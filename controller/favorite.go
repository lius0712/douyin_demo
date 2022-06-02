package controller

import (
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	uid := c.GetInt64("uid")
	vid, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "video_id is not valid",
		})
		return
	}

	t, err := strconv.ParseInt(c.Query("action_type"), 10, 64)
	if err != nil || t < 1 || t > 2 {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "action_type is not valid",
		})
		return
	}

	f := service.FavoriteService{}
	f.Uid = uid
	f.Vid = vid

	if t == 1 {
		err = f.FavorateAction()
	} else {
		err = f.UnFavorateAction()
	}

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "Operation Failed",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "Operation Success",
	})
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	//c.JSON(http.StatusOK, VideoListResponse{
	//	Response: Response{
	//		StatusCode: 0,
	//	},
	//	VideoList: DemoVideos,
	//})
}
