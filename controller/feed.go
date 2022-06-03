package controller

import (
	"net/http"
	"time"

	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	var videos []entity.Video
	var err error
	videos, err = service.VideoInfoAll()
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{StatusCode: 1, StatusMsg: "查找失败！"},
		})
		return
	}

	lenVideo := len(videos)
	allVideos := make([]Video, 0, lenVideo)
	uid := c.GetInt64("uid")
	for _, videoItem := range videos {
		vid := VideoByEntity(videoItem)
		author, err := VideoAuthor(videoItem)
		if err != nil {
			continue
		}
		vid.Author = *author
		if uid > 0 {
			fav := service.FavoriteService{Uid: uid, Vid: vid.Id}
			vid.IsFavorite = fav.UserIsFavorited()
		}
		allVideos = append(allVideos, vid)
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: allVideos,
		NextTime:  time.Now().Unix(),
	})
}
