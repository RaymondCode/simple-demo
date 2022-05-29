package repository

import (
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
)

type UserRepository struct{}

func (ur *UserRepository) QueryUserById(id int64) (*model.User, error) {
	var user model.User
	err := global.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) QueryUserByNameAndPassword(username, password string) (*model.User, error) {
	var user model.User
	err := global.DB.Where("username = ? and password = ?", username, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// IfNameExist 检查用户名是不是用过了。用过了返回1，没用过返回0。int具有比boolean更好的扩展性。
func (ur *UserRepository) IfNameExist(name string) int64 {
	var cnt int64
	global.DB.Model(model.User{}).Where("name = ?", name).Count(&cnt)
	return cnt
}

// SaveUser 把新创建的User存到数据库中
func (ur *UserRepository) SaveUser(user *model.User) int64 {
	global.DB.Save(user)
	return user.Id
}
