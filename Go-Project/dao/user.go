package dao

import (
	"errors"

	"github.com/life-studied/douyin-simple/global"
	"gorm.io/gorm"
)

// 查询用户名是否存在
func QueryName(username string) bool {
	//使用gorm查询用户名是否存在
	var users []User
	result := global.DB.Where("name=?", username).First(&users)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			if result.RowsAffected == 0 {
				return true
			}
		} else {
			panic(result.Error)
		}

	}
	return false
}

// 将用户信息存入数据库
func AddUserInfo(id int64, username string, password string) error {

	//将id username password 存入数据库中
	user := User{
		ID:       id,
		Name:     username,
		Password: password,
	}
	tResult := global.DB.Create(&user)
	if tResult.Error != nil {
		return tResult.Error
	}

	return nil
}

// 获取所有数据
func GetAllUsers() ([]User, error) {
	//数据库连接
	var users []User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil

}

//查询Token是否存在
func QueryToken（token string）bool{
	var tokens []Token
	result := global.DB.Where("token=?", token).First(&tokens)
	
	// 检查查询结果和错误
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			if result.RowsAffected == 0 {
				return true
			}
		} else {
			panic(result.Error)
		}
	}
	
	return false
}