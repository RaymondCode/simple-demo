package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func CreateToken(userId int64, secret string, expire int64) (string, error) {
	userClaims := UserClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expire) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "app-api",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	signed, err := token.SignedString([]byte(secret))
	return signed, err
}

func ParseToken(tokenString string, secret string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if token == nil {
		return nil, err
	}

	userClaims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, err
	}
	return userClaims, err
}
