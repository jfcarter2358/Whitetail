// routes.go

package main

import (
	"whitetail/api"
	"whitetail/page"
)

func initializeRoutes(basePath string) {
	router.Static(basePath + "/resources/css", "./static/css")
	router.Static(basePath + "/resources/img", "./static/img")
	router.Static(basePath + "/resources/js", "./static/js")

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
