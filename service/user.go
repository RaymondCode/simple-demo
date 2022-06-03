package service

import (
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
)

type UserService struct{}

func (us *UserService) QueryUser(userId int64, token string) (*model.User, error) {
	rawUser, err := repository.GroupApp.UserRepository.QueryUserById(userId)
	if err != nil {
		return nil, err
	}
	return rawUser, nil
}

// QueryUserByNameAndPassword 根据用户名和密码查询用户，返回用户id。
func (us *UserService) QueryUserByNameAndPassword(username, password string) (int64, error) {
	var user model.User
	err := global.DB.Where("name = ? and password = ?", username, password).First(&user).Error
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

// IfNameExist 检查用户名是不是用过了。用过了返回1，没用过返回0。int具有比boolean更好的扩展性。
func (us *UserService) IfNameExist(name string) int64 {
	return repository.GroupApp.UserRepository.IfNameExist(name)
}

// SaveUser 把新创建的User存到数据库中
func (us *UserService) SaveUser(username, password, nickname string) int64 {
	user := model.User{Name: username, Password: password, Nickname: nickname}
	return repository.GroupApp.UserRepository.SaveUser(&user)
}
