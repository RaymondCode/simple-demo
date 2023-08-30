package model

type User struct {
	Id       int64  `json:"id,omitempty" gorm:"primaryKey"`
	Name     string `json:"name,omitempty"`
	Password string `json:"-" `
}
