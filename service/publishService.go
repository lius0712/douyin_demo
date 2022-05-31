package service

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/repository"
)

<<<<<<< HEAD
=======

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/repository"
)

type VideoInsertInfo struct {
	Author    string
	PlayerUrl string
	CoverUrl  string
	Title     string
}

type VideoListByName struct {
	Name string
}

//视频信息插入数据库
func (v *VideoInsertInfo) VideoInsert() error {
	var video entity.Video
	video.Author = v.Author
	video.PlayerUrl = v.PlayerUrl
	video.CoverUrl = v.CoverUrl
	video.Title = v.Title

	err := repository.DB.Create(&video).Error
	return err
}

//从数据中查询用户发布的视频,得到视频列表
func (v *VideoListByName) VideoInfoByName() ([]entity.Video, error) {
	var video []entity.Video
	err := repository.DB.Where(&entity.Video{Author: v.Name}).Find(&video).Error
	return video, err
}

//从数据中查询所有视频,得到视频列表
func VideoInfoAll() ([]entity.Video, error) {
	var video []entity.Video
	err := repository.DB.Find(&video).Error
	return video, err
}
>>>>>>> dev_ls
