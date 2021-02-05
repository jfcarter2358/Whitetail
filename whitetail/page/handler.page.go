// pages.go

package Page

import (
	// "log"
	"github.com/gin-gonic/gin"
	"net/http"
	"whitetail/logging"
	"whitetail/config"
)

func RedirectIndexPage(c *gin.Context) {
	c.Redirect(301, Config.Config.BasePath + "/ui/home")
}

func ShowHomePage(c *gin.Context) {
	// Render the logs-selection.html page
	render(c, gin.H{
		"title":   "Home",
		"basePath": Config.Config.BasePath,
		"location": "Home"}, 
		"page.index.html")
}

func ShowLogsPage(c *gin.Context) {
	// Render the logs-selection.html page
	render(c, gin.H{
		"title":   "Logs",
		"basePath": Config.Config.BasePath,
		"location": "Logs",
		"services": Logging.Services}, 
		"page.logs.html")
}

func ShowSettingsPage(c *gin.Context) {
	// Render the logs-selection.html page
	render(c, gin.H{
		"title":   "Settings",
		"basePath": Config.Config.BasePath,
		"location": "Settings",
		"primary_color": Config.Config.Branding.PrimaryColor,
		"secondary_color": Config.Config.Branding.SecondaryColor,
		"tertiary_color": Config.Config.Branding.TertiaryColor,
		"INFO_color": Config.Config.Branding.INFOColor,
		"WARN_color": Config.Config.Branding.WARNColor,
		"DEBUG_color": Config.Config.Branding.DEBUGColor,
		"TRACE_color": Config.Config.Branding.TRACEColor,
		"ERROR_color": Config.Config.Branding.ERRORColor}, 
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
