package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	error2 "tikstart/common/error"
	"tikstart/common/model"
	"tikstart/service/rpc/user/internal/svc"
	"tikstart/service/rpc/user/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryByNameLogic {
	return &QueryByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryByNameLogic) QueryByName(in *user.QueryByNameRequest) (*user.QueryResponse, error) {
	username := in.Username

	userRecord := model.User{}
	err := l.svcCtx.DB.Where("username = ?", username).First(&userRecord).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, error2.ErrUserNotFound.Err()
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
