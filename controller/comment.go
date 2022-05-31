package controller

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	userId, err1 := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err1 != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "数据转化错误"})
		return
	}
	videoId, err2 := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err2 != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "数据转化错误"})
		return
	}
	actionType := c.Query("action_type")

	userInfoByName := service.UserInfo{
		Username: token,
	}

	_, err := userInfoByName.UserInfoByName()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	if actionType == "1" { //发布评论
		text := c.Query("comment_text")
		commentInsertInfo := service.CommentInsertInfo{
			UserId:     userId,
			VideoId:    videoId,
			Comment:    text,
			CreateDate: time.Now().String(),
		}
		errInsert := commentInsertInfo.CommentInsert()
		if errInsert != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Comment insert failed!"})
			return
		}

		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}

	//if user, exist := usersLoginInfo[token]; exist {
	//	if actionType == "1" {
	//		text := c.Query("comment_text")
	//		c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
	//			Comment: Comment{
	//				Id:         1,
	//				User:       user,
	//				Content:    text,
	//				CreateDate: "05-01",
	//			}})
	//		return
	//	}
	//	c.JSON(http.StatusOK, Response{StatusCode: 0})
	//} else {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	//}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
