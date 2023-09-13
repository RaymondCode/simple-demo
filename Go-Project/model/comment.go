package model

type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true"`
	UserId     int64  `json:"-"`
	VideoId    int64  `json:"-"`
	User       User   `json:"user,omitempty" gorm:"foreignKey:user_id;references:id;"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
	ID         uint
}

type Video struct {
	Video_id      int64  `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY;not null" json:"video_id"`
	Author_id     int64  `gorm:"column:author_id;type:int;not null" json:"author_id"`
	Play_url      string `gorm:"column:play_url;type:varchar(255);not null" json:"play_url"`
	Cover_url     string `gorm:"column:name;type:varchar(255);not null" json:"cover_url"`
	Comment_count int64  `gorm:"column:comment_count;type:int;default:0" json:"comment_count"`
	Title         string `gorm:"column:title;type:varchar(255);not null" json:"title"`
	Publish_time  int64  `gorm:"column:publish_time;type:bigint;not null" json:"publish_time"`
}
