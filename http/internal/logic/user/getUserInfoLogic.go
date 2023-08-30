package user

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/types/known/anypb"
	"tikstart/common"
	"tikstart/common/utils"
	"tikstart/http/internal/svc"
	"tikstart/http/internal/types"
	"tikstart/http/schema"
	"tikstart/rpc/user/user"
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
		if st, match := utils.MatchError(err, common.ErrUserNotFound); match {
			return nil, schema.ApiError{
				StatusCode: 422,
				Code:       42202,
				Message:    "用户名不存在",
			}
		} else {
			for index, item := range st.Details() {
				detail := item.(*anypb.Any)
				fmt.Printf("%d: %s\n", index, string(detail.Value))
			}

			return nil, schema.ServerError{
				ApiError: schema.ApiError{
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
