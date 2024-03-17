package main

import (
	"controllers"
	"models"
	"net/http"

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

	//load html file
	r.LoadHTMLGlob("templates/**/*.html")

	//static path
	r.Use(static.Serve("/assets", static.LocalFile("./assets", false)))

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/index")
	})
	r.GET("/index", controllers.Home)

	r.Run(":8080")

}
