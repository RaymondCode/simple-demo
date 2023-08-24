package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/life-studied/douyin-simple/dao"
)

const (
	MaxUsernameLength = 32 //用户名最大长度
	MaxPasswordLength = 32 //密码最大长度
	MinPasswordLength = 6  //密码最小长度
)

// 合法性加密
func IsUserLegal(userName string, passWord string) error {
	//1.用户名检验
	if userName == "" {
		return errors.New("用户名为空")
	}
	if len(userName) > MaxUsernameLength {
		return errors.New("用户名长度不符合规范")
	}
	//2.密码检验
	if passWord == "" {
		return errors.New("密码为空")
	}
	if len(passWord) > MaxPasswordLength || len(passWord) < MinPasswordLength {
		return errors.New("密码长度不符合规范")
	}
	return nil
}

// 对token进行加密
func Encryption(username, password string) string {
	token := username + password
	hash := sha256.New()
	hash.Write([]byte(token))
	enToken := hex.EncodeToString(hash.Sum(nil))
	return enToken
}

// 查询用户名是否重复
func IsUsernameExists(username string) bool {
	return dao.QueryName(username)
}

// 创建用户并存入数据库
func CreateInfo(id int64, username string, password string) error {
	err := dao.AddUserInfo(id, username, password)
	if err != nil {
		return err
	}
	return nil
}

// 获取数据库所有数据
func RequireAllUser() ([]dao.User, error) {
	users, err := dao.GetAllUsers()
	if err != nil {
		fmt.Println("获取用户数据出错：", err)
		return nil, err
	}
	return users, nil
}
