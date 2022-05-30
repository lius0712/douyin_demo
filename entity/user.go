package entity

type User struct {
	ID            int64  `gorm:"column:uid"`
	Name          string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	FollowCount   int64  `gorm:"column:follow_count"`
	FollowerCount int64  `gorm:"column:follower_count"`
}

func (User) TableName() string {
	return "t_user"
}
