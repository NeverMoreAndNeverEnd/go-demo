package service

import (
	"gocrud/model"
	"gocrud/serializer"
)

// ShowVideoService 视频详情的服务
type ShowVideoService struct {

}

// Show 获取视频
func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.ParamErr("获取视频详情失败", err)
	}

	return serializer.BuildVideoResponse(video)
}
