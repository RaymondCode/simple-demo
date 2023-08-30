package logic

import (
	"context"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tikstart/common"
	"tikstart/model"
	"tikstart/service/rpc/user/internal/svc"
	"tikstart/service/rpc/user/user"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *user.CreateRequest) (*user.CreateResponse, error) {
	username := in.Username
	password := in.Password

	var count int64
	err := l.svcCtx.DB.Model(&model.User{}).Where("username = ?", username).Count(&count).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if count > 0 {
		return nil, common.ErrUserAlreadyExists.Err()
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	newUser := &model.User{
		Username: username,
		Password: hashedPassword,
	}

	err = l.svcCtx.DB.Create(newUser).Error
	if err != nil {
		st, _ := status.New(codes.Internal, "error creating user record").WithDetails(
			&any.Any{
				Value: []byte(err.Error()),
			})
		return nil, st.Err()
	}

	return &user.CreateResponse{
		UserId: newUser.UserId,
	}, nil
}
