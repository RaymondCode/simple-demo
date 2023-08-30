package user

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/anypb"
	error2 "tikstart/common/error"
	"tikstart/common/utils"
	"tikstart/service/rpc/user/user"

	"tikstart/internal/svc"
	"tikstart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	username := req.Username
	password := req.Password

	res, err := l.svcCtx.UserRpc.QueryByName(l.ctx, &user.QueryByNameRequest{
		Username: username,
	})
	if err != nil {
		if st, match := utils.MatchError(err, error2.ErrUserNotFound); match {
			return nil, error2.ApiError{
				StatusCode: 422,
				Code:       42202,
				Message:    "用户名不存在",
			}
		} else {
			for index, item := range st.Details() {
				detail := item.(*anypb.Any)
				fmt.Printf("%d: %s\n", index, string(detail.Value))
			}

			return nil, error2.ServerError{
				ApiError: error2.ApiError{
					StatusCode: 500,
					Code:       50000,
					Message:    "Internal Server Error",
				},
				Detail: err,
			}
		}
	}

	err = bcrypt.CompareHashAndPassword(res.Password, []byte(password))
	if err != nil {
		return nil, error2.ApiError{
			StatusCode: 422,
			Code:       42203,
			Message:    "密码错误",
		}
	}

	tokenString, err := utils.CreateToken(res.UserId, l.svcCtx.Config.JwtAuth.Secret, l.svcCtx.Config.JwtAuth.Expire)
	if err != nil {
		return nil, error2.ServerError{
			ApiError: error2.ApiError{
				StatusCode: 500,
				Code:       50000,
				Message:    "Internal Server Error",
			},
			Detail: err,
		}
	}

	return &types.LoginResponse{
		BasicResponse: types.BasicResponse{
			StatusCode: 0,
			StatusMsg:  "",
		},
		UserId: res.UserId,
		Token:  tokenString,
	}, nil
}
