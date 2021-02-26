package Ceres

import (
    "bytes"
    "encoding/json"
    // "fmt"
    "log"
    "net/http"
)

type LogResponse struct {
    Status int `json:"status"`
	Data   []LogResponseDatum `json:"data"`
}

type LogResponseDatum struct {
	Year      int    `json:"year"`
	Month     int    `json:"month"`
	Day       int    `json:"day"`
	Hour      int    `json:"hour"`
	Minute    int    `json:"minute"`
	Second    int    `json:"second"`
	Service   string `json:"service"`
	Message   string `json:"message"`
	Level     string `json:"level"`
	Timestamp string `json:"timestamp"`
	ID        string `json:"id"`
}

var CeresHost string

func InitConfig(ceresHost string) {
	CeresHost = ceresHost
}

func Query(query string) []LogResponseDatum {
	values := map[string]string{"query": query}
    json_data, err := json.Marshal(values)

    if err != nil {
        log.Fatal(err)
    }

    resp, err := http.Post(CeresHost + "/query", "application/json", bytes.NewBuffer(json_data))

    if err != nil {
        log.Fatal(err)
    }

    var data LogResponse

    json.NewDecoder(resp.Body).Decode(&data)

    return data.Data
}

func Insert(data []map[string]interface{}) {
	values := make(map[string]interface{})
	values["messages"] = data
    json_data, err := json.Marshal(values)

    if err != nil {
        log.Fatal(err)
    }

    _, err = http.Post(CeresHost + "/insert", "application/json", bytes.NewBuffer(json_data))

    if err != nil {
        log.Fatal(err)
    }

    // var res map[string]interface{}

    // json.NewDecoder(resp.Body).Decode(&res)

    // return res["data"]
}