package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
	"time"
)

const (
	TokenExpireDuration = time.Hour * 24 //token过期时间
)

var mySecret = []byte("going小分队")

// 仅用于根据token值返回密钥
func keyFunc(token *jwt.Token) (interface{}, error) {
	return mySecret, nil
}

type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenToken 根据用户id生成token，如需修改过期时间请修改上面的常量 TokenExpireDuration
func GenToken(userID int64) (Token string, err error) {
	// 建立自己的token字段
	c := MyClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    viper.GetString("app.name"),
		},
	}
	// 加密并获的完整编码后的Token
	if Token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret); err != nil {
		return "", err
	}
	return
}

// ParseToken 解析token，返回包含信息的结构体
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	var claims = new(MyClaims) //解析好的存放在mc中

	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
