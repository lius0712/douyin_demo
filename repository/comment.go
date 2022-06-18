package repository

import (
	"sync"

	"github.com/RaymondCode/simple-demo/entity"
	"gorm.io/gorm"
)

type CommentDao struct {
}

var commentDao *CommentDao
var commentOnce sync.Once

func NewCommentDao() *CommentDao {
	commentOnce.Do(func() {
		commentDao = new(CommentDao)
	})
	return commentDao
}

//将评论插入数据库

func (c *CommentDao) CommentInsert(videoId int64, comment *entity.Comment) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		var video entity.Video
		video.ID = videoId
		err := DB.Create(&comment).Error
		if err != nil {
			return err
		}

		err = tx.Find(&video, &video).Error
		if err != nil {
			return err
		}
		video.CommentCount += 1
		err = tx.Save(&video).Error
		return err
	})
	return err
}

//通过videoUid和commentUid来查询评论id

func (c *CommentDao) CommentInfoByVideoUidAndCommentUid(videoUid int64, commentUid int64) (entity.Comment, error) {
	var comment entity.Comment
	err := DB.Where(&entity.Comment{Uid: commentUid, Vid: videoUid}).First(&comment).Error
	return comment, err
}

//通过commentId来删除评论内容

func (c *CommentDao) DeleteComment(vid int64, cid int64) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		var comment entity.Comment
		comment.ID = cid
		err := DB.Delete(&comment, &comment).Error

		if err != nil {
			return err
		}

		var video entity.Video
		video.ID = vid
		err = tx.Find(&video, &video).Error
		if err != nil {
			return err
		}

		video.CommentCount -= 1
		err = tx.Save(&video).Error
		return err
	})
	return err
}

//通过videoId来查找评论内容

func (c *CommentDao) QueryCommentInfoByVideoId(videoId int64) ([]entity.Comment, error) {
	var comment []entity.Comment
	err := DB.Where(&entity.Comment{Vid: videoId}).Order("create_date DESC").Find(&comment).Error
	return comment, err
}

// videoCommentInc Increments the comment count of the video by 1.

func (c *CommentDao) VideoCommentInc(count int64, videoId int64) error {
	var video entity.Video
	video.ID = videoId
	err := DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Find(&video, &video).Error
		if err != nil {
			return err
		}
		video.CommentCount += count
		err = tx.Save(&video).Error
		return err
	})
	return err
}
