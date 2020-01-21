package service

import (
	"gocrud/model"
	"gocrud/serializer"
)

// CreateVideoService 视频投稿的服务
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=1,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
}

// create 创建视频
func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}
	if err := model.DB.Create(&video).Error; err != nil {
		return serializer.ParamErr("创建视频失败", err)
	}

	return serializer.BuildVideoResponse(video)
}


