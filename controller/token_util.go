package controller

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey string = "testSecretKey"

func generateToken(username string) string {
	// create a token object
	token := jwt.New(jwt.SigningMethodHS256)
	// set claims of the token
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix() // expired time is 48 hours

	// generate toke string
	tokenString, _ := token.SignedString([]byte(secretKey))

	return tokenString
}

func checkToken(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil //secret key
	})
	if token.Valid {
		//fmt.Println("token checked")
		return true
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		return false
	} else {
		fmt.Println("Unable to handle token", err)
		return false
	}
}

func getUsername(tokenString string) (string, error) {
	if checkToken(tokenString) {
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 返回签名密钥
			return []byte(secretKey), nil
		})
		claims := token.Claims.(jwt.MapClaims)
		username := claims["username"].(string)
		return username, nil
	} else {
		err := errors.New("unable to get user name with token")
		return "", err
	}
}
