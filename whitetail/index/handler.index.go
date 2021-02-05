package Index

import (
    "fmt"
    "strings"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "gorm.io/driver/postgres"
    "log"
    "net/url"
    "errors"
	"regexp"
	"whitetail/config"
	"time"
	"strconv"
)

type Index struct {
    IDs string `json:"ids"`
    ID string `json:"id"`
}


var DB *gorm.DB

/* -- Database interactions -- */

func ConnectDataBase(db_type string, postgresConfig *Config.PostgresConfigObject, sqliteConfig *Config.SqliteConfigObject) {
    var err error
    var database *gorm.DB

    if db_type == "postgres" {
        dsn := url.URL{
            User:     url.UserPassword(postgresConfig.Username, postgresConfig.Password),
            Scheme:   db_type,
            Host:     fmt.Sprintf("%s:%d", postgresConfig.Host, postgresConfig.Port),
            Path:     "whitetail",
            RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
        }
    	database, err = gorm.Open(postgres.Open(dsn.String()), &gorm.Config{})
    } else if db_type == "sqlite" {
        database, err = gorm.Open(sqlite.Open(sqliteConfig.Path), &gorm.Config{})
    }

	if err != nil {
		log.Println(err)
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Index{})

    DB = database
	
	GetAllIndices()
}

func DeleteIndexByKey(id string) error{
	index, err := GetIndexByKey(id)
	if err != nil {
		return errors.New("Index")
	}
	DB.Delete(&index)
	return nil
}

func GetAllIndices() []Index{
	var indices []Index
	DB.Find(&indices)
	
	return indices
}


func GetIndexByKey(key string) (*Index, error) {

	var index Index
	if err := DB.Where("id = ?", key).First(&index).Error; err != nil {
		return nil, errors.New("Index not found")
	}
	return &index, nil
}

func CreateNewIndex(key, id string) *Index {
	index := Index{ID: key, IDs: id}

	DB.Create(&index)

	return &index
}

// Update an existsing job with the data provided
func AddIDToIndex(key, newID string){
	index, err := GetIndexByKey(key)
	if err != nil {
		index = CreateNewIndex(key, newID)
	} else {
		newIndex := Index{}
		newIndex.ID = key
		newIndex.IDs = index.IDs + "," + newID
		DB.Model(index).Updates(newIndex)
	}
}

func ParseLog(text, id, timestamp, level, service string) {
	reg, err := regexp.Compile("[^a-zA-Z0-9\\s]+")
    if err != nil {
        log.Fatal(err)
    }
	processedString := reg.ReplaceAllString(text, "")
	words := strings.Split(processedString, " ")
	for _, key := range words {
		AddIDToIndex(key, id)
	}
	layout := "2006-01-02T15:04:05"
	t, err := time.Parse(layout, timestamp)
	yearString := strconv.Itoa(t.Year())
	monthString := strconv.Itoa(int(t.Month()))
	dayString := strconv.Itoa(t.Day())
	hourString := strconv.Itoa(t.Hour())
	minuteString := strconv.Itoa(t.Minute())
	secondString := strconv.Itoa(t.Second())
	AddIDToIndex("@year:" + yearString, id)
	AddIDToIndex("@month:" + monthString, id)
	AddIDToIndex("@day:" + dayString, id)
	AddIDToIndex("@hour:" + hourString, id)
	AddIDToIndex("@minute:" + minuteString, id)
	AddIDToIndex("@second:" + secondString, id)
	AddIDToIndex("@level:" + level, id)
	AddIDToIndex("@service:" + service, id)
	AddIDToIndex("@all", id)
}