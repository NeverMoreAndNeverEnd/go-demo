package service

import (
	"gocrud/model"
	"gocrud/serializer"
)

// ListVideoService 视频列表的服务
type ListVideoService struct {
}

// List 获取视频
func (service *ListVideoService) List() serializer.Response {
	var videos []model.Video
	if err := model.DB.Find(&videos).Error; err != nil {
		return serializer.ParamErr("获取所有视频失败", err)
	}

	return serializer.BuildVideosResponse(videos)
}
