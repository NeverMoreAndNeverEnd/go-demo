package service

import (
	"gocrud/model"
	"gocrud/serializer"
)

// DeleteVideoService 删除视频的服务
type DeleteVideoService struct {
}

// delete 删除视频
func (service *DeleteVideoService) Delete(id string) serializer.Response {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.ParamErr("获取视频详情失败", err)
	}

	if err := model.DB.Delete(&video).Error; err != nil {
		return serializer.ParamErr("删除视频失败", err)
	}

	return serializer.Response{}
}
