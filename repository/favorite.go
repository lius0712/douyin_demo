package repository

import (
	"gorm.io/gorm"
	"sync"

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
	err := DB.Where(&entity.Favorite{Uid: uid, Vid: vid}).Find(&fav).Error
	if err == nil || err == gorm.ErrRecordNotFound {
		fav.Uid = uid
		fav.Vid = vid
		err = DB.Create(&fav).Error
	}
	return err
}

func (c *FavoriteDao) UnFavoriteVideo(uid int64, vid int64) error {
	var fav entity.Favorite
	fav.Uid = uid
	fav.Vid = vid
	err := DB.Delete(&fav, &fav).Error
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
