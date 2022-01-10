package main

import (
	"namgay/jampa/models"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()

	models.ConnectDatabase()
	initializeRoutes()

	router.Run()
}

