package response

import (
	"github.com/RaymondCode/simple-demo/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserListResponse struct {
	Response
	UserList []*model.User `json:"user_list"`
}

func OkWithUserList(userList []*model.User, c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: SUCCESS,
		},
		UserList: userList,
	})
}
