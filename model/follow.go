package model

type Follow struct {
	UserID uint `gorm:"column:user_id"`
	FanID  uint `gorm:"column:fan_id"`
}

func (f *Follow) TableName() string {
	return "follow"
}
