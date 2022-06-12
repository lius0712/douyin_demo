package service

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/repository"
)

type VideoInsertInfo struct {
	Author     string
	PlayerUrl  string
	CoverUrl   string
	Title      string
	CreateDate string
}

type VideoListByName struct {
	Name string
}

//视频信息插入数据库
func (v *VideoInsertInfo) VideoInsert() (int64, error) {
	var video entity.Video
	video.Author = v.Author
	video.Title = v.Title
	video.CreateDate = v.CreateDate

	err := repository.NewVideoDao().VideoInsert(&video)
	return video.ID, err
}

//从数据中查询用户发布的视频,得到视频列表
func (v *VideoListByName) VideoInfoByName() ([]entity.Video, error) {
	video, err := repository.NewVideoDao().VideoInfoByName(v.Name)
	return video, err
}

//从数据中查询所有视频,得到视频列表
func VideoInfoAll() ([]entity.Video, error) {
	video, err := repository.NewVideoDao().VideoInfoAll()
	return video, err
}
