// pages.go

package Page

import (
	// "log"
	"github.com/gin-gonic/gin"
	"net/http"
	"whitetail/logging"
	"whitetail/config"
)

var config *Config.ConfigObject

func InitConfig(newConfig *Config.ConfigObject) {
	config = newConfig
}

func RedirectIndexPage(c *gin.Context) {
	// log.Println(config.BasePath)
	c.Redirect(301, config.BasePath + "/ui/home")
}

func ShowHomePage(c *gin.Context) {
	// Render the logs-selection.html page
	render(c, gin.H{
		"title":   "Home",
		"basePath": config.BasePath,
		"location": "Home"}, 
		"page.index.html")
}

func ShowLogsPage(c *gin.Context) {
	// Render the logs-selection.html page
	render(c, gin.H{
		"title":   "Logs",
		"basePath": config.BasePath,
		"location": "Logs",
		"services": Logging.Services}, 
		"page.logs.html")
}

func ShowSettingsPage(c *gin.Context) {
	// Render the logs-selection.html page
	render(c, gin.H{
		"title":   "Settings",
		"basePath": config.BasePath,
		"location": "Settings"}, 
		"page.settings.html")
}

func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
