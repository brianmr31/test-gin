package controllers

import (
	"models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	var content models.Content
	content = models.GetContentByType("HOME")
	youtube, logo, github, email := models.GetParamHome()
	c.HTML(http.StatusOK, "public/index.tmpl", gin.H{
		"title":   content.Title,
		"content": content.Content,
		"youtube": youtube,
		"logo":    logo,
		"github":  github,
		"email":   email,
	})
}
