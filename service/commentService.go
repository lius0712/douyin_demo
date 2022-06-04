package service

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/repository"
	"gorm.io/gorm"
)

type CommentInfo struct {
	Cid        int64
	UserId     int64
	VideoId    int64
	Comment    string
	CreateDate string
}

//将评论插入数据库

func (c *CommentInfo) CommentInsert() error {
	var comment entity.Comment
	comment.Vid = c.VideoId
	comment.Uid = c.UserId
	comment.Comment = c.Comment
	comment.CreateDate = c.CreateDate

	err := repository.DB.Create(&comment).Error

	if err != nil {
		return err
	}

	err = c.videoCommentInc(1)
	return err
}

//通过videoUid和commentUid来查询评论id

func (c *CommentInfo) CommentInfoByVideoUidAndCommentUid() (entity.Comment, error) {
	var comment entity.Comment
	err := repository.DB.Where(&entity.Comment{Uid: c.UserId, Vid: c.VideoId}).First(&comment).Error
	return comment, err
}

//通过commentId来删除评论内容

func (c *CommentInfo) DeleteCommentByCid() error {
	var comment entity.Comment
	comment.ID = c.Cid
	err := repository.DB.Delete(&comment, &comment).Error

	if err != nil {
		return err
	}

	err = c.videoCommentInc(-1)

	return err
}

//通过videoId来查找评论内容

func (c *CommentInfo) QueryCommentInfoByVideoId() ([]entity.Comment, error) {
	var comment []entity.Comment
	err := repository.DB.Where(&entity.Comment{Vid: c.VideoId}).Find(&comment).Error
	return comment, err
}

// videoCommentInc Increments the comment count of the video by 1.
func (c *CommentInfo) videoCommentInc(count int64) error {
	var video entity.Video
	video.ID = c.VideoId
	err := repository.DB.Transaction(func(tx *gorm.DB) error {
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
