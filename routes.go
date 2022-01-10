package main

import "namgay/jampa/controllers"

func initializeRoutes() {
	router.GET("/campaigns", controllers.GetCampaigns)
	router.GET("/campaigns/:id", controllers.GetCampaign)
	router.POST("/campaigns", controllers.CreateCampaign)
	router.PUT("/campaigns/:id", controllers.UpdateCampaign)
	router.DELETE("/campaigns/:id", controllers.DeleteCampaign)
}
