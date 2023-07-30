package controller

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const DatabaseAddress string = "root:shanwer666@tcp(localhost:3306)/momotok" //SQL address(1024)

// sql statements of the table,database name:momotok
// CREATE TABLE user (
//	id INT PRIMARY KEY AUTO_INCREMENT,
//	username VARCHAR(50) NOT NULL UNIQUE,
//	ip VARCHAR(15),
//	password VARCHAR(60),
//	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );

// usersLoginInfo use map to store user info, token is created by jwt
// user data will be stored in the 1024 workspace MySQL server
var usersLoginInfo = map[string]User{
	//TODO:其他模块需要改写为从数据库中获取
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	ip := c.ClientIP()
	hashedPassword, err := hashPassword(password)

	if err != nil {
		fmt.Println("Hash password failed：", err)
		return
	}
	db, err := sql.Open("mysql", DatabaseAddress)
	if err != nil {
		fmt.Println("Database connected failed: ", err)
	}
	_, err = db.Exec("INSERT INTO user (username, ip, password) VALUES (?, ?, ?)", username, ip, hashedPassword)
	if err != nil {
		fmt.Println("Register failed：", err)
		return
	}
	id := int64(0)
	err = db.QueryRow("SELECT ID FROM user WHERE username = ?", username).Scan(&id)
	tokenString := generateToken(username)

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   id,
		Token:    tokenString,
	})
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	db, err := sql.Open("mysql", DatabaseAddress)
	if err != nil {
		fmt.Println("Database connected failed: ", err)
	}
	storedPassword := ""
	id := int64(0)
	err = db.QueryRow("SELECT password, id FROM user WHERE username = ?", username).Scan(&storedPassword, &id)
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
	if err != nil {
		//fmt.Println("Wrong username or password: ", err)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "Wrong username or password"},
		})
		return
	}
	tokenString := generateToken(username)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   id,
		Token:    tokenString,
	})
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	//TODO:这里要改为从数据库查询
	if user, exist := usersLoginInfo[token]; exist { //这里改为checkToken函数
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

// generate hashed password
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
