package models

import "time"

type Content struct {
	ID      uint64 `json:"id" gorm:"primary_key"`
	TITLE   string `json:"title"`
	CONTENT string `json:"content"`
	AUTHOR  string `json:"author"`
	TYPE    string `json:"type"`

	CREATED_BY   string    `json:"created_by"`
	CREATED_DATE time.Time `jsno:"created_date"`

	MODIFIED_BY   string    `json:"created_by"`
	MODIFIED_DATE time.Time `jsno:"created_date"`
}
