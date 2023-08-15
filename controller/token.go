package controller

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(username string) (string, error) {
	expireDuration, _ := time.ParseDuration("23h59m59s")
	expireTime := time.Now().Add(expireDuration)
	claims := jwt.StandardClaims{
		Audience:  username,
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "tiktok",
		NotBefore: time.Now().Unix(),
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("tik tok"))
	return token, err
}

func ParseToken(token string) (string, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("tik tok"), nil
	})

	if err == nil && tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*jwt.StandardClaims); ok && tokenClaims.Valid {
			return claims.Audience, nil
		}
	}

	return "", err
}
