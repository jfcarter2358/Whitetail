package logging

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	"whitetail/config"

	"github.com/jfcarter2358/ceresdb-go/connection"
)

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Block) Do() {
	if tcf.Finally != nil {

		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

type LogMessageInput struct {
	Timestamp  string      `json:"@timestamp"`
	Message    string      `json:"message"`
	Service    string      `json:"appName"`
	LoggerName string      `json:"logger_name"`
	Level      string      `json:"level"`
	StackTrace string      `json:"stack_trace"`
	Fields     *FieldInput `json:"fields"`
}

type FieldInput struct {
	Level      string `json:"severity"`
	LoggerName string `json:"application"`
	Service    string `json:"hostname"`
}

type Log struct {
	Text      string
	Level     string
	Timestamp string
	Service   string
	ID        string
	Year      int
	Month     int
	Day       int
	Hour      int
	Minute    int
	Second    int
}

type LogRequestInput struct {
	LineLimit   string   `json:"limit"`
	KeywordList string   `json:"keyword-list"`
	LogLevels   []string `json:"log-levels"`
}

var Services []string

var allClients map[*Client]int

type Client struct {
	// incoming chan string
	reader     *bufio.Reader
	conn       net.Conn
	connection *Client
}

func (client *Client) Read() {
	line, err := client.reader.ReadString('\n')
	if err == nil || err == io.EOF {
		parseData(line)
	}

	Block{
		Try: func() {
			client.conn.Close()
			delete(allClients, client)
			if client.connection != nil {
				client.connection.connection = nil
			}
			client = nil
		},
		Catch: func(e Exception) {
			fmt.Printf("Caught %v\n", e)
		},
		Finally: func() {
			// fmt.Println("Finally...")
		},
	}.Do()
}

func (client *Client) Listen() {
	client.Read()
}

func NewClient(connection net.Conn) *Client {
	reader := bufio.NewReader(connection)

	client := &Client{
		// incoming: make(chan string),
		conn:   connection,
		reader: reader,
	}
	client.Listen()

	return client
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func InitLogger() {
	data, _ := connection.Query(fmt.Sprintf("GET RECORD %s.logs service", config.Config.DB.Name))
	for _, datum := range data {
		if contains(Services, datum["service"].(string)) == false {
			Services = append(Services, datum["service"].(string))
		}
	}
}

func formatLogMessage(data *LogMessageInput) string {
	// log line format is
	// [ TIMESTAMP ] [ SERVICE ] [ LEVEL ] [ LOGGER ] MESSAGE
	message := "[" + data.Timestamp + "] "
	message = message + "[" + data.Service + "] "
	message = message + "[<span class=\"" + data.Level + "\">" + data.Level + "</span>] "
	if config.Config.Logging.ConciseLogger {
		loggerList := strings.Split(data.LoggerName, ".")
		loggerName := ""
		for index, part := range loggerList {
			if index == len(loggerList)-1 {
				loggerName = loggerName + part
			} else {
				loggerName = loggerName + string(part[0]) + "."
			}
		}
		if config.Config.Logging.HoverableLongLogger {
			message = message + "[<span class=\"w3-tooltip\">" + loggerName + "<span class=\"w3-text\">&nbsp;(" + data.LoggerName + ")</span></span>] "
		} else {
			message = message + "[" + loggerName + "] "
		}
		message = strings.TrimSuffix(message, "\n")
		message = strings.TrimSuffix(message, "<br>")
	} else {
		message = message + "[" + data.LoggerName + "] "
	}
	message = message + data.Message
	if data.StackTrace != "" {
		message += "<br>" + strings.ReplaceAll(strings.ReplaceAll(data.StackTrace, "\t", "&emsp;&emsp;&emsp;"), "\n", "<br>")
	}
	return message
}

func Query(query string) ([]string, string) {
	logs := make([]string, 0)

	data, err := connection.Query(query)
	if err != nil {
		return []string{}, err.Error()
	}
	for _, datum := range data {
		logs = append(logs, datum["message"].(string))
	}

	return logs, ""
}

func CreateNewLog(text, level, timestamp, service, rawMessage string) (*Log, error) {

	layout := "2006-01-02T15:04:05"
	t, _ := time.Parse(layout, timestamp)

	log := make(map[string]interface{})
	log["message"] = text
	log["level"] = level
	log["service"] = service
	log["timestamp"] = timestamp
	log["year"] = t.Year()
	log["month"] = int(t.Month())
	log["day"] = t.Day()
	log["hour"] = t.Hour()
	log["minute"] = t.Minute()
	log["second"] = t.Second()

	json_data, _ := json.Marshal(log)

	connection.Query(fmt.Sprintf("POST RECORD %s.logs %v", config.Config.DB.Name, string(json_data)))

	if config.Config.PrintElevatedMessages {
		if (level == "INFO") || (level == "DEBUG") || (level == "TRACE") {
			// ignore these log levels
		} else {
			fmt.Printf(">>> %v %v %v\n", service, level, rawMessage)
		}
	}

	return nil, nil
}

/* -- TCP and UDP -- */

func StartTCPServer(conn_port int) {
	port := strconv.Itoa(conn_port)

	fmt.Println("Listening to TCP on " + ":" + port)

	allClients = make(map[*Client]int)
	listener, _ := net.Listen("tcp", ":"+port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}
		client := NewClient(conn)
		for clientList, _ := range allClients {
			if clientList.connection == nil {
				client.connection = clientList
				clientList.connection = client
				// fmt.Println("Connected")
			}
		}
		allClients[client] = 1
	}
}

func StartUDPServer(conn_port int) {
	ServerAddr, err := net.ResolveUDPAddr("udp", ":"+strconv.Itoa(conn_port))
	if err != nil {
		panic(err)
	}
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	if err != nil {
		panic(err)
	}
	defer ServerConn.Close()
	fmt.Println("Listening to UDP on " + ":" + strconv.Itoa(conn_port))
	buf := make([]byte, 1024)
	for {
		n, _, _ := ServerConn.ReadFromUDP(buf)
		parseData(string(buf[0:n]))
	}
}

func parseData(data string) {
	messages := strings.Split(data, "\n")

	fmt.Printf("%v\n", messages)
	// messages[0] = leftover + messages[0]
	for i := 0; i < len(messages); i++ {
		shouldParse := true
		messages[i] = strings.TrimSpace(messages[i])
		messages[i] = strings.TrimSuffix(messages[i], "\n")
		if i == len(messages)-1 {
			if strings.HasSuffix(messages[i], "}") == false {
				shouldParse = false
			}
		}
		if shouldParse {
			var input *LogMessageInput
			var byt = []byte(messages[i])
			byt = bytes.Trim(byt, "\x00")

			var readErr = json.Unmarshal(byt, &input)
			if readErr != nil {
				fmt.Println(readErr)
			}

			if input.Fields != nil {
				input.Service = input.Fields.Service
				input.Level = strings.ToUpper(input.Fields.Level)
				input.LoggerName = input.Fields.LoggerName
			}
			if input.Service != "" {

				if input.Timestamp == "" {
					input.Timestamp = time.Now().Format("2006-01-02T15:04:05")
				}

				formatted := formatLogMessage(input)

				contained := false
				for _, service := range Services {
					if service == input.Service {
						contained = true
						break
					}
				}
				if contained == false {
					Services = append(Services, input.Service)
				}

				CreateNewLog(formatted, input.Level, input.Timestamp, input.Service, input.Message)
			}
		}
	}
}

// Handles incoming requests.
func handleTCPRequest(port string) {
	// Listen for incoming connections.
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}

	defer l.Close()
	// Make a buffer to hold incoming data.
	buf := make([]byte, 65536)

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}

	// go handleTCPRequest(port)

	defer conn.Close()

	// Read the incoming connection into the buffer.
	// reqLen, err := conn.Read(buf)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	buf = bytes.Trim(buf, "\x00")
	parseData(string(buf))

	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.

	return
}

