// handler.api.go

package API

import (
	"github.com/gin-gonic/gin"
	"whitetail/logging"
	"net/http"
	"log"
	"strconv"
)

func GetLogsByService(c *gin.Context) {
	service := c.Param("service")
	var input Logging.LogRequestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Unable to bind JSON")
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logs := []Logging.Log{}
	lineLimit, err := strconv.Atoi(input.LineLimit)
	
	if err == nil {
		logs = Logging.GetLogsByService(input.KeywordList, service, lineLimit)
	} else {
		logs = Logging.GetLogsByService(input.KeywordList, service, 1000)
	}

	logMessages := []string{}
	for _, log := range logs {
		logMessages = append(logMessages, log.Text)
	}

	c.JSON(http.StatusOK, gin.H{"logs": logMessages})
}