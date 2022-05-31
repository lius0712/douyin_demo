package entity

type Comment struct {
	ID         int64  `gorm:"column:cid"`
	Vid        int64  `gorm:"column:video_uid"`
	Uid        int64  `gorm:"column:comment_uid"`
	Comment    string `gorm:"column:comment"`
	CreateDate string `gorm:"column:create_date"`
}

func (Comment) TableName() string {
	return "t_comment"
}
