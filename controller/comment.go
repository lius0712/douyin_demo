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
	username := c.GetString("username") //用户鉴权token
	//fmt.Println(c.Query("user_id")) // 接口中无user_id
	//userId, err1 := strconv.ParseInt(c.Query("user_id"), 10, 64) //将字符类型转化成整形
	//if err1 != nil {
	//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "数据转化错误"})
	//	return
	//}
	videoId, err2 := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err2 != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "数据转化错误"})
		return
	}
	actionType := c.Query("action_type") //发布评论1， 删除评论2

	userInfoByName := service.UserInfo{
		Username: username,
	}

	entityUser, err := userInfoByName.UserInfoByName()
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	if actionType == "1" { //发布评论
		text := c.Query("comment_text") //评论内容
		now := time.Now()
		commentInsertInfo := service.CommentInfo{
			UserId:     entityUser.ID,
			VideoId:    videoId,
			Comment:    text,
			CreateDate: now.Format("2006-01-02 15:04"),
		}
		errInsert := commentInsertInfo.CommentInsert()
		if errInsert != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Comment insert failed!"})
			return
		}

		commentQueryInfo := service.CommentInfo{
			UserId:  entityUser.ID,
			VideoId: videoId,
		}

		entityComment, errQuery := commentQueryInfo.CommentInfoByVideoUidAndCommentUid()
		if errQuery != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "Comment query failed",
			})
		}

		var user User
		user.Id = entityUser.ID
		user.Name = entityUser.Name
		user.FollowCount = entityUser.FollowCount
		user.FollowerCount = entityUser.FollowerCount
		user.IsFollow = entityUser.IsFollow

		c.JSON(http.StatusOK, CommentActionResponse{
			Comment: Comment{
				Id:         entityComment.ID,
				User:       user,
				Content:    entityComment.Comment,
				CreateDate: entityComment.CreateDate,
			},
		})
	} else { //删除评论
		commentId, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "数据转化错误"})
			return
		}
		commentInfoDelete := service.CommentInfo{Cid: commentId, VideoId: videoId}
		errDelete := commentInfoDelete.DeleteCommentByCid()

		if errDelete != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  "delete comment failed!",
			})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {

	videoId, errData := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if errData != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "数据转化错误"})
		return
	}

	queryCommentInfoByVid := service.CommentInfo{VideoId: videoId}
	entityComment, err := queryCommentInfoByVid.QueryCommentInfoByVideoId()

	if err != nil {
		c.JSON(http.StatusOK, CommentListResponse{Response: Response{
			StatusCode: 1,
			StatusMsg:  "comment query failed",
		}})
		return
	}

	lenComment := len(entityComment)

	var DemoComment []Comment
	var CommentUser []User

	DemoComment = make([]Comment, lenComment)
	CommentUser = make([]User, lenComment)

	for i, commentItem := range entityComment {
		commentUid := commentItem.Uid //获取每条评论的用户id
		queryUserInfoByUid := service.UserInfo{Uid: commentUid}
		userEntity, errUser := queryUserInfoByUid.UserInfoByUid()
		if errUser != nil {
			c.JSON(http.StatusOK, Response{
				StatusCode: 1,
				StatusMsg:  "用户查询失败！",
			})
			return
		}
		CommentUser[i].Id = userEntity.ID
		CommentUser[i].Name = userEntity.Name
		CommentUser[i].FollowCount = userEntity.FollowCount
		CommentUser[i].FollowerCount = userEntity.FollowerCount
		CommentUser[i].IsFollow = userEntity.IsFollow

		DemoComment[i].Id = commentItem.ID
		DemoComment[i].User = CommentUser[i]
		DemoComment[i].Content = commentItem.Comment
		DemoComment[i].CreateDate = commentItem.CreateDate
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: DemoComment,
	})
}
