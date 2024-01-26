// main.go

package main

import (
	// "os"
	"log"
	"whitetail/config"
	"whitetail/datastore"
	"whitetail/logger"
	"whitetail/operation"
	"whitetail/probe"

	"github.com/gin-gonic/gin"

	"strconv"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	config.LoadConfig()
	logger.SetLevel(config.Config.LogLevel)
	logger.SetFormat(config.Config.LogFormat)
	routerPort := ":" + strconv.Itoa(config.Config.Port)

	// Read in the compass data from the json file

	datastore.Init()

	log.Print("Running with port: " + strconv.Itoa(config.Config.Port))

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	operation.LoadOperation()
	logger.Debugf("", "loaded operation: %v", operation.Operations)
	for oName, o := range operation.Operations.Observers {
		for sName := range o.Streams {
			if err := datastore.AddStream(oName, sName, o); err != nil {
				logger.Fatalf("", "Cannot add stream: %s__%s", oName, sName)
			}
		}
	}

	// Initialize the routes
	initializeRoutes()

	logger.Info("", "Creating probes")
	for pName, p := range operation.Operations.Probes {
		logger.Debugf("", "Checking probe %s", pName)
		probeArgs := make(map[string]map[string]map[string]interface{})
		for oName, o := range operation.Operations.Observers {
			logger.Debugf("", "Checking observer %s", oName)
			for sName, s := range o.Streams {
				logger.Debugf("", "Checking stream %s", sName)
				if s.Probe == pName {
					logger.Infof("Associating probe %s with stream %s/%s", pName, oName, sName)
					if _, ok := probeArgs[oName]; !ok {
						probeArgs[oName] = make(map[string]map[string]interface{})
					}
					probeArgs[oName][sName] = s.Arguments
				}
			}
		}
		logger.Tracef("", "Probe args: %v", probeArgs)
		go probe.Run(p, probeArgs)
	}

	// Start serving the application
	router.Run(routerPort)
}
