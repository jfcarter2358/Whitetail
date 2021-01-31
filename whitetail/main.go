// main.go

package main

import (
	// "os"
	"github.com/gin-gonic/gin"
	"log"
	"whitetail/logging"
	"whitetail/config"
	"whitetail/index"
	"strconv"
)

var router *gin.Engine
var config *Config.ConfigObject

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	config = Config.ReadConfigFile()
	routerPort := ":" + strconv.Itoa(config.HTTPPort)

	db_type := config.Database.Type
	db_host := config.Database.Host
	db_port := strconv.Itoa(config.Database.Port)
	db_user := config.Database.Username
	db_pass := config.Database.Password

	// Read in the compass data from the json file
	Logging.ConnectDataBase(db_type, db_user, db_pass, db_host, db_port)
	Index.ConnectDataBase(db_type, db_user, db_pass, db_host, db_port)
	log.Print("Running on port: " + routerPort)

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