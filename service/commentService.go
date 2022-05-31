package service

import (
	"github.com/RaymondCode/simple-demo/entity"
	"github.com/RaymondCode/simple-demo/repository"
)

type CommentInsertInfo struct {
	UserId     int64
	VideoId    int64
	Comment    string
	CreateDate string
}

//将评论插入数据库
func (c *CommentInsertInfo) CommentInsert() error {
	var comment entity.Comment
	comment.Vid = c.VideoId
	comment.Uid = c.UserId
	comment.Comment = c.Comment
	comment.CreateDate = c.CreateDate

	err := repository.DB.Create(&comment).Error

	return err
}
