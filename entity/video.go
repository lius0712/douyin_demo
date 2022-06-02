package entity

type Video struct {
	ID            int64  `gorm:"column:vid"`
	Author        string `gorm:"column:author"`
	Title         string `gorm:"column:title"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	IsFavorite    bool   `gorm:"column:is_favorite"`
}

func (Video) TableName() string {
	return "t_video"
}
