package service

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/repository"
	"gorm.io/gorm"
)

type RelationInfo struct {
	FromUid int64
	ToUid   int64
}

//adds the relation to the user's relation list and increments the from_user's follow_count, And to_user's follower_count .

func (r *RelationInfo) RelationAction() error {
	var relation entity.Relation
	relation.FromUid = r.FromUid
	relation.ToUid = r.ToUid
	err := repository.NewRelationDao().RelationAction(&relation)

	if err != nil {
		return err
	}

	err = repository.NewRelationDao().UserFollowerCountInc(r.ToUid, 1)
	//err = r.UserFollowerCountInc(1)

	if err != nil {
		return err
	}

	err = repository.NewRelationDao().UserFollowCountInc(r.FromUid, 1)
	//err = r.UserFollowCountInc(1)
	return err
}

func (r *RelationInfo) UnRelationAction() error {
	var relation entity.Relation
	relation.FromUid = r.FromUid
	relation.ToUid = r.ToUid
	err := repository.DB.Delete(&relation, &relation).Error

	if err != nil {
		return err
	}

	err = repository.NewRelationDao().UserFollowerCountInc(r.ToUid, -1)
	//err = r.UserFollowerCountInc(-1)

	if err != nil {
		return err
	}

	err = repository.NewRelationDao().UserFollowCountInc(r.FromUid, -1)
	//err = r.UserFollowerCountInc(-1)
	return err
}

// UserIsRelationed checks if the user has relationed the other user.

func (r *RelationInfo) UserIsRelationed() bool {
	err := repository.NewRelationDao().UserIsRelationed(r.FromUid, r.ToUid)
	if err == gorm.ErrRecordNotFound {
		return false
	}
	return err == nil
}

//查找用户关注列表

func (r *RelationInfo) UserFollowList() ([]entity.User, error) {
	users, err := repository.NewRelationDao().UserFollowList(r.FromUid)
	return users, err
}

//查找用户粉丝列表

func (r *RelationInfo) UserFollowerList() ([]entity.User, error) {
	users, err := repository.NewRelationDao().UserFollowerList(r.ToUid)
	return users, err
}
