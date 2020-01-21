package service

import (
	"gocrud/model"
	"gocrud/serializer"
)

type VideoService struct {
}

// create 创建视频
func (service *VideoService) CreateVideo(video model.Video) serializer.Response {
	if err := model.DB.Create(&video).Error; err != nil {
		return serializer.ParamErr("创建视频失败", err)
	}

	return serializer.BuildVideoResponse(video)
}

// List 获取视频
func (service *VideoService) ListVideos() serializer.Response {
	var videos []model.Video
	if err := model.DB.Find(&videos).Error; err != nil {
		return serializer.ParamErr("获取所有视频失败", err)
	}

	return serializer.BuildVideosResponse(videos)
}

// Show 获取视频
func (service *VideoService) ShowVideo(id string) serializer.Response {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.ParamErr("获取视频详情失败", err)
	}

	return serializer.BuildVideoResponse(video)
}

// update 更新视频
func (service *VideoService) UpdateVideo(id string, video model.Video) serializer.Response {
	var videoNew model.Video
	if err := model.DB.First(&videoNew, id).Error; err != nil {
		return serializer.ParamErr("获取视频详情失败", err)
	}
	videoNew.Title = video.Title
	videoNew.Info = video.Info
	if err := model.DB.Save(&videoNew).Error; err != nil {
		return serializer.ParamErr("更新视频失败", err)
	}

	return serializer.BuildVideoResponse(video)
}

// delete 删除视频
func (service *VideoService) DeleteVideo(id string) serializer.Response {
	var video model.Video
	if err := model.DB.First(&video, id).Error; err != nil {
		return serializer.ParamErr("获取视频详情失败", err)
	}

	if err := model.DB.Delete(&video).Error; err != nil {
		return serializer.ParamErr("删除视频失败", err)
	}

	return serializer.Response{}
}
