package dao

import (
	"errors"
	"github.com/life-studied/douyin-simple/model"

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

// 查询用户名和密码
func GetUserByUsernameAndPassword(username, password string) (User, error) {
	var users []User
	result := global.DB.Where("name = ? AND password = ?", username, password).First(&users)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return User{}, errors.New("User not found")
		}
		return User{}, result.Error
	}
	return User{}, nil
}

func QueryUserById(id int64) (*model.User, error) {
	var user model.User
	err := global.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
