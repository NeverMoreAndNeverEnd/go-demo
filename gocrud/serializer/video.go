package serializer

import "gocrud/model"

// Video 视频序列化器
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}

// BuildVideo 序列化视频
func BuildVideo(video model.Video) Video {
	return Video{
		ID:        video.ID,
		Title:     video.Title,
		Info:      video.Info,
		CreatedAt: video.CreatedAt.Unix(),
	}
}

// BuildVideoResponse 序列化视频响应
func BuildVideoResponse(video model.Video) Response {
	return Response{
		Data: BuildVideo(video),
	}
}

// BuildVideosResponse 序列化视频列表响应
func BuildVideosResponse(items []model.Video) Response {
	var videos []Video
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}

	return Response{
		Data: videos,
	}
}
