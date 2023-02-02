package models

import (
	"fmt"
	"go-scristi/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Post struct {
	gorm.Model
	Name    string `json:"name"`
	Author  string `json:"author"`
	Date    string `json:"date"`
	Content string `json:"content"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	err := db.AutoMigrate(&Post{})
	if err != nil {
		fmt.Println("Error! fail at Automigrating Post model")
	}
}

func (p *Post) CreatePost() *Post {
	db.Create(&p)
	return p
}

func ReturnLastPosts(numberOf int) []Post {
	var posts []Post
	db.Model(&Post{}).Limit(numberOf).Find(posts)
	return posts
}
