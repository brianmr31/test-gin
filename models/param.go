package models

import (
	"time"
)

type Param struct {
	id    uint `gorm:"primary_key"`
	Code  string
	Value string

	Type string

	CreatedBy string
	CreatedAt time.Time

	UpdatedBy string
	UpdatedAt time.Time
}

var searchs = []string{"url_youtube", "url_logo", "url_github", "url_email"}

func InitParam() {
	var params [4]Param
	datas := []Param{}
	count := int64(0)

	var isYoutubeExits bool = true
	var isLogoExits bool = true
	var isGithubExits bool = true
	var isEmailExits bool = true

	DB.Where("code in ? ", searchs).Find(&datas)

	for i := 0; i < len(searchs); i++ {
		for j := 0; j < len(datas); j++ {
			if searchs[i] == datas[j].Code {
				if searchs[i] == "url_youtube" {
					isYoutubeExits = false
					break
				} else if searchs[i] == "url_logo" {
					isLogoExits = false
					break
				} else if searchs[i] == "url_github" {
					isGithubExits = false
					break
				} else if searchs[i] == "url_email" {
					isEmailExits = false
					break
				}

			}
		}
	}

	if isYoutubeExits {
		params[count] = Param{Code: "url_youtube", Value: "https://www.youtube.com/@whyme-vr7xx", Type: "SYSTEM", CreatedBy: "Brian"}
		count += 1
	}

	if isLogoExits {
		params[count] = Param{Code: "url_logo", Value: "/assets/image/logo.png", Type: "SYSTEM", CreatedBy: "Brian"}
		count += 1
	}

	if isGithubExits {
		params[count] = Param{Code: "url_github", Value: "https://github.com/whymecek", Type: "SYSTEM", CreatedBy: "Brian"}
		count += 1
	}

	if isEmailExits {
		params[count] = Param{Code: "url_email", Value: "why.me.cek@gmail.com", Type: "SYSTEM", CreatedBy: "Brian"}
		count += 1
	}

	if count > 0 {
		DB.Create(&params)
	}
}

func GetParamHome() (string, string, string, string) {
	datas := []Param{}

	var youtube string = ""
	var logo string = ""
	var github string = ""
	var email string = ""

	DB.Where("code in ? ", searchs).Find(&datas)

	for i := 0; i < len(searchs); i++ {
		for j := 0; j < len(datas); j++ {
			if searchs[i] == datas[j].Code {
				if searchs[i] == "url_youtube" {
					youtube = datas[j].Value
					break
				} else if searchs[i] == "url_logo" {
					logo = datas[j].Value
					break
				} else if searchs[i] == "url_github" {
					github = datas[j].Value
					break
				} else if searchs[i] == "url_email" {
					email = datas[j].Value
					break
				}

			}
		}
	}

	return youtube, logo, github, email
}
