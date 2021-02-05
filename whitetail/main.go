// main.go

package main

import (
	// "os"
	"github.com/gin-gonic/gin"
	"log"
	"whitetail/logging"
	"whitetail/config"
	"whitetail/index"
	// "whitetail/ast"
	"strconv"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	Config.ReadConfigFile()
	routerPort := ":" + strconv.Itoa(Config.Config.HTTPPort)

	// Read in the compass data from the json file
	Logging.ConnectDataBase(Config.Config.Database.Type, Config.Config.Database.Postgres, Config.Config.Database.Sqlite)
	Index.ConnectDataBase(Config.Config.Database.Type, Config.Config.Database.Postgres, Config.Config.Database.Sqlite)
	basePath := Config.Config.BasePath
	log.Print("Running with base path: " + basePath)

	go Logging.StartTCPServer(Config.Config.TCPPort)
	go Logging.StartUDPServer(Config.Config.UDPPort)

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