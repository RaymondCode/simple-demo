package service

import (
	"github.com/RaymondCode/simple-demo/models"
	"golang.org/x/crypto/bcrypt"
)

const (
	MaxUsernameLength = 32
	MaxPasswordLength = 32
	MinPasswordLength = 6
)

func IsUserLegal(username string, password string) error {
	if username == "" {
		return ErrorUserNameNull
	}
	if len(username) > MaxUsernameLength {
		return ErrorUserNameExtend
	}
	if password == "" {
		return ErrorPasswordNull
	}
	if len(password) > MaxPasswordLength || len(password) < MinPasswordLength {
		return ErrorPasswordLength
	}
	return nil
}

func isUserExistByName(username string) bool {
	if _, err := models.NewUserDaoInstance().FindUserByName(username); err != nil {
		return false
	}
	return true
}

func CreateRegisterUser(username string, password string) (int64, error) {
	newPassword, _ := HashAndSalt(password)
	newUser := models.User{
		Name:     username,
		Password: newPassword,
	}
	if isUserExistByName(username) {
		return -1, ErrorUserExit
	} else {
		userId, err := models.NewUserDaoInstance().CreateUser(&newUser)
		if err != nil {
			panic(err)
		}
		return userId, err
	}
}

func FindLoginUser(username string, password string) (int64, error) {
	login, err := models.NewUserDaoInstance().FindUserByName(username)
	if err != nil {
		return -1, ErrorFullPossibility
	}
	if !ComparePasswords(login.Password, password) {
		return -1, ErrorPasswordFalse
	}
	return login.UserId, nil
}

func HashAndSalt(pwdstr string) (string, error) {
	pwd := []byte(pwdstr)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	pwdHash := string(hash)
	return pwdHash, nil
}

func ComparePasswords(pwdHash string, pwdPlain string) bool {
	hash := []byte(pwdHash)
	plain := []byte(pwdPlain)
	err := bcrypt.CompareHashAndPassword(hash, plain)
	if err != nil {
		return false
	}
	return true
}
