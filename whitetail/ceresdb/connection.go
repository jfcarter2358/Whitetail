package ceresdb

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"whitetail/logger"
)

type Client struct {
	Connection string
	Auth       string
	Databases  map[string]*Database
}

type Collection struct {
	Name       string
	Auth       string
	Connection string
	Schema     interface{}
	Database   string
}

type Database struct {
	Name        string
	Connection  string
	Auth        string
	Collections map[string]*Collection
}

type Cursor struct {
	Data   []interface{}
	Length int
	Idx    int
	Closed bool
}

type Filter map[string]interface{}

func Connect(connectionString string) (*Client, error) {
	u, err := url.Parse(connectionString)
	if err != nil {
		panic(err)
	}

	// Validate connection string
	if u.Scheme != "https" && u.Scheme != "http" {
		return nil, fmt.Errorf("invalid connection scheme: %s", u.Scheme)
	}
	host, port, err := net.SplitHostPort(u.Host)
	if err != nil {
		return nil, err
	}
	password, present := u.User.Password()
	if !present {
		return nil, fmt.Errorf("no password provided")
	}

	dbs := make(map[string]*Database)
	connectionString = fmt.Sprintf("%s://%s:%s/api/query", u.Scheme, host, port)
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", u.User.Username(), password)))
	if len(u.Path) > 0 {
		dbName := u.Path[1:len(u.Path)]
		if dbName != "" {
			dbs[dbName] = &Database{
				Name:        dbName,
				Auth:        auth,
				Connection:  connectionString,
				Collections: make(map[string]*Collection),
			}
		}
	}

	c := &Client{
		Connection: connectionString,
		Databases:  dbs,
		Auth:       auth,
	}

	return c, nil
}

func (c *Client) Ping() error {
	_, err := RawQuery(c.Connection, c.Auth, "get database")
	return err
}

func (c *Client) Database(name string) (*Database, error) {
	d := &Database{
		Name:        name,
		Auth:        c.Auth,
		Connection:  c.Connection,
		Collections: make(map[string]*Collection),
	}

	databases, err := RawQuery(c.Connection, c.Auth, "get database")
	if err != nil {
		return nil, err
	}

	for _, db := range databases {
		if db.(string) == name {
			return nil, &DatabaseExists{}
		}
	}
	if _, err := RawQuery(c.Connection, c.Auth, fmt.Sprintf("create database %s", name)); err != nil {
		return nil, err
	}

	c.Databases[name] = d

	return d, nil
}

func (c *Client) Delete(name string) error {
	_, err := RawQuery(c.Connection, c.Auth, fmt.Sprintf("delete database %s", name))
	return err
}

func (d *Database) Collection(name string, schema interface{}) (*Collection, error) {

	cl := &Collection{
		Name:       fmt.Sprintf("%s.%s", d.Name, name),
		Auth:       d.Auth,
		Connection: d.Connection,
		Schema:     schema,
		Database:   d.Name,
	}

	collections, err := RawQuery(d.Connection, d.Auth, fmt.Sprintf("get collection from %s", d.Name))
	if err != nil {
		return nil, err
	}

	for _, col := range collections {
		if col.(string) == name {
			return nil, &CollectionExists{}
		}
	}

	schemaData, _ := json.Marshal(&schema)
	if _, err := RawQuery(d.Connection, d.Auth, fmt.Sprintf("create collection %s in %s with schema %s", name, d.Name, string(schemaData))); err != nil {
		return nil, err
	}

	d.Collections[name] = cl

	return cl, err
}

func (d *Database) Delete(name string) error {
	_, err := RawQuery(d.Connection, d.Auth, fmt.Sprintf("delete collection %s from %s", name, d.Name))
	return err
}

func (cr *Cursor) Next() bool {
	if cr.Closed {
		return false
	}
	cr.Idx += 1
	if cr.Idx == cr.Length {
		return false
	}
	return true
}

func (cr *Cursor) Close() {
	cr.Closed = true
}

func (cr *Cursor) Decode(out interface{}) error {
	if cr.Idx == -1 {
		return fmt.Errorf("need to call next on cursor before decoding data")
	}
	data, err := json.Marshal(&cr.Data[cr.Idx])
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, out)
	return err
}

func (cl *Collection) Find(filter Filter) (*Cursor, error) {
	cr := &Cursor{
		Idx:    -1,
		Closed: false,
	}

	filterBytes, err := json.Marshal(filter)
	if err != nil {
		return nil, fmt.Errorf("cannot marshal filter object to json")
	}
	filterString := string(filterBytes)

	var data []interface{}
	if len(filter) > 0 {
		data, err = RawQuery(cl.Connection, cl.Auth, fmt.Sprintf("get record from %s WHERE %s", cl.Name, filterString))
		if err != nil {
			return nil, err
		}
	} else {
		data, err = RawQuery(cl.Connection, cl.Auth, fmt.Sprintf("get record from %s", cl.Name))
		if err != nil {
			return nil, err
		}
	}

	cr.Data = data
	cr.Length = len(data)

	return cr, err
}

func (cl *Collection) InsertOne(in interface{}) error {
	inData, _ := json.Marshal(&in)
	_, err := RawQuery(cl.Connection, cl.Auth, fmt.Sprintf("insert record %s into %s", inData, cl.Name))
	return err
}

func (cl *Collection) InsertMany(in []interface{}) error {
	inData, _ := json.Marshal(&in)
	_, err := RawQuery(cl.Connection, cl.Auth, fmt.Sprintf("insert record %s into %s", inData, cl.Name))
	return err
}

func (cl *Collection) Upsert(filter Filter, obj interface{}) error {
	filterBytes, err := json.Marshal(filter)
	if err != nil {
		return fmt.Errorf("cannot marshal filter object to json")
	}
	filterString := string(filterBytes)

	dataBytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("cannot marshal upsert object to json")
	}
	dataString := string(dataBytes)

	_, err = RawQuery(cl.Connection, cl.Auth, fmt.Sprintf("update record %s with %s where %s", cl.Name, dataString, filterString))
	return err
}

func (cl *Collection) Delete(filter Filter) error {
	filterBytes, err := json.Marshal(filter)
	if err != nil {
		return fmt.Errorf("cannot marshal filter object to json")
	}
	filterString := string(filterBytes)

	_, err = RawQuery(cl.Connection, cl.Auth, fmt.Sprintf("delete record from %s where %s", cl.Name, filterString))
	return err
}

func RawQuery(connectionString, authString, queryString string) ([]interface{}, error) {
	payload := map[string]string{
		"auth":  authString,
		"query": queryString,
	}
	data, err := json.Marshal(payload)
	if err != nil {
		logger.Errorf("", "An error occurred marshalling payload: %v", err)
		return nil, err
	}

	payloadBytes := bytes.NewBuffer(data)
	resp, err := http.Post(connectionString, "application/json", payloadBytes)
	if err != nil {
		logger.Errorf("", "An error occurred talking to the CeresDB instance: %v", err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Errorf("", "An error occurred reading response body %v", err)
		return nil, err
	}
	outputString := string(body)

	if resp.StatusCode >= 400 {
		logger.Errorf("", "An error occurred in CeresDB: %d | %v", resp.StatusCode, outputString)
		return nil, errors.New(outputString)
	}

	if outputString == "null" {
		return nil, nil
	}

	// We got back a dictionary not a list, so that means an error was thrown
	var outputData []interface{}
	err = json.Unmarshal([]byte(outputString), &outputData)
	if err != nil {
		return nil, err
	}
	return outputData, nil
}
