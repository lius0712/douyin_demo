package repository

import (
	"sync"

	"gorm.io/gorm"

	"github.com/RaymondCode/simple-demo/entity"
)

type FavoriteDao struct {
}

var favoriteDao *FavoriteDao
var favoriteOnce sync.Once

func NewFavoriteDao() *FavoriteDao {
	favoriteOnce.Do(func() {
		favoriteDao = &FavoriteDao{}
	})
	return favoriteDao
}

func (c *FavoriteDao) FavoriteVideo(uid int64, vid int64) error {
	var fav entity.Favorite
	var video entity.Video
	video.ID = vid
	err := DB.Transaction(func(tx *gorm.DB) error {
		fav.Uid = uid
		fav.Vid = vid
		err := DB.Create(&fav).Error

		if err != nil {
			return err
		}

		err = tx.Find(&video, &video).Error
		if err != nil {
			return err
		}

		video.FavoriteCount += 1
		err = tx.Save(&video).Error
		return err

	})

	return err
}

func (c *FavoriteDao) UnFavoriteVideo(uid int64, vid int64) error {
	var fav entity.Favorite
	var video entity.Video
	video.ID = vid
	err := DB.Transaction(func(tx *gorm.DB) error {
		err := DB.Delete(&fav, &fav).Error

		if err != nil {
			return err
		}

		err = tx.Find(&video, &video).Error
		if err != nil {
			return err
		}

		video.FavoriteCount -= 1
		err = tx.Save(&video).Error
		return err

	})

	return err
}

func (c *FavoriteDao) QueryFavoriteInfoByUserId(userId int64) ([]entity.Video, error) {
	var favs []entity.Favorite
	var videos []entity.Video

	err := DB.Where(&entity.Favorite{Uid: userId}).Find(&favs).Error
	if err != nil {
		return nil, err
	}

	for _, fav := range favs {
		v := entity.Video{ID: fav.Vid}
		if err := DB.Where(&v).Take(&v).Error; err != nil {
			continue
		}
		videos = append(videos, v)
	}
	return videos, nil
}

func (c *FavoriteDao) QueryFavoriteInfo(userId int64, videoId int64) (entity.Favorite, error) {
	var fav entity.Favorite
	err := DB.Where(&entity.Favorite{Uid: userId, Vid: videoId}).Take(&fav).Error
	return fav, err
}
