package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"github.com/BaiZe1998/douyin-simple-demo/dto"
	"github.com/BaiZe1998/douyin-simple-demo/pkg/util"
	"github.com/BaiZe1998/douyin-simple-demo/service"
	"github.com/gin-gonic/gin"
)

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	userClaims, _ := util.ParseToken(token)
	userModel, _ := model.QueryUserById(context.Background(), userClaims.ID)
	users := dto.User{
		Id:            userModel.ID,
		Name:          userModel.Name,
		FollowCount:   userModel.FollowCount,
		FollowerCount: userModel.FollowerCount,
		IsFollow:      false,
	}

	if userModel.ID > 0 {
		if actionType == "1" {
			text := c.Query("comment_text")
			//comment addcomment
			responseComment := service.AddComment(text, users, videoId)
			c.JSON(http.StatusOK,
				dto.CommentActionResponse{
					Response: dto.Response{StatusCode: 0},
					Comment:  responseComment,
				})
		} else {
			commentId, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
			//comment delete
			model.DeleteCommnet(context.Background(), videoId, commentId)
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	res, total, _ := model.QueryComment(context.Background(), videoId, 10, 0)
	c.JSON(http.StatusOK, dto.CommentListResponse{
		Response:    dto.Response{StatusCode: 0, StatusMsg: ""},
		CommentList: res,
	})
	fmt.Println(total)
}
