package service

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/repository"
)

type RelationInfo struct {
	FromUid int64
	ToUid   int64
}

//adds the relation to the user's relation list and increments the from_user's follow_count, And to_user's follower_count .

func (r *RelationInfo) FollowAction() error {
	err := repository.NewRelationDao().FollowAction(r.FromUid, r.ToUid)
	return err
}

func (r *RelationInfo) UnFollowAction() error {
	err := repository.NewRelationDao().UnFollowAction(r.FromUid, r.ToUid)
	return err
}

// UserIsRelationed checks if the user has relationed the other user.

func (r *RelationInfo) UserIsRelationed() bool {
	err := repository.NewRelationDao().UserIsRelationed(r.FromUid, r.ToUid)
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
