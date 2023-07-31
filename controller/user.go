package controller

import (
	"Momotok-Server/rpc"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"strconv"
)

const DatabaseAddress string = "root:shanwer666@tcp(localhost:3306)/momotok" //SQL address(1024)
//const DatabaseAddress string = "root:root@tcp(localhost:3306)/momotok" //SQL address(1024)
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
	if err != nil { //duplicate username check
		mysqlErr, ok := err.(*mysql.MySQLError)
		if !ok {
			fmt.Println("Register failed：", err)
		}
		if mysqlErr.Number == 1062 {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 1, StatusMsg: "User:" + username + "already exists!"},
			})
		} else {
			fmt.Println("Register failed：", err)
		}
		return
	}
	id := int64(0)
	err = db.QueryRow("SELECT ID FROM user WHERE username = ?", username).Scan(&id)
	token := generateToken(username)

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   id,
		Token:    token,
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
	token := generateToken(username)
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   id,
		Token:    token,
	})
}

func UserInfo(c *gin.Context) {
	uid := c.Query("user_id")
	if !checkToken(c.Query("token")) {
		return
	}

	resp, _ := rpc.HttpRequest("GET", "https://v1.hitokoto.cn/?c=a&c=d&c=i&c=k&encode=text", nil)
	if resp.Body != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(resp.Body)
	}
	signature, _ := io.ReadAll(resp.Body)

	id, _ := strconv.ParseInt(uid, 10, 64)
	userInfo := User{
		Id:              id,
		Signature:       string(signature),
		Avatar:          "https://acg.suyanw.cn/sjtx/random.php",
		BackgroundImage: "https://acg.suyanw.cn/api.php",
	}

	db, err := sql.Open("mysql", DatabaseAddress)
	if err != nil {
		fmt.Println("Database connected failed: ", err)
		return
	}
	err = db.QueryRow("SELECT username FROM user WHERE id = ?", uid).Scan(&userInfo.Name)
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "user not found"},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "Success",
		"user":        userInfo,
		//"signature":   string(signature),
		//"avatar":      avatar,
		//"background_image": background_image,
	})
	return
}

// generate hashed password
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
