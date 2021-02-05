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
		apiRoutes.POST("/logs/service/:service", API.GetLogsByService)
		apiRoutes.POST("/logs/query", API.QueryLogs)
		apiRoutes.POST("/settings/colors/update", API.UpdateColors)
		apiRoutes.POST("/settings/logo/update", API.UpdateLogo)
		apiRoutes.POST("/settings/icon/update", API.UpdateIcon)
		apiRoutes.POST("/settings/colors/default", API.DefaultColors)
		apiRoutes.POST("/settings/logo/default", API.DefaultLogo)
		apiRoutes.POST("/settings/icon/default", API.DefaultIcon)
	}

	uiRoutes := router.Group(basePath + "/ui")
	{
		uiRoutes.GET("/home", Page.ShowHomePage)
		uiRoutes.GET("/logs", Page.ShowLogsPage)
		uiRoutes.GET("/query", Page.ShowQueryPage)
		uiRoutes.GET("/settings", Page.ShowSettingsPage)
	}
}
