package Logging

import (
    "fmt"
    "net"
    "os"
    "encoding/json"
    "strings"
    "bytes"
    "time"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "log"
    "net/url"
    "errors"
    "github.com/google/uuid"
    "strconv"
    "whitetail/index"
    "sort"
)

type LogMessageInput struct {
	Timestamp  string `json:"@timestamp"`
	Message    string `json:"message"`
    Service    string `json:"appName"`
    LoggerName string `json:"logger_name"`
    Level      string `json:"level"`
    StackTrace string `json:"stack_trace"`
}

type Log struct {
    Text       string
    Level      string
    Timestamp  string
    Service    string
    ID         string
}

type LogRequestInput struct {
    LineLimit   string `json:"limit"`
    KeywordList string `json:"keyword-list"`
}

var DB *gorm.DB
var Services []string

func formatLogMessage(data *LogMessageInput) string{
    // log line format is
    // [ TIMESTAMP ] [ LEVEL ] [ LOGGER ] MESSAGE
    message := "[" + data.Timestamp + "] "
    message = message + "[<span class=\"" + data.Level + "\">" + data.Level + "</span>] "
    message = message + "[" + data.LoggerName + "] "
    message = message + data.Message + "<br>"
    if data.StackTrace != "" {
        message += strings.ReplaceAll(strings.ReplaceAll(data.StackTrace, "\t", "&emsp;&emsp;&emsp;"), "\n", "<br>")
    }
    return message
}

/* -- Database interactions -- */

func ConnectDataBase(db_type, db_user, db_pass, db_host, db_port string) {
	dsn := url.URL{
		User:     url.UserPassword(db_user, db_pass),
		Scheme:   db_type,
		Host:     fmt.Sprintf("%s:%s", db_host, db_port),
		Path:     "whitetail",
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	database, err := gorm.Open("postgres", dsn.String())

	if err != nil {
		log.Println(err)
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Log{})

    DB = database

    allLogs := GetAllLogs()

    for _, log := range allLogs {
        contained := false
        for _, service := range Services {
            if service == log.Service {
                contained = true
                break
            }
        }
        if contained == false {
            Services = append(Services, log.Service)
        }
    }
}

func GetLogCount() int {
	logs := GetAllLogs()
	return len(logs)
}

func GetAllLogs() []Log {
	var logs []Log
	DB.Find(&logs)
	
	return logs
}

func GetLogsFromIndex(keyString, service string, limit int) []Log {
    keys := strings.Split(keyString, ",")
    logs := []Log{}
    for _, key := range keys {
        index, err := Index.GetIndexByKey(key) 
        if err != nil {
            continue
        } else {
            ids := strings.Split(index.IDs, ",")
            for _, id := range ids {
                newLog := &Log{}
                if service != "" {
                    newLog, err = GetLogByServiceAndID(id, service)
                } else {
                    newLog, err = GetLogByID(id)
                }
                if err != nil {
                    continue
                }
                logs = append(logs, *newLog)
            }
        }
    }
    sort.Slice(logs[:], func(i, j int) bool {
        return logs[i].Timestamp < logs[j].Timestamp
    })
    if limit < len(logs) {
        return logs[len(logs) - limit:]
    }
    return logs
}

func DeleteLogByID(id string) error{
	log, err := GetLogByID(id)
	if err != nil {
		return errors.New("Log not found")
	}
	DB.Delete(&log)
	return nil
}

func GetLogByID(id string) (*Log, error) {
	var log Log

	if err := DB.Where("id = ?", id).First(&log).Error; err != nil {
		return nil, errors.New("Log not found")
	}
	return &log, nil
}

func GetLogsByService(keyString, service string, limit int) []Log {
    var logs []Log
    
    if keyString != "" {
        logs = GetLogsFromIndex(keyString, service, limit)
    } else {
        DB.Where("service = ?", service).Find(&logs)

        sort.Slice(logs[:], func(i, j int) bool {
            return logs[i].Timestamp < logs[j].Timestamp
        })
        if limit < len(logs) {
            return logs[len(logs) - limit:]
        }
    }
    return logs
}

func GetLogByServiceAndID(id, service string) (*Log, error) {
	var log Log

	if err := DB.Where("id = ? AND service = ?", id, service).First(&log).Error; err != nil {
		return nil, errors.New("Log not found")
	}
	return &log, nil
}

func CreateNewLog(text, level, timestamp, service, rawMessage string) (*Log, error) {
	id := uuid.New().String()

	log := Log{Text: text, Level: level, Timestamp: timestamp, Service: service, ID: id}

    DB.Create(&log)
    
    Index.ParseLog(rawMessage, id)

	return &log, nil
}

/* -- TCP and UDP -- */

func StartTCPServer(conn_port int) {
    port := strconv.Itoa(conn_port)
    // Listen for incoming connections.
    l, err := net.Listen("tcp", ":" + port)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening to TCP on " + ":" + port)
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleTCPRequest(conn)
    }
}

func StartUDPServer(conn_port int) {
    ServerConn, _ := net.ListenUDP("udp", &net.UDPAddr{IP:[]byte{},Port:conn_port,Zone:""})
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
    // messages[0] = leftover + messages[0]
    for i := 0; i < len(messages); i++ {
        log.Println(messages[i])
        shouldParse := true
        messages[i] = strings.TrimSpace(messages[i])
        messages[i] = strings.TrimSuffix(messages[i], "\n")
        if i == len(messages) - 1 {
            if strings.HasSuffix(messages[i], "}") == false {
                shouldParse = false
            }
        }
        if shouldParse {
            var input *LogMessageInput
            var byt = []byte(messages[i])
            byt = bytes.Trim(byt, "\x00")
            fmt.Println("TO UNMARSHAL : " + messages[i])
            var readErr = json.Unmarshal(byt, &input)
            if readErr != nil {
                fmt.Println(readErr)
            }

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

// Handles incoming requests.
func handleTCPRequest(conn net.Conn) {
    // Make a buffer to hold incoming data.
    buf := make([]byte, 65536)
    // Read the incoming connection into the buffer.
    // reqLen, err := conn.Read(buf)
    _, err := conn.Read(buf)
    if err != nil {
        fmt.Println("Error reading:", err.Error())
    }
    buf = bytes.Trim(buf, "\x00")
    parseData(string(buf))

    // Send a response back to person contacting us.
    conn.Write([]byte("Message received."))
    // Close the connection when you're done with it.
    conn.Close()
}