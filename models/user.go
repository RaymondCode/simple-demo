package models

import (
	"sync"
)

type User struct {
	UserId int64 `gorm:"primary_key"`
	Name   string
	// FollowingCount  int64  `gorm:"default:(-)"`
	// FollowerCount   int64  `gorm:"default:(-)"`
	Password string
	// Avatar          string `gorm:"default:(-)"`
	// BackgroundImage string `gorm:"default:(-)"`
	Signature string
	// TotalFavorited  int64  `gorm:"default:(-)"`
	// WorkCount       int64  `gorm:"default:(-)"`
	// FavoriteCount   int64  `gorm:"default:(-)"`
	// CreateAt        time.Time
	// DeleteAt        time.Time
}

type UserDao struct{}

var userDao *UserDao
var userOnce sync.Once

func NewUserDaoInstance() *UserDao {
	userOnce.Do(
		func() {
			userDao = &UserDao{}
		},
	)
	return userDao
}

/*
*
根据用户名和密码，创建一个新的User，返回UserId
*/
func (*UserDao) CreateUser(user *User) (int64, error) {
	/*user := User{Name: username, Password: password, FollowingCount: 0, FollowerCount: 0, CreateAt: time.Now()}*/

	result := SqlSession.Create(&user)

	if result.Error != nil {
		return -1, result.Error
	}

	return user.UserId, nil
}

/*
*
根据用户名，查找用户实体
*/
func (*UserDao) FindUserByName(username string) (*User, error) {
	user := User{Name: username}

	result := SqlSession.Where("name = ?", username).First(&user)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

/*
*
根据用户id，查找用户实体
*/
func (d *UserDao) FindUserById(id int64) (*User, error) {
	user := User{UserId: id}

	result := SqlSession.Where("user_id = ?", id).First(&user)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return &user, err
}
