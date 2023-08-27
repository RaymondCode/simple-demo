package dao

import (
	"log"

	"github.com/life-studied/douyin-simple/global"
)

func GetUserByUserId(id int64) (User, error) {
	user := User{}
	if err := global.DB.Where("id = ?", id).First(&user).Error; err != nil {
		log.Println(err.Error())
		return user, err
	}
	return user, nil
}
