package entity

type Relation struct {
	FromUid int64 `gorm:"column:from_uid"`
	ToUid   int64 `gorm:"column:to_uid"`
}

func (Relation) TableName() string {
	return "t_relation"
}
