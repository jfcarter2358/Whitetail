// handler.config.go

package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"whitetail/logger"
)

const DEFAULT_CONFIG_PATH = "/whitetail/config/config.json"
const ENV_PREFIX = "WHITETAIL_"

type ConfigObject struct {
	Host                    string               `json:"host" env:"HOST"`
	Port                    int                  `json:"port" env:"PORT"`
	TCPPort                 int                  `json:"tcp_port" env:"TCP_PORT"`
	UDPPort                 int                  `json:"udp_port" env:"UDP_PORT"`
	BasePath                string               `json:"base_path" env:"BASE_PATH"`
	Branding                BrandingConfigObject `json:"branding" env:"BRANDING"`
	CeresDBConnectionString string               `json:"ceres_db_connection_string" env:"CERESDB_CONNECTION_STRING"`
	CeresDBDatabaseName     string               `json:"ceresdb_database_name" env:"CERESDB_DATABASE_NAME"`
	LogLevel                string               `json:"log_level" env:"LOG_LEVEL"`
	LogFormat               string               `json:"log_format" env:"LOG_FORMAT"`
	OperationPath           string               `json:"operation_path" env:"OPERATION_PATH"`
}

type BrandingConfigObject struct {
	PrimaryColor   ColorConfigObject `json:"primary_color" binding:"required"`
	SecondaryColor ColorConfigObject `json:"secondary_color" binding:"required"`
	TertiaryColor  ColorConfigObject `json:"tertiary_color" binding:"required"`
	INFOColor      string            `json:"INFO_color" binding:"required"`
	WARNColor      string            `json:"WARN_color" binding:"required"`
	DEBUGColor     string            `json:"DEBUG_color" binding:"required"`
	TRACEColor     string            `json:"TRACE_color" binding:"required"`
	ERRORColor     string            `json:"ERROR_color" binding:"required"`
}

type ColorConfigObject struct {
	Background string `json:"background" binding:"required"`
	Text       string `json:"text" binding:"required"`
}

var Config ConfigObject

var Defaults BrandingConfigObject

func LoadConfig() {
	configPath := os.Getenv(ENV_PREFIX + "CONFIG_PATH")
	if configPath == "" {
		configPath = DEFAULT_CONFIG_PATH
	}

	Defaults = BrandingConfigObject{
		PrimaryColor: ColorConfigObject{
			Background: "#C3C49E",
			Text:       "#000000",
		},
		SecondaryColor: ColorConfigObject{
			Background: "#8F7E4F",
			Text:       "#ffffff",
		},
		TertiaryColor: ColorConfigObject{
			Background: "#524632",
			Text:       "#ffffff",
		},
		INFOColor:  "#4F772D",
		WARNColor:  "#E24E1B",
		DEBUGColor: "#2B50AA",
		TRACEColor: "#610345",
		ERRORColor: "#95190C",
	}

	Config = ConfigObject{
		Host:                    "localhost",
		Port:                    9001,
		TCPPort:                 9002,
		UDPPort:                 9003,
		BasePath:                "",
		CeresDBConnectionString: "http://ceresdb:ceresdb@ceresdb:7437/whitetail",
		CeresDBDatabaseName:     "whitetail",
		Branding:                Defaults,
		LogLevel:                logger.LOG_LEVEL_INFO,
		LogFormat:               logger.LOG_FORMAT_CONSOLE,
		OperationPath:           "/tmp/whitetail/operations/operation.yaml",
	}

	// if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
	// 	configData, _ := json.MarshalIndent(Config, "", " ")

	// 	_ = ioutil.WriteFile(configPath, configData, 0644)
	// }

	// jsonFile, err := os.Open(configPath)
	// if err != nil {
	// 	log.Println("Unable to read json file")
	// 	panic(err)
	// }

	// log.Printf("Successfully Opened %v", configPath)

	// byteValue, _ := ioutil.ReadAll(jsonFile)

	// json.Unmarshal(byteValue, &Config)

	jsonFile, err := os.Open(configPath)
	if err == nil {
		log.Printf("Successfully Opened %v", configPath)

		byteValue, _ := ioutil.ReadAll(jsonFile)

		json.Unmarshal(byteValue, &Config)
	}

	v := reflect.ValueOf(Config)
	t := reflect.TypeOf(Config)

	for i := 0; i < v.NumField(); i++ {
		field, found := t.FieldByName(v.Type().Field(i).Name)
		if !found {
			continue
		}

		value := field.Tag.Get("env")
		if value != "" {
			val, present := os.LookupEnv(ENV_PREFIX + value)
			if present {
				w := reflect.ValueOf(&Config).Elem().FieldByName(t.Field(i).Name)
				x := getAttr(&Config, t.Field(i).Name).Kind().String()
				if w.IsValid() {
					switch x {
					case "int", "int64":
						i, err := strconv.ParseInt(val, 10, 64)
						if err == nil {
							w.SetInt(i)
						}
					case "int8":
						i, err := strconv.ParseInt(val, 10, 8)
						if err == nil {
							w.SetInt(i)
						}
					case "int16":
						i, err := strconv.ParseInt(val, 10, 16)
						if err == nil {
							w.SetInt(i)
						}
					case "int32":
						i, err := strconv.ParseInt(val, 10, 32)
						if err == nil {
							w.SetInt(i)
						}
					case "string":
						w.SetString(val)
					case "float32":
						i, err := strconv.ParseFloat(val, 32)
						if err == nil {
							w.SetFloat(i)
						}
					case "float", "float64":
						i, err := strconv.ParseFloat(val, 64)
						if err == nil {
							w.SetFloat(i)
						}
					case "bool":
						i, err := strconv.ParseBool(val)
						if err == nil {
							w.SetBool(i)
						}
					default:
						objValue := reflect.New(field.Type)
						objInterface := objValue.Interface()
						err := json.Unmarshal([]byte(val), objInterface)
						obj := reflect.ValueOf(objInterface)
						if err == nil {
							w.Set(reflect.Indirect(obj).Convert(field.Type))
						} else {
							log.Println(err)
						}
					}
				}
			}
		}
	}

	// UpdateBranding()
	// InitLogoIcon()
}

func getAttr(obj interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		panic("not struct")
	}
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		panic("not found:" + fieldName)
	}
	return curField
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
