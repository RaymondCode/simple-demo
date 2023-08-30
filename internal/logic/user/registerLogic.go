package user

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/anypb"
	error2 "tikstart/common/error"
	"tikstart/common/utils"
	"tikstart/service/rpc/user/user"

	"tikstart/internal/svc"
	"tikstart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	username := req.Username
	password := req.Password

	usernamePattern := "^[^\\s]{1,20}$"
	passwordPattern := "^[!-~]{8,24}$"

	if !utils.MatchRegexp(usernamePattern, username) || !utils.MatchRegexp(passwordPattern, password) {
		return nil, error2.ApiError{
			StatusCode: 422,
			Code:       42201,
			Message:    "Invalid username or password",
		}
	}

	res, err := l.svcCtx.UserRpc.Create(l.ctx, &user.CreateRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		if st, match := utils.MatchError(err, error2.ErrUserAlreadyExists); match {
			return nil, error2.ApiError{
				StatusCode: 422,
				Code:       42201,
				Message:    "用户名已被使用",
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

	return &types.RegisterResponse{
		BasicResponse: types.BasicResponse{
			StatusCode: 0,
			StatusMsg:  "",
		},
		UserId: res.UserId,
		Token:  tokenString,
	}, nil
}
