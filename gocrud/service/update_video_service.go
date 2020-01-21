package service

import (
	"gocrud/model"
	"gocrud/serializer"
)

// UpdateVideoService 更新视频的服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=1,max=30"`
	Info  string `form:"info" json:"info" binding:"max=300"`
}

// update 更新视频
func (service *UpdateVideoService) Update(id string) serializer.Response {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.ParamErr("获取视频详情失败", err)
	}

	video.Title = service.Title
	video.Info = service.Info
	if err := model.DB.Save(&video).Error; err != nil {
		return serializer.ParamErr("更新视频失败", err)
	}

	return serializer.BuildVideoResponse(video)
}
