// handler.config.go

package Config

import (
	"log"
	"os"
	"io/ioutil"
	"encoding/json"
	"strings"
	"io"
)

type ConfigObject struct {
	HTTPPort int                  `json:"http-port" binding:"required"`
	TCPPort  int                  `json:"tcp-port" binding:"required"`
	UDPPort  int                  `json:"udp-port" binding:"required"`
	BasePath string               `json:"basepath" binding:"required"`
	Database DatabaseConfigObject `json:"database" binding:"required"`
	Logging  LoggingConfigObject  `json:"logging" binding:"required"`
	Branding BrandingConfigObject `json:"branding" binding:"required"`
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
	Database string `json:"database"`
}

type BrandingConfigObject struct {
	PrimaryColor   ColorConfigObject `json:"primary_color" binding:"required"`
	SecondaryColor ColorConfigObject `json:"secondary_color" binding:"required"`
	TertiaryColor  ColorConfigObject `json:"tertiary_color" binding:"required"`
	INFOColor      string `json:"INFO_color" binding:"required"`
	WARNColor      string `json:"WARN_color" binding:"required"`
	DEBUGColor     string `json:"DEBUG_color" binding:"required"`
	TRACEColor     string `json:"TRACE_color" binding:"required"`
	ERRORColor     string `json:"ERROR_color" binding:"required"`
}

type ColorConfigObject struct {
	Background string `json:"background" binding:"required"`
	Text string `json:"text" binding:"required"`
}

type SqliteConfigObject struct {
	Path string `json:"path"`
}

type LoggingConfigObject struct { 
	MaxAgeDays int `json:"max-age-days"`
	PollRate string `json:"poll-rate"`
	ConciseLogger bool `json:"concise-logger"`
}

var Config ConfigObject

var Defaults BrandingConfigObject

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

	json.Unmarshal(byteValue, &Config)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	if Config.Database.Postgres != nil {
		log.Println("postgres")
		Config.Database.Type = "postgres"
	} else if Config.Database.Sqlite != nil {
		log.Println("sqlite")
		Config.Database.Type = "sqlite"
	} else {
		panic("No database config found")
	}

	UpdateBranding()
	InitLogoIcon()

	Defaults.PrimaryColor = ColorConfigObject{Background: "#C3C49E", Text: "#000000"}
	Defaults.SecondaryColor = ColorConfigObject{Background: "#8F7E4F", Text: "#000000"}
	Defaults.TertiaryColor = ColorConfigObject{Background: "#524632", Text: "#ffffff"}
	Defaults.INFOColor = "#4F772D"
	Defaults.WARNColor = "#E24E1B"
	Defaults.DEBUGColor = "#2B50AA"
	Defaults.TRACEColor = "#610345"
	Defaults.ERRORColor = "#95190C"

	return &Config
}

func UpdateBranding() {
	dat, err := ioutil.ReadFile("static/css/branding.template.css")
    if err != nil {
        panic(err)
	}
	out := string(dat)
	out = strings.ReplaceAll(out, "[[T1]]", Config.Branding.PrimaryColor.Text)
	out = strings.ReplaceAll(out, "[[BG1]]", Config.Branding.PrimaryColor.Background)
	out = strings.ReplaceAll(out, "[[T2]]", Config.Branding.SecondaryColor.Text)
	out = strings.ReplaceAll(out, "[[BG2]]", Config.Branding.SecondaryColor.Background)
	out = strings.ReplaceAll(out, "[[T3]]", Config.Branding.TertiaryColor.Text)
	out = strings.ReplaceAll(out, "[[BG3]]", Config.Branding.TertiaryColor.Background)

	out = strings.ReplaceAll(out, "[[INFO]]", Config.Branding.INFOColor)
	out = strings.ReplaceAll(out, "[[TRACE]]", Config.Branding.TRACEColor)
	out = strings.ReplaceAll(out, "[[DEBUG]]", Config.Branding.DEBUGColor)
	out = strings.ReplaceAll(out, "[[WARN]]", Config.Branding.WARNColor)
	out = strings.ReplaceAll(out, "[[ERROR]]", Config.Branding.ERRORColor)

	f, err := os.Create("static/css/branding.css")
    if err != nil {
        panic(err)
    }
	defer f.Close()
	
	_, err = f.WriteString(out)
	if err != nil { 
		panic(err)
	}
}

func InitLogoIcon() {

	if _, err := os.Stat("config/custom/logo/logo.png"); err == nil {
		log.Println("Copying custom logo file")
		source, err := os.Open("config/custom/logo/logo.png")
        if err != nil {
                panic(err)
        }
        defer source.Close()

        destination, err := os.Create("static/img/logo.png")
        if err != nil {
                panic(err)
        }
        defer destination.Close()
        _, err = io.Copy(destination, source)
	}

	if _, err := os.Stat("config/custom/icon/favicon.png"); err == nil {
		log.Println("Copying custom icon file")
		source, err := os.Open("config/custom/icon/favicon.png")
        if err != nil {
                panic(err)
        }
        defer source.Close()

        destination, err := os.Create("static/img/favicon.png")
        if err != nil {
                panic(err)
        }
        defer destination.Close()
        _, err = io.Copy(destination, source)
	}
}