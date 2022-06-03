package service

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/repository"
	"gorm.io/gorm"
)

type FavoriteService struct {
	Uid int64
	Vid int64
}

// FavoriteAction adds the video to the user's favorite list and increments the video's favorite count by 1.
func (f *FavoriteService) FavorateAction() error {
	var fav entity.Favorite
	fav.Uid = f.Uid
	fav.Vid = f.Vid
	err := repository.DB.Create(&fav).Error

	if err != nil {
		return err
	}

	err = f.videoFavoriteInc(1)
	return err
}

// UnFavoriteAction removes the video from the user's favorite list and decrements the video's favorite count by 1
func (f *FavoriteService) UnFavorateAction() error {
	var fav entity.Favorite
	fav.Uid = f.Uid
	fav.Vid = f.Vid
	err := repository.DB.Delete(&fav, &fav).Error

	if err != nil {
		return err
	}

	err = f.videoFavoriteInc(-1)
	return err
}

// UserIsFavorited checks if the user has favorited the video.
func (f *FavoriteService) UserIsFavorited() bool {
	var fav entity.Favorite
	fav.Uid = f.Uid
	fav.Vid = f.Vid
	err := repository.DB.Where(&fav).Take(&fav).Error
	return err == nil
}

// UserFavoritedVideos returns a list of videos favorited by the user.
func (f *FavoriteService) UserFavoriteVideos() ([]entity.Video, error) {
	var favs []entity.Favorite
	var videos []entity.Video

	err := repository.DB.Where(&entity.Favorite{Uid: f.Uid}).Find(&favs).Error
	if err != nil {
		return nil, err
	}

	for _, fav := range favs {
		v := entity.Video{ID: fav.Vid}
		if err := repository.DB.Where(&v).Take(&v).Error; err != nil {
			continue
		}
		videos = append(videos, v)
	}
	return videos, nil
}

// videoFavoriteInc Increments the favorite count of the video by 1.
func (f *FavoriteService) videoFavoriteInc(count int64) error {
	var video entity.Video
	video.ID = f.Vid
	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Find(&video, &video).Error
		if err != nil {
			return err
		}
		video.FavoriteCount += count
		err = tx.Save(&video).Error
		return err
	})
	return err
}
