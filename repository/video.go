package repository

import (
	"github.com/RaymondCode/simple-demo/entity"
	"sync"
)

type VideoDao struct {
}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDao() *VideoDao {
	videoOnce.Do(func() {
		videoDao = new(VideoDao)
	})
	return videoDao
}

//将视频信息插入数据库

func (v *VideoDao) VideoInsert(video *entity.Video) error {
	err := DB.Create(&video).Error
	return err
}

//从数据中查询用户发布的视频,得到视频列表

func (v *VideoDao) VideoInfoByName(username string) ([]entity.Video, error) {
	var video []entity.Video
	err := DB.Where(&entity.Video{Author: username}).Find(&video).Error
	return video, err
}

//从数据中查询所有视频,得到视频列表
func (v *VideoDao) VideoInfoAll() ([]entity.Video, error) {
	var video []entity.Video
	err := DB.Find(&video).Error
	return video, err
}
