// main.go

package main

import (
	// "os"
	"github.com/gin-gonic/gin"
	"log"
	"whitetail/logging"
	"whitetail/config"
	"whitetail/index"
	"whitetail/page"
	"strconv"
)

var router *gin.Engine
var config *Config.ConfigObject

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	config = Config.ReadConfigFile()
	routerPort := ":" + strconv.Itoa(config.HTTPPort)

	Page.InitConfig(config)

	// Read in the compass data from the json file
	Logging.ConnectDataBase(config.Database.Type, config.Database.Postgres, config.Database.Sqlite)
	Index.ConnectDataBase(config.Database.Type, config.Database.Postgres, config.Database.Sqlite)
	basePath := config.BasePath
	log.Print("Running with base path: " + basePath)

	go Logging.StartTCPServer(config.TCPPort)
	go Logging.StartUDPServer(config.UDPPort)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes(basePath)

	// Start serving the application
	router.Run(routerPort)
}