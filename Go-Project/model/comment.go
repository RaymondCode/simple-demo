package model

type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true"`
	UserId     int64  `json:"-"`
	VideoId    int64  `json:"-"`
	User       User   `json:"user" gorm:"foreignKey:user_id;references:id;"`
	Content    string `json:"content,omitempty"`
	CreateDate int64  `json:"create_date,omitempty"`

}
