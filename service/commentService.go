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

	err := repository.NewCommentDao().CommentInsert(c.VideoId, &comment)

	return err
}

//通过videoUid和commentUid来查询评论id

func (c *CommentInfo) CommentInfoByVideoUidAndCommentUid() (entity.Comment, error) {
	comment, err := repository.NewCommentDao().CommentInfoByVideoUidAndCommentUid(c.VideoId, c.UserId)
	return comment, err
}

//通过commentId来删除评论内容

func (c *CommentInfo) DeleteComment() error {
	err := repository.NewCommentDao().DeleteComment(c.VideoId, c.Cid)

	return err
}

//通过videoId来查找评论内容

func (c *CommentInfo) QueryCommentInfoByVideoId() ([]entity.Comment, error) {
	comment, err := repository.NewCommentDao().QueryCommentInfoByVideoId(c.VideoId)
	return comment, err
}
