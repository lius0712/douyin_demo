package controller

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
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
	allVideos := make([]Video, lenVideo)
	allAuthor := make([]User, lenVideo)
	var errUserItem error
	var UserItem entity.User
	for i, videoItem := range videos {
		userItemName := videoItem.Author //获取视频作者用户名
		userItemInfoByName := service.UserInfo{ //根据作者用户名查找用户信息
			Username: userItemName,
		}
		UserItem, errUserItem = userItemInfoByName.UserInfoByName()
		if errUserItem != nil {
			c.JSON(http.StatusOK, FeedResponse{
				Response: Response{StatusCode: 1, StatusMsg: "查找失败！"},
			})
			return
		}
		allAuthor[i].Id = UserItem.ID
		allAuthor[i].Name = UserItem.Name
		allAuthor[i].FollowCount = UserItem.FollowCount
		allAuthor[i].FollowerCount = UserItem.FollowerCount
		allAuthor[i].IsFollow = UserItem.IsFollow

		allVideos[i].Id = videoItem.ID
		allVideos[i].Author = allAuthor[i]
		allVideos[i].PlayUrl = videoItem.PlayerUrl
		allVideos[i].CoverUrl = videoItem.CoverUrl
		allVideos[i].FavoriteCount = videoItem.FavoriteCount
		allVideos[i].CommentCount = videoItem.CommentCount
		allVideos[i].Title = videoItem.Title
		allVideos[i].IsFavorite = videoItem.IsFavorite

	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: allVideos,
		NextTime:  time.Now().Unix(),
	})
}
