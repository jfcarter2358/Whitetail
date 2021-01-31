// pages.go

package Page

import (
	"os"
	"github.com/gin-gonic/gin"
	"net/http"
	"whitetail/logging"
)

func RedirectIndexPage(c *gin.Context) {
	basePath := os.Getenv("BASEPATH")
	c.Redirect(301, basePath + "/ui/home")
}

func ShowHomePage(c *gin.Context) {
	basePath := os.Getenv("BASEPATH")
	// Render the logs-selection.html page
	render(c, gin.H{
		"title":   "Home",
		"basePath": basePath,
		"location": "Home"}, 
		"page.index.html")
}

func ShowLogsPage(c *gin.Context) {
	basePath := os.Getenv("BASEPATH")
	// Render the logs-selection.html page
	render(c, gin.H{
		"title":   "Logs",
		"basePath": basePath,
		"location": "Logs",
		"services": Logging.Services}, 
		"page.logs.html")
}

func ShowSettingsPage(c *gin.Context) {
	basePath := os.Getenv("BASEPATH")
	// Render the logs-selection.html page
	render(c, gin.H{
		"title":   "Settings",
		"basePath": basePath,
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
