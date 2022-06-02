package controller

import (
	"fmt"
	"net/http"
	"os"

	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/video"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {

	username := c.GetString("username")
	title := c.PostForm("title")

	userByNameSelect := service.UserInfo{
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

	uuid, err := uuid.NewRandom()
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	tmpname := uuid.String()

	tmpPath := video.GetVideoLocalPath(tmpname)

	if err := c.SaveUploadedFile(data, tmpPath); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	videoInsertService := service.VideoInsertInfo{
		Author: user.Name,
		Title:  title,
	}

	vid, err := videoInsertService.VideoInsert()
	if err != nil {
		os.Remove(tmpPath)
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "insert mysql failed",
		})
		return
	}

	realPath := video.GetVideoLocalPath(fmt.Sprintf("%v", vid))
	err = os.Rename(tmpPath, realPath)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	go video.ExtractFrame(realPath, video.GetCoverLocalPath(fmt.Sprintf("%v", vid)))

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  tmpname + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	username := c.GetString("username")
	userId := c.Query("user_id") //user_id 一直为0, 接口问题？
	fmt.Println(userId)
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
		DemoVideo[i].PlayUrl = video.GetVideoRemotePath(fmt.Sprintf("%d", videoList.ID))
		DemoVideo[i].PlayUrl = video.GetCoverRemotePath(fmt.Sprintf("%d", videoList.ID))
		DemoVideo[i].FavoriteCount = videoList.FavoriteCount
		DemoVideo[i].CommentCount = videoList.CommentCount
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
