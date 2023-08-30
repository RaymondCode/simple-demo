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

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoRequest) (resp *types.GetUserInfoResponse, err error) {
	queryId := req.UserId
	res, err := l.svcCtx.UserRpc.QueryById(l.ctx, &user.QueryByIdRequest{
		UserId: queryId,
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
	return &types.GetUserInfoResponse{
		BasicResponse: types.BasicResponse{
			StatusCode: 200,
			StatusMsg:  "",
		},
		User: types.User{
			Id:   res.UserId,
			Name: res.Username,
		},
	}, nil
}
