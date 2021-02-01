// handler.config.go

package Config

import (
	"log"
	"os"
	"io/ioutil"
	"encoding/json"
)

type ConfigObject struct {
	HTTPPort int                  `json:"http-port"`
	TCPPort  int                  `json:"tcp-port"`
	UDPPort  int                  `json:"udp-port"`
	BasePath string               `json:"basepath"`
	Database DatabaseConfigObject `json:"database"`
	Logging  LoggingConfigObject  `json:"logging"`
}

type DatabaseConfigObject struct {
	Type     string
	Postgres *PostgresConfigObject `json:"postgres"`
	Sqlite   *SqliteConfigObject   `json:"sqlite"`
}

type PostgresConfigObject struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type SqliteConfigObject struct {
	Path string `json:"path"`
}

type LoggingConfigObject struct { 
	MaxAge string `json:"max-age"`
}

func ReadConfigFile() *ConfigObject{
	// Open our jsonFile
	jsonFile, err := os.Open("config/config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println("Unable to read json file")
		panic(err)
	}
	
	log.Println("Successfully Opened config/config.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var config ConfigObject
	json.Unmarshal(byteValue, &config)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	log.Println(config.Database.Postgres)
	log.Println(config.Database.Sqlite)

	if config.Database.Postgres != nil {
		log.Println("postgres")
		config.Database.Type = "postgres"
	} else if config.Database.Sqlite != nil {
		log.Println("sqlite")
		config.Database.Type = "sqlite"
	} else {
		panic("No database config found")
	}
	return &config
}