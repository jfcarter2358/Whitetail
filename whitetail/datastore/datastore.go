package datastore

import (
	"encoding/json"
	"fmt"
	"strings"
	"whitetail/config"
	"whitetail/logger"
	"whitetail/observer"

	"whitetail/ceresdb"
)

var Database *ceresdb.Database

func Init() error {
	c, err := ceresdb.Connect(config.Config.CeresDBConnectionString)
	if err != nil {
		return err
	}

	Database = c.Databases[config.Config.CeresDBDatabaseName]
	if _, err := c.Database(config.Config.CeresDBDatabaseName); err != nil {
		if _, ok := err.(*ceresdb.DatabaseExists); !ok {
			return nil
		}
	}

	return nil
}

func AddStream(oo, ss string, o observer.Observer) error {
	logger.Tracef("", "adding stream %s for observer %s with schema %v", ss, oo, o.Streams[ss].Schema)
	cl, err := Database.Collection(fmt.Sprintf("%s__%s", oo, ss), o.Streams[ss].Schema)
	if err != nil {
		if _, ok := err.(*ceresdb.CollectionExists); !ok {
			return nil
		}
	}
	Database.Collections[fmt.Sprintf("%s__%s", oo, ss)] = cl
	return nil
}

func AddData(data, stream, observer string) error {
	lines := strings.Split(data, "\n")
	if len(lines) == 0 {
		return nil
	}

	records := []interface{}{}

	for _, line := range lines {
		var datum interface{}
		if err := json.Unmarshal([]byte(line), &datum); err == nil {
			records = append(records, datum)
		}
	}

	return Database.Collections[fmt.Sprintf("%s__%s", observer, stream)].InsertMany(records)
}

func GetData(stream, observer string, filter interface{}) ([]interface{}, error) {
	f := filter.(ceresdb.Filter)
	cur, err := Database.Collections[fmt.Sprintf("%s__%s", observer, stream)].Find(f)
	if err != nil {
		return nil, err
	}
	return cur.Data, nil
}
