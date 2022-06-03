package utils

import (
	"github.com/RaymondCode/simple-demo/global"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(global.CONFIG.JWT.SigningKey)

type Claims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken 根据userId生成token
func GenerateToken(userId int64) (string, error) {
	claims := Claims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + global.CONFIG.JWT.ExpiresTime, // 过期时间7天
			Issuer:    global.CONFIG.JWT.Issuer,                          // 签名的发行者
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
