package model

import "github.com/jinzhu/gorm"

// Video 视频模型
type Video struct {
	gorm.Model
	Title string `form:"title" json:"title" binding:"required,min=1,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
}