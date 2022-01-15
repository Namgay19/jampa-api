package main

import (
	"fmt"
	"io"
	"namgay/jampa/controllers"
	"namgay/jampa/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var router *gin.Engine

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout);
	router = gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
				param.ClientIP,
				param.TimeStamp.Format(time.RFC1123),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.Request.UserAgent(),
				param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())


	models.ConnectDatabase()
	initializeRoutes()
	router.Static("/public", "./public")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("CampaignDateValidation", controllers.CampaignDateValidation)
	}

	router.Run()
}
