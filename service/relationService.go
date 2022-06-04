package service

import (
	"fmt"
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
	fmt.Println(relation) //FromUid一直为0， ？？？， bug
	err := repository.DB.Create(&relation).Error

	if err != nil {
		return err
	}

	err = r.UserFollowerCountInc(1)

	if err != nil {
		return err
	}

	err = r.UserFollowerCountInc(1)
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

	err = r.UserFollowerCountInc(-1)

	if err != nil {
		return err
	}

	err = r.UserFollowerCountInc(-1)
	return err
}

func (r *RelationInfo) UserFollowCountInc(count int64) error {
	var user entity.User
	user.ID = r.FromUid
	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Find(&user, &user).Error
		if err != nil {
			return err
		}
		user.FollowCount += count
		err = tx.Save(&user).Error
		return err
	})
	return err
}

func (r *RelationInfo) UserFollowerCountInc(count int64) error {
	var user entity.User
	user.ID = r.ToUid
	err := repository.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Find(&user, &user).Error
		if err != nil {
			return err
		}
		user.FollowCount += count
		err = tx.Save(&user).Error
		return err
	})
	return err
}

// UserIsRelationed checks if the user has relationed the other user.

func (r *RelationInfo) UserIsRelationed() bool {
	var relation entity.Relation
	relation.FromUid = r.FromUid
	relation.ToUid = r.ToUid
	err := repository.DB.Where(&relation).Take(&relation).Error
	return err == nil
}

//查找用户关注列表

func (r *RelationInfo) UserFollowList() ([]entity.User, error) {
	var relation []entity.Relation
	var users []entity.User

	err := repository.DB.Where(&entity.Relation{FromUid: r.FromUid}).Find(&relation).Error

	if err != nil {
		return nil, err
	}

	for _, rel := range relation {
		u := entity.User{ID: rel.ToUid}
		if err := repository.DB.Where(&u).Take(&u).Error; err != nil {
			continue
		}
		users = append(users, u)
	}
	return users, nil
}

//查找用户粉丝列表

func (r *RelationInfo) UserFollowerList() ([]entity.User, error) {
	var relation []entity.Relation
	var users []entity.User

	err := repository.DB.Where(&entity.Relation{ToUid: r.ToUid}).Find(&relation).Error

	if err != nil {
		return nil, err
	}

	for _, rel := range relation {
		u := entity.User{ID: rel.FromUid}
		if err := repository.DB.Where(&u).Take(&u).Error; err != nil {
			continue
		}
		users = append(users, u)
	}
	return users, nil
}
