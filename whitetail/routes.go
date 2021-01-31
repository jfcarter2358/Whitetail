// routes.go

package main

import (
	"whitetail/api"
	"whitetail/Page"
)

func initializeRoutes(basePath string) {
	router.Static("/resources/css", "./static/css")
	router.Static("/resources/img", "./static/img")
	router.Static("/resources/js", "./static/js")

	router.GET(basePath + "/", Page.RedirectIndexPage)

	apiRoutes := router.Group(basePath + "/api")
	{
		apiRoutes.POST("/logs/:service", API.GetLogsByService)
	}

	uiRoutes := router.Group(basePath + "/ui")
	{
		uiRoutes.GET("/home", Page.ShowHomePage)
		uiRoutes.GET("/logs", Page.ShowLogsPage)
		uiRoutes.GET("/settings", Page.ShowSettingsPage)
	}
}
