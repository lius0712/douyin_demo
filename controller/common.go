package controller

import (
	"fmt"

	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/video"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	Title         string `json:"title"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
}

type Comment struct {
	Id         int64  `json:"id"`
	User       User   `json:"user"`
	Content    string `json:"content"`
	CreateDate string `json:"create_date"`
}

type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow"`
}

func UserByEntity(user entity.User) *User {
	return &User{
		Id:            user.ID,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      user.IsFollow,
	}
}

func VideoByEntity(v entity.Video) (*Video, error) {
	author := service.UserInfo{Username: v.Author}
	user, err := author.UserInfoByName()
	if err != nil {
		return nil, err
	}

	fav := service.FavoriteService{Uid: user.ID, Vid: v.ID}
	return &Video{
		Id:            v.ID,
		Author:        *UserByEntity(user),
		PlayUrl:       video.GetVideoRemotePath(fmt.Sprintf("%d", v.ID)),
		CoverUrl:      video.GetCoverRemotePath(fmt.Sprintf("%d", v.ID)),
		Title:         v.Title,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    fav.UserIsFavorited(),
	}, nil
}
