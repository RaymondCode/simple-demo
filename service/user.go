package service

import (
	"errors"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/config"
	"github.com/warthecatalyst/douyin/dao"
	"github.com/warthecatalyst/douyin/erromsg"
	"github.com/warthecatalyst/douyin/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserExist(userId int64) (bool, error) {
	user, err := dao.NewUserDaoInstance().GetUserById(userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return user != nil, nil
}


// UserInfo 获取用户信息
// 包括 user_id,name,follow_count,follower_count,is_favorite
// 最后一个字段应该和具体业务有关（我暂时还不太理解）
func UserInfo(id int64) (api.UserInfo, error) {
	// 处理错误
	handleError := func(errorType *erromsg.Eros) api.UserInfo {
		return api.UserInfo{
			Response: api.Response{StatusCode: errorType.Code, StatusMsg: errorType.Message},
		}
	}

	// 获取到 user_id,follow_count,follower_count
	u, err := dao.NewUserDaoInstance().GetUserById(id)
	if err != nil {
		return handleError(erromsg.ErrQueryUserInfoFail), err
	}

	// 获取 user_name
	name, err := getUserName(id)
	if err != nil {
		return handleError(erromsg.ErrQueryUserNameFail), err
	}
	return api.UserInfo{
		Response: api.Response{StatusCode: 0, StatusMsg: "success"},
		User:   model.UserQuery{ID: u.UserID, FollowCount: u.FollowCount, FollowerCount: u.FollowerCount, Name: name},
	}, nil
}

// 密码加密
func encryptPassWord(password string) (string, error) {

	p, err := bcrypt.GenerateFromPassword([]byte(password), config.BcryptCost)
	if err != nil {
		return "", errors.New("创建用户失败")
	}
	return string(p), nil
}

// 操作 dao层获取 user_name
func getUserName(userid int64) (string, error) {
	u, err := dao.NewUserDaoInstance().GetUserById(userid)
	if err != nil {
		return "", err
	}

	return u.UserName, nil
}
