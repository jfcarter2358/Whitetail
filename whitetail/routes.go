// routes.go

package main

import (
	"whitetail/api"
	"whitetail/middleware"
	"whitetail/page"
)

func initializeRoutes() {
	router.Static("/static/css", "./static/css")
	router.Static("/static/img", "./static/img")
	router.Static("/static/js", "./static/js")

	router.GET("/", page.RedirectIndexPage)

	healthRoutes := router.Group("/health", middleware.CORSMiddleware())
	{
		healthRoutes.GET("/healthy", api.GetHealth)
	}

	apiRoutes := router.Group("/api", middleware.CORSMiddleware())
	{
		v1Routes := apiRoutes.Group("/v1")
		{
			baseStationRoutes := v1Routes.Group("/basestation")
			{
				baseStationRoutes.POST("/:observer/:stream", api.UpdateBaseStation)
				baseStationRoutes.GET("/:observer/:stream", api.GetBaseStation)
			}
			operationRoutes := v1Routes.Group("/operation")
			{
				operationRoutes.GET("/operation", api.GetOperation)
				operationRoutes.POST("/operation", api.UpdateOperation)
			}
			v1Routes.GET("/status", api.GetStatus)
			v1Routes.POST("/query", api.Query)
		}
	}

	// settingsRoutes := router.Group("/settings", middleware.CORSMiddleware())
	// {
	// 	colorRoutes := settingsRoutes.Group("/colors")
	// 	{
	// 		colorRoutes.POST("/settings/colors/update", api.UpdateColors)
	// 		colorRoutes.POST("/settings/colors/default", api.DefaultColors)
	// 	}
	// 	logoRoutes := settingsRoutes.Group("/logo")
	// 	{
	// 		logoRoutes.POST("/settings/logo/update", api.UpdateLogo)
	// 		logoRoutes.POST("/settings/logo/default", api.DefaultLogo)
	// 	}
	// 	iconRoutes := settingsRoutes.Group("/icon")
	// 	{
	// 		iconRoutes.POST("/settings/icon/update", api.UpdateIcon)
	// 		iconRoutes.POST("/settings/icon/default", api.DefaultIcon)
	// 	}
	// }

	uiRoutes := router.Group("/ui", middleware.CORSMiddleware())
	{
		uiRoutes.GET("/home", page.ShowHomePage)
		uiRoutes.GET("/query", page.ShowQueryPage)
		// uiRoutes.GET("/settings", page.ShowSettingsPage)
		dashboardRoutes := uiRoutes.Group("/dashboard")
		{
			dashboardRoutes.GET("/:name", page.ShowDashboardPage)
		}
	}
}
