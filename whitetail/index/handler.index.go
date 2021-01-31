package Index

import (
    "fmt"
    "strings"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "log"
    "net/url"
    "errors"
	"regexp"
)

type Index struct {
    IDs string `json:"ids"`
    ID string `json:"id"`
}


var DB *gorm.DB

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

func ParseLog(text, id string) {
	reg, err := regexp.Compile("[^a-zA-Z0-9\\s]+")
    if err != nil {
        log.Fatal(err)
    }
	processedString := reg.ReplaceAllString(text, "")
	words := strings.Split(processedString, " ")
	for _, key := range words {
		AddIDToIndex(key, id)
	}
}