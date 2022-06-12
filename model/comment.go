package model

type Comment struct {
	CommentID uint `gorm:"primarykey"`
	UserID    uint
	VideoID   uint
	Content   string `gorm:"not null"`
	CreatedAt int
}

func (c *Comment) TableName() string {
	return "comment"
}
