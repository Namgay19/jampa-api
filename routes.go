package main

import "namgay/jampa/controllers"

func initializeRoutes() {
	api := router.Group("/api")
	{
	api.GET("/campaigns", controllers.GetCampaigns)
	api.GET("/campaigns/:id", controllers.GetCampaign)
	api.POST("/campaigns", authenticateUser(), controllers.CreateCampaign)
	api.PUT("/campaigns/:id", controllers.UpdateCampaign)
	api.DELETE("/campaigns/:id", controllers.DeleteCampaign)

	api.POST("/campaigns/:id/donate", controllers.CreateDonation)
	api.GET("/campaigns/:id/donations", controllers.GetDonations)

	api.POST("/users", controllers.RegisterUser)
	api.POST("/users/login", controllers.Login)
	}
}