func Cleanup() {
	if config.Config.Logging.MaxAgeDays == -1 {
		return
	}
	for true {
		poll, err := time.ParseDuration(config.Config.Logging.PollRate)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Sleeping before cleanup...")
		time.Sleep(poll)
		log.Println("Cleaning up...")
		t := time.Now()
		cutoff := t.AddDate(0, 0, -1*config.Config.Logging.MaxAgeDays)

		ids, _ := connection.Query(fmt.Sprintf("GET RECORD %s.logs .id | FILTER \"year\" <= %d AND \"month\" <= %d AND \"day\" <= %d AND \"hour\" <= %d AND \"minute\" <= %d AND \"second\" <= %d", config.Config.DB.Name, cutoff.Year(), int(cutoff.Month()), cutoff.Day(), cutoff.Hour(), cutoff.Minute(), cutoff.Second()))
		json_data, _ := json.Marshal(ids)
		_, errorMessage := connection.Query(fmt.Sprintf("DELETE RECORD %s.logs %v", config.Config.DB.Name, string(json_data)))

		log.Println(errorMessage)
		log.Println("Cleaned up logs")
	}
}

func ParseFileData(data, service string) {
	messages := strings.Split(data, "\n")

	for i := 0; i < len(messages); i++ {

		messages[i] = strings.TrimSpace(messages[i])
		messages[i] = strings.TrimSuffix(messages[i], "\n")

		var input LogMessageInput = LogMessageInput{}

		input.Message = messages[i]
		input.Service = service
		input.Level = "INFO"
		input.LoggerName = "wt.filelogger"
		input.Timestamp = time.Now().Format("2006-01-02T15:04:05")

		formatted := formatLogMessage(&input)

		contained := false
		for _, service := range Services {
			if service == input.Service {
				contained = true
				break
			}
		}
		if contained == false {
			Services = append(Services, input.Service)
		}

		CreateNewLog(formatted, input.Level, input.Timestamp, input.Service, input.Message)
	}
}
