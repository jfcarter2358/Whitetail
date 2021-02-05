// handler.api.go

package API

import (
	"github.com/gin-gonic/gin"
	"whitetail/logging"
	"whitetail/config"
	"net/http"
	"log"
	"strconv"
	"path/filepath"
	"whitetail/ast"
	"os"
	"io"
	/*
	"github.com/leekchan/timeutil"
	"time"
	*/
)

type AnalyticsInput struct {
	Services []string `json:"services"`
	TimeStart string `json:"time_start"`
	TimeEnd string `json:"time_end"`
}

type QueryInput struct {
	Query string `json:"query"`
}

func QueryLogs(c *gin.Context) {
	var input QueryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Unable to bind JSON")
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ids := AST.Parse(input.Query)
	
	logMessages := []string{}
	for _, id := range(ids) {
		log, err := Logging.GetLogByID(id)
		if err == nil {
			logMessages = append(logMessages, log.Text)
		}
	}

	c.JSON(http.StatusOK, gin.H{"logs": logMessages})
}

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
		logs = Logging.GetLogsByService(input.KeywordList, service, lineLimit, input.LogLevels)
	} else {
		logs = Logging.GetLogsByService(input.KeywordList, service, 1000, input.LogLevels)
	}

	logMessages := []string{}
	for _, log := range logs {
		logMessages = append(logMessages, log.Text)
	}

	c.JSON(http.StatusOK, gin.H{"logs": logMessages})
}

func UpdateLogo(c *gin.Context) {
	file, err := c.FormFile("file")

    // The file cannot be received.
    if err != nil {
		log.Println(err)
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "message": "No file is received",
        })
        return
    }

    // Retrieve file information
    extension := filepath.Ext(file.Filename)
	if extension != ".png" {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "message": "File is of wrong format",
        })
        return
    }

    // The file is received, so let's save it
    if err := c.SaveUploadedFile(file, "static/img/logo.png"); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "message": "Unable to save the file",
        })
        return
    }
	Config.UpdateBranding()
	
	c.Redirect(302, Config.Config.BasePath + "/ui/settings")
}

func UpdateIcon(c *gin.Context) {
	file, err := c.FormFile("file")

    // The file cannot be received.
    if err != nil {
		log.Println(err)
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "message": "No file is received",
        })
        return
    }

    // Retrieve file information
    extension := filepath.Ext(file.Filename)
	if extension != ".png" {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "message": "File is of wrong format",
        })
        return
    }

    // The file is received, so let's save it
    if err := c.SaveUploadedFile(file, "static/img/favicon.png"); err != nil {
        c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
            "message": "Unable to save the file",
        })
        return
    }
	Config.UpdateBranding()
	
	c.Redirect(302, Config.Config.BasePath + "/ui/settings")
}

func UpdateColors(c *gin.Context) {
	var input Config.BrandingConfigObject
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Unable to bind JSON")
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Config.Config.Branding.PrimaryColor = input.PrimaryColor
	Config.Config.Branding.SecondaryColor = input.SecondaryColor
	Config.Config.Branding.TertiaryColor = input.TertiaryColor
	Config.Config.Branding.INFOColor = input.INFOColor
	Config.Config.Branding.WARNColor = input.WARNColor
	Config.Config.Branding.DEBUGColor = input.DEBUGColor
	Config.Config.Branding.TRACEColor = input.TRACEColor
	Config.Config.Branding.ERRORColor = input.ERRORColor
	Config.UpdateBranding()
	c.JSON(http.StatusOK, gin.H{})
}

func DefaultColors(c *gin.Context) {
	Config.Config.Branding.PrimaryColor = Config.Defaults.PrimaryColor
	Config.Config.Branding.SecondaryColor = Config.Defaults.SecondaryColor
	Config.Config.Branding.TertiaryColor = Config.Defaults.TertiaryColor
	Config.Config.Branding.INFOColor = Config.Defaults.INFOColor
	Config.Config.Branding.WARNColor = Config.Defaults.WARNColor
	Config.Config.Branding.DEBUGColor = Config.Defaults.DEBUGColor
	Config.Config.Branding.TRACEColor = Config.Defaults.TRACEColor
	Config.Config.Branding.ERRORColor = Config.Defaults.ERRORColor

	Config.UpdateBranding()
	c.JSON(http.StatusOK, gin.H{})
}

func DefaultLogo(c *gin.Context) {
	log.Println("Copying default logo file")
	source, err := os.Open("static/img/logo.default.png")
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
	c.JSON(http.StatusOK, gin.H{})
}

func DefaultIcon(c *gin.Context) {
	log.Println("Copying default icon file")
	source, err := os.Open("static/img/favicon.default.png")
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
	c.JSON(http.StatusOK, gin.H{})
}

func getLogs(c *gin.Context) {
	increments, _ := strconv.Atoi(c.Param(":increments"))
	/*
	base := time.Now()
    td := timeutil.Timedelta{Days: 0, Minutes: 1, Seconds: 0}
	counts := []int{}
	labels := []string{}
	*/

	i := 0
	for i < increments {
		// count := 
	}
	// result := base.Add(td.Duration())
	
    // fmt.Println(result) // "2015-02-13 00:17:56 +0000 UTC"
}