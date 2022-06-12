package service

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/repository"
)

type FavoriteService struct {
	Uid int64
	Vid int64
}

// FavoriteAction adds the video to the user's favorite list and increments the video's favorite count by 1.
func (f *FavoriteService) FavorateAction() error {
	fdao := repository.NewFavoriteDao()
	vdao := repository.NewVideoDao()

	err := fdao.FavoriteVideo(f.Uid, f.Vid)
	if err != nil {
		return err
	}
	err = vdao.VideoFavoriteInc(f.Vid, 1)
	return err
}

// UnFavoriteAction removes the video from the user's favorite list and decrements the video's favorite count by 1
func (f *FavoriteService) UnFavorateAction() error {
	fdao := repository.NewFavoriteDao()
	vdao := repository.NewVideoDao()

	err := fdao.UnFavoriteVideo(f.Uid, f.Vid)
	if err != nil {
		return err
	}
	err = vdao.VideoFavoriteInc(f.Vid, -1)
	return err
}

// UserIsFavorited checks if the user has favorited the video.
func (f *FavoriteService) UserIsFavorited() bool {
	fdao := repository.NewFavoriteDao()
	_, err := fdao.QueryFavoriteInfo(f.Uid, f.Vid)
	return err != nil
}

// UserFavoritedVideos returns a list of videos favorited by the user.
func (f *FavoriteService) UserFavoriteVideos() ([]entity.Video, error) {
	dao := repository.NewFavoriteDao()
	infos, err := dao.QueryFavoriteInfoByUserId(f.Uid)
	if err != nil {
		return nil, err
	}

	return infos, nil
}
