package entity

type Favorite struct {
	Uid int64 `gorm:"column:uid"`
	Vid int64 `gorm:"column:vid"`
}

func (Favorite) TableName() string {
	return "t_favorite"
}
