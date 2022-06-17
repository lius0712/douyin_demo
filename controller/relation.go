package controller

import (
	"net/http"
	"strconv"

	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	FromUid := c.GetInt64("uid")
	ToUid, err := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "to_user_id is not valid"})
	}

	r := service.RelationInfo{FromUid: FromUid, ToUid: ToUid}

	actionType := c.Query("action_type") //关注1， 取消关注2

	if actionType == "1" {
		err = r.FollowAction()
	} else {
		err = r.UnFollowAction()
	}

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "Operation Failed",
		})
		return
	}

	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "Operation Success"})
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {

	Uid, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "数据转化错误"})
		return
	}

	relationInfo := service.RelationInfo{FromUid: Uid}
	users, err := relationInfo.UserFollowList()

	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}

	userList := make([]User, 0, len(users))

	for _, u := range users {
		uItem := UserByEntity(u)
		relation := service.RelationInfo{FromUid: Uid, ToUid: uItem.Id}
		uItem.IsFollow = relation.UserIsRelationed()
		userList = append(userList, uItem)
	}

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userList,
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	Uid, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "数据转化错误"})
		return
	}

	relationInfo := service.RelationInfo{ToUid: Uid}
	users, err := relationInfo.UserFollowerList()

	if err != nil {
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
	}

	userList := make([]User, 0, len(users))

	for _, u := range users {
		uItem := UserByEntity(u)
		relation := service.RelationInfo{FromUid: Uid, ToUid: uItem.Id}
		uItem.IsFollow = relation.UserIsRelationed()
		userList = append(userList, uItem)
	}

	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: userList,
	})
}
