package main

import (
	"github.com/gin-gonic/gin"
	"hoangvuphone/internal/config"
	"hoangvuphone/internal/render"
)

func main() {

	//Kết nối database
	config.ConnectDatabase()
	render.LoadTemplates()

	router := gin.Default()
	router.LoadHTMLGlob("web/templates/**/*")
	router.Static("/static", "./web/static")

	router.Run(":" + config.GetPort())
}