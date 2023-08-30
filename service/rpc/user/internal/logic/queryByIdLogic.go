package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"tikstart/common"
	"tikstart/model"

	"tikstart/service/rpc/user/internal/svc"
	"tikstart/service/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryByIdLogic {
	return &QueryByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryByIdLogic) QueryById(in *user.QueryByIdRequest) (*user.QueryResponse, error) {
	userId := in.UserId

	userRecord := model.User{}
	err := l.svcCtx.DB.Where("user_id = ?", userId).First(&userRecord).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrUserNotFound.Err()
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &user.QueryResponse{
		UserId:    userRecord.UserId,
		Username:  userRecord.Username,
		Password:  userRecord.Password,
		CreatedAt: userRecord.CreatedAt.Unix(),
		UpdatedAt: userRecord.UpdatedAt.Unix(),
	}, nil
}
