package models

import (
	"time"
)

type Content struct {
	Id      uint `gorm:"primary_key"`
	Title   string
	Content string
	Author  string
	Type    string

	CreatedBy string
	CreatedAt time.Time

	UpdatedBy string
	UpdatedAt time.Time
}

func InitContent() {
	var count int64
	DB.Model(&Content{}).Where("Type = ? ", "HOME").Count(&count)
	if count == 0 {
		home := Content{Title: "Home", Content: "Just a website to fill my free time, this website was made from go gin with the MVC concept. I'll add some cool features sometime, if I have time. You can see my YouTube and GitHub, there are several source codes that can be used. Thank you for visiting", Author: "Brian", Type: "HOME", CreatedBy: "Brian"}
		DB.Create(&home)
	}
}

func GetContentByType(t string) Content {
	content := Content{}
	DB.First(&content, "TYPE = ? ", t)
	return content
}
