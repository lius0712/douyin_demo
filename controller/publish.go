package controller

import (
	"fmt"
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	username := c.GetString("username")

	userByNameSelect := service.UserInfo{ //这里token为用户名，根据用户名进行查找，后续进行优化
		Username: username,
	}
	user, errUser := userByNameSelect.UserInfoByName()

	if errUser != nil { //如果不存在该用户名
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)

	finalName := fmt.Sprintf("%d_%s", user.ID, filename)

	saveFile := filepath.Join("./public/", finalName)

	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	IP := "http://10.128.229.176:8080/static/" //改为app连接服务端的地址，否则无法播放，后续优化
	playerUrl := IP + finalName                //视频源
	coverUrl := IP + "2.jpg"                   //封面图片，后续优化：截取视频某一帧作为封面

	videoInsertService := service.VideoInsertInfo{
		Author:    user.Name,
		PlayerUrl: playerUrl,
		CoverUrl:  coverUrl,
		Title:     "Test video",
	}

	err = videoInsertService.VideoInsert()
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "insert mysql failed",
		})
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	username := c.GetString("username")
	userQuery := service.UserInfo{
		Username: username,
	}
	user, errUser := userQuery.UserInfoByName() //查询到该用户信息
	if errUser != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "token已过期！",
			},
		})
		return
	}

	var videos []entity.Video
	var err error
	videoListByName := service.VideoListByName{
		Name: user.Name,
	}
	videos, err = videoListByName.VideoInfoByName()
	if err != nil {
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  "查询出错！",
			},
		})
		return
	}

	lenVideo := len(videos)

	var DemoVideo []Video
	var VideoUser User
	DemoVideo = make([]Video, lenVideo)

	VideoUser.Id = user.ID //类型转换
	VideoUser.Name = user.Name
	VideoUser.FollowCount = user.FollowCount
	VideoUser.FollowerCount = user.FollowerCount
	VideoUser.IsFollow = user.IsFollow

	for i, videoList := range videos {
		DemoVideo[i].Author = VideoUser
		DemoVideo[i].PlayUrl = videoList.PlayerUrl
		DemoVideo[i].CoverUrl = videoList.CoverUrl
		DemoVideo[i].FavoriteCount = videoList.FavoriteCount
		DemoVideo[i].CommentCount = videoList.CommentCount
		DemoVideo[i].IsFavorite = videoList.IsFavorite
	}
	//fmt.Println("************")
	//fmt.Println(DemoVideo)

	c.JSON(
		http.StatusOK,
		VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: DemoVideo,
		},
	)
}
