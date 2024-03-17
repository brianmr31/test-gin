package main

import (
	"controllers"
	"models"

	"fmt"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Start Server")

	// Connect to database
	models.ConnectDatabase()

	r := gin.Default()
	r.Use(gin.Logger())
	r.Delims("{{", "}}")
	// r.LoadHTMLGlob("templates/*.tmpl")

	//load html file
	r.LoadHTMLGlob("templates/**/*.tmpl")

	//static path
	r.Use(static.Serve("/assets", static.LocalFile("./assets", false)))

	r.GET("/index", controllers.Home)
	r.GET("/", controllers.Home)
	r.Run(":8080")

}
