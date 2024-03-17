package controllers

import (
	"models"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	var contents []models.Content
	models.DB.Find(&contents)

	fmt.Println(contents[0].TITLE)

	c.HTML(http.StatusOK, "public/index.tmpl", gin.H{
		"title": "Main website",
		"datas": contents,
	})
}
