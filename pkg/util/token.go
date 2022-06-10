package util

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

//用户信息类，作为生成token的参数
type UserClaims struct {
	ID       int64  `json:"user_id"`
	Name     string `json:"name"`
	PassWord string `json:"password"`
	//jwt-go提供的标准claim
	jwt.StandardClaims
}

var (
	//自定义的token秘钥
	secret = []byte("16849841325189456f487")
)

// 生成token
func GenerateToken(claims *UserClaims) (string, error) {
	//生成token
	sign, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(secret)
	if err != nil {
		//这里因为项目接入了统一异常处理，所以使用panic并不会使程序终止，如不接入，可使用原始方式处理错误
		//接入统一异常可参考 https://blog.csdn.net/u014155085/article/details/106733391
		panic(err)
	}
	return sign, nil
}

// 解析Token
func ParseToken(tokenString string) (*UserClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("token is invalid")
	}
	return claims, nil
}
