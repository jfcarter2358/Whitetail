package Logging

import (
    "fmt"
    "net"
    "os"
    "encoding/json"
    "strings"
    "bytes"
    "time"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "gorm.io/driver/postgres"
    "log"
    "net/url"
    "errors"
    "github.com/google/uuid"
    "strconv"
    "whitetail/index"
    "sort"
    "whitetail/config"
    "gorm.io/gorm/logger"
    // "github.com/expectedsh/go-sonic/sonic"
)

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
    Text       string
    Level      string
    Timestamp  string
    Service    string
    ID         string
    Year       int
    Month      int
    Day        int
    Hour       int
    Minute     int
    Second     int
}

type LogRequestInput struct {
    LineLimit   string   `json:"limit"`
    KeywordList string   `json:"keyword-list"`
    LogLevels   []string `json:"log-levels"`
}

var DB *gorm.DB
var Services []string
// var ingester sonic.Ingestable
// var search sonic.Searchable

func formatLogMessage(data *LogMessageInput) string{
    // log line format is
    // [ TIMESTAMP ] [ SERVICE ] [ LEVEL ] [ LOGGER ] MESSAGE
    message := "[" + data.Timestamp + "] "
    message = message + "[" + data.Service + "] "
    message = message + "[<span class=\"" + data.Level + "\">" + data.Level + "</span>] "
    if Config.Config.Logging.ConciseLogger {
        loggerList := strings.Split(data.LoggerName, ".")
        loggerName := ""
        for index, part := range loggerList {
            if index == len(loggerList) - 1 {
                loggerName = loggerName + part
            } else {
                loggerName = loggerName + string(part[0]) + "."
            }
        }
        if Config.Config.Logging.HoverableLongLogger {
            message = message + "[<span class=\"w3-tooltip\">" + loggerName + "<span class=\"w3-text\">&nbsp;(" + data.LoggerName + ")</span></span>] "
        } else {
            message = message + "[" + loggerName + "] "
        }
    } else {
        message = message + "[" + data.LoggerName + "] "
    }
    message = message + data.Message + "<br>"
    if data.StackTrace != "" {
        message += strings.ReplaceAll(strings.ReplaceAll(data.StackTrace, "\t", "&emsp;&emsp;&emsp;"), "\n", "<br>")
    }
    return message
}

/* -- Database interactions -- */

func ConnectDataBase(db_type string, postgresConfig *Config.PostgresConfigObject, sqliteConfig *Config.SqliteConfigObject) {
    var err error
    var database *gorm.DB

    if db_type == "postgres" {
        dsn := url.URL{
            User:     url.UserPassword(postgresConfig.Username, postgresConfig.Password),
            Scheme:   db_type,
            Host:     fmt.Sprintf("%s:%d", postgresConfig.Host, postgresConfig.Port),
            Path:     postgresConfig.Database,
            RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
        }
    	database, err = gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
            Logger: logger.Default.LogMode(logger.Silent),
        })
    } else if db_type == "sqlite" {
        database, err = gorm.Open(sqlite.Open(sqliteConfig.Path), &gorm.Config{
            Logger: logger.Default.LogMode(logger.Silent),
        })
    }

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

    // connect to sonic
    /*
    ingester, err = sonic.NewIngester("localhost", 1491, "Pi314159")
    if err != nil {
        panic(err)
    }
    search, err = sonic.NewSearch("localhost", 1491, "Pi314159")
    if err != nil {
        panic(err)
    }
    */
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

func GetLogsFromIndex(keyString, service string, limit int, logLevels []string) []Log {
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
    filteredLogs := []Log{}
    for _, log := range logs {
        levelMatch := false
        for _, level := range logLevels {
            if log.Level == level {
                levelMatch = true
                break
            }
        }
        if levelMatch == true {
            filteredLogs = append(filteredLogs, log)
        }
    }
    if limit < len(filteredLogs) {
        return filteredLogs[len(filteredLogs) - limit:]
    }
    return filteredLogs
}

func DeleteLogByID(id string) error{
	log, err := GetLogByID(id)
	if err != nil {
		return errors.New("Log not found")
	}
	DB.Delete(&log)
    Index.DeleteElementFromIndices(id)
	return nil
}

func GetLogByID(id string) (*Log, error) {
	var log Log

	if err := DB.Where("id = ?", id).First(&log).Error; err != nil {
		return nil, errors.New("Log not found")
	}
	return &log, nil
}

func GetLogsByService(keyString, service string, limit int, logLevels []string) []Log {
    var logs []Log
    
    if keyString != "" {
        logs = GetLogsFromIndex(keyString, service, limit, logLevels)
    } else {
        DB.Where("service = ?", service).Find(&logs)

        sort.Slice(logs[:], func(i, j int) bool {
            return logs[i].Timestamp < logs[j].Timestamp
        })
        filteredLogs := []Log{}
        for _, log := range logs {
            levelMatch := false
            for _, level := range logLevels {
                if log.Level == level {
                    levelMatch = true
                    break
                }
            }
            if levelMatch == true {
                filteredLogs = append(filteredLogs, log)
            }
        }
        if limit < len(filteredLogs) {
            return filteredLogs[len(filteredLogs) - limit:]
        }
        return filteredLogs
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

func Query(query string) []Log{
    var logs []Log
    log.Println(query)
    // DB.Where(query).Find(&logs)
    DB.Where("service = ?", "foobar").Find(&logs)
    log.Println(len(logs))
    return logs
}

func CreateNewLog(text, level, timestamp, service, rawMessage string) (*Log, error) {
	id := uuid.New().String()

    layout := "2006-01-02T15:04:05"
	t, _ := time.Parse(layout, timestamp)

	log := Log{Text: text, Level: level, Timestamp: timestamp, Service: service, ID: id, Year: t.Year(), Month: int(t.Month()), Day: t.Day(), Hour: t.Hour(), Minute: t.Minute(), Second: t.Second()}

    DB.Create(&log)
    
    Index.ParseLog(rawMessage, id, timestamp, level, service)

	return &log, nil
}

// func AddLogToSonic(text) {

// }

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
    ServerAddr, err := net.ResolveUDPAddr("udp", ":" + strconv.Itoa(conn_port))
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
    // messages[0] = leftover + messages[0]
    for i := 0; i < len(messages); i++ {
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
            // fmt.Println("TO UNMARSHAL : " + messages[i])
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

func Cleanup() {
    if Config.Config.Logging.MaxAgeDays == -1 {
        return
    }
    for true {
        poll, err := time.ParseDuration(Config.Config.Logging.PollRate)
        if err != nil {
            log.Println(err)
            return
        }
        log.Println("Sleeping before cleanup...")
        time.Sleep(poll)
        log.Println("Cleaning up...")
        t := time.Now()
        cutoff := t.AddDate(0, 0, -1 * Config.Config.Logging.MaxAgeDays).Format("2006-01-02T15:04:05")
        logs := GetAllLogs()
        cleanupCount := 0
        for _, log := range logs {
            if log.Timestamp < cutoff{
                cleanupCount += 1
                DeleteLogByID(log.ID)
            }
        }
        log.Println("Cleaned up " + strconv.Itoa(cleanupCount) + " logs")
    }
}