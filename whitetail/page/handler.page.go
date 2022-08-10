// pages.go

package Page

import (
	// "log"
	"net/http"
	"whitetail/config"
	"whitetail/logging"

	"github.com/gin-gonic/gin"
)

func RedirectIndexPage(c *gin.Context) {
	c.Redirect(301, config.Config.BasePath+"/ui/home")
}

func ShowHomePage(c *gin.Context) {
	// Render the logs-selection.html page
	render(c, gin.H{
		"title":    "Home",
		"basePath": config.Config.BasePath,
		"location": "Home"},
		"page.index.html")
}

func ShowLogsPage(c *gin.Context) {
	// Render the logs-selection.html page
	render(c, gin.H{
		"title":    "Logs",
		"basePath": config.Config.BasePath,
		"location": "Logs",
		"services": logging.Services,
		"db_name":  config.Config.DB.Name},
		"page.logs.html")
}

func ShowQueryPage(c *gin.Context) {
	// Render the logs-selection.html page
	render(c, gin.H{
		"title":    "Logs",
		"basePath": config.Config.BasePath,
		"location": "Query",
		"db_name":  config.Config.DB.Name},
		"page.query.html")
}

func ShowSettingsPage(c *gin.Context) {
	// Render the logs-selection.html page
	render(c, gin.H{
		"title":           "Settings",
		"basePath":        config.Config.BasePath,
		"location":        "Settings",
		"primary_color":   config.Config.Branding.PrimaryColor,
		"secondary_color": config.Config.Branding.SecondaryColor,
		"tertiary_color":  config.Config.Branding.TertiaryColor,
		"INFO_color":      config.Config.Branding.INFOColor,
		"WARN_color":      config.Config.Branding.WARNColor,
		"DEBUG_color":     config.Config.Branding.DEBUGColor,
		"TRACE_color":     config.Config.Branding.TRACEColor,
		"ERROR_color":     config.Config.Branding.ERRORColor},
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
