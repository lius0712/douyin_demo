package service

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/repository"
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
	err := repository.DB.Delete(&entity.Comment{ID: c.Cid}).Error
	return err
}
