package controllers

import (
	"models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	var content models.Content
	content = models.GetContentByType("HOME")

	c.HTML(http.StatusOK, "public/index.tmpl", gin.H{
		"title":   content.TITLE,
		"content": content.CONTENT,
	})
}
