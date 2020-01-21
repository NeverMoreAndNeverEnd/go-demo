package api

import (
	"github.com/gin-gonic/gin"
	"gocrud/model"
	"gocrud/service"
)

//CreateVideo 视频投稿接口
func CreateVideo(c *gin.Context) {
	var service service.VideoService
	var video model.Video
	if err := c.ShouldBind(&video); err == nil {
		res := service.CreateVideo(video)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//ShowVideo 视频详情接口
func ShowVideo(c *gin.Context) {
	var service service.VideoService
	res := service.ShowVideo(c.Param("id"))
	c.JSON(200, res)

}

//ListVideo 查询所有视频详情接口
func ListVideo(c *gin.Context) {
	var service service.VideoService
	res := service.ListVideos()
	c.JSON(200, res)
}

//UpdateVideo 更新视频接口
func UpdateVideo(c *gin.Context) {
	var service service.VideoService
	var video model.Video
	if err := c.ShouldBind(&video); err == nil {
		res := service.UpdateVideo(c.Param("id"), video)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

//DeleteVideo 删除视频的接口
func DeleteVideo(c *gin.Context) {
	var service service.VideoService
	res := service.DeleteVideo(c.Param("id"))
	c.JSON(200, res)

}
