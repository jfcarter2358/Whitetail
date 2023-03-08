// main.go

package main

import (
	// "os"
	"log"
	"whitetail/ceresdb"
	"whitetail/config"
	"whitetail/logging"

	"github.com/gin-gonic/gin"
	"github.com/jfcarter2358/ceresdb-go/connection"

	// "whitetail/index"
	// "whitetail/ast"

	"strconv"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	config.LoadConfig()
	routerPort := ":" + strconv.Itoa(config.Config.HTTPPort)

	// Read in the compass data from the json file
	// Logging.ConnectDataBase(Config.Config.Database.Type, Config.Config.Database.Postgres, Config.Config.Database.Sqlite)
	// Index.ConnectDataBase(Config.Config.Database.Type, Config.Config.Database.Postgres, Config.Config.Database.Sqlite)

	connection.Initialize(config.Config.DB.Username, config.Config.DB.Password, config.Config.DB.Host, config.Config.DB.Port)

	if err := ceresdb.VerifyDatabase(config.Config.DB.Name); err != nil {
		panic(err)
	}
	if err := ceresdb.VerifyCollections(config.Config.DB.Name); err != nil {
		panic(err)
	}

	basePath := config.Config.BasePath
	log.Print("Running with base path: " + basePath)
	log.Print("Running with port: " + strconv.Itoa(config.Config.HTTPPort))

	logging.InitLogger()

	go logging.StartTCPServer(config.Config.TCPPort)
	go logging.StartUDPServer(config.Config.UDPPort)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes(basePath)

	// Kick off the log cleanup check
	go logging.Cleanup()

	// Start serving the application
	router.Run(routerPort)
}
