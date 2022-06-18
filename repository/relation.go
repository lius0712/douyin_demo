package repository

import (
	"sync"

	"github.com/RaymondCode/simple-demo/entity"
	"gorm.io/gorm"
)

type RelationDao struct {
}

var relationDao *RelationDao
var relationOnce sync.Once

func NewRelationDao() *RelationDao {
	relationOnce.Do(func() {
		relationDao = new(RelationDao)
	})
	return relationDao
}

func (r *RelationDao) FollowAction(from, to int64) error {
	var user entity.User
	err := DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&entity.Relation{FromUid: from, ToUid: to}).Error
		if err != nil {
			return err
		}
		err = tx.Find(&user, &entity.User{ID: from}).Error
		if err != nil {
			return err
		}
		user.FollowCount += 1
		err = tx.Save(&user).Error
		if err != nil {
			return err
		}

		err = tx.Find(&user, &entity.User{ID: to}).Error
		if err != nil {
			return err
		}
		user.FollowerCount += 1
		err = tx.Save(&user).Error
		return err
	})
	return err
}

func (r *RelationDao) UnFollowAction(from, to int64) error {
	var user entity.User
	err := DB.Transaction(func(tx *gorm.DB) error {
		var relation entity.Relation
		relation.FromUid = from
		relation.ToUid = to
		err := tx.Delete(&relation, &relation).Error
		if err != nil {
			return err
		}
		err = tx.Find(&user, &entity.User{ID: from}).Error
		if err != nil {
			return err
		}
		user.FollowCount -= 1
		err = tx.Save(&user).Error
		if err != nil {
			return err
		}

		err = tx.Find(&user, &entity.User{ID: to}).Error
		if err != nil {
			return err
		}
		user.FollowerCount -= 1
		err = tx.Save(&user).Error
		return err
	})
	return err
}

func (r *RelationDao) UserIsRelationed(fromUid int64, toUid int64) error {
	var relation entity.Relation
	relation.FromUid = fromUid
	relation.ToUid = toUid
	err := DB.Where(&relation).Take(&relation).Error
	return err
}

//查找用户关注列表

func (r *RelationDao) UserFollowList(fromUid int64) ([]entity.User, error) {
	var relation []entity.Relation
	var users []entity.User

	err := DB.Where(&entity.Relation{FromUid: fromUid}).Find(&relation).Error

	if err != nil {
		return nil, err
	}

	for _, rel := range relation {
		u := entity.User{ID: rel.ToUid}
		if err := DB.Where(&u).Take(&u).Error; err != nil {
			continue
		}
		users = append(users, u)
	}
	return users, nil
}

//查找用户粉丝列表

func (r *RelationDao) UserFollowerList(toUid int64) ([]entity.User, error) {
	var relation []entity.Relation
	var users []entity.User

	err := DB.Where(&entity.Relation{ToUid: toUid}).Find(&relation).Error

	if err != nil {
		return nil, err
	}

	for _, rel := range relation {
		u := entity.User{ID: rel.FromUid}
		if err := DB.Where(&u).Take(&u).Error; err != nil {
			continue
		}
		users = append(users, u)
	}
	return users, nil
}

//关注人数增减count

func (r *RelationDao) UserFollowCountInc(uid int64, count int64) error {
	var user entity.User
	user.ID = uid
	err := DB.Transaction(func(tx *gorm.DB) error {
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

//粉丝数增减count

func (r *RelationDao) UserFollowerCountInc(uid int64, count int64) error {
	var user entity.User
	user.ID = uid
	err := DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Find(&user, &user).Error
		if err != nil {
			return err
		}
		user.FollowerCount += count
		err = tx.Save(&user).Error
		return err
	})
	return err
}
