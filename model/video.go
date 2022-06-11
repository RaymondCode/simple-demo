package model

type Video struct {
	VideoID    uint `gorm:"primarykey"`
	UserID     uint
	Title      string `gorm:"not null"`
	OSSVideoID string `gorm:"not null"`
	CreatedAt  int
	Comments   []*Comment
	Likes      []*User `gorm:"many2many:like;joinForeignKey:video_id;joinReferences:user_id;"`
}

func (v *Video) TableName() string {
	return "video"
}
