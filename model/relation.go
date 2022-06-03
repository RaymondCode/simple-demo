package model

//每个用户都有的一个关注表
type Follower struct {
	Id       int64 `json:"id" gorm:"primaryKey"`
	UserId   int64 `json:"user_id"`
	ToUserId int64 `json:"to_user_id"`
}

func (Follower) TableName() string {
	return "user_follower"
}
