package error

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"tikstart/internal/types"
)

var (
	ErrUserAlreadyExists = status.New(codes.AlreadyExists, "user already exists")
	ErrUserNotFound      = status.New(codes.NotFound, "user not found")
)

type ApiError struct {
	StatusCode int
	Code       int32
	Message    string
}

type ServerError struct {
	ApiError
	Detail error
}

func (e ApiError) Error() string {
	return fmt.Sprintf("(%d) %s", e.Code, e.Message)
}

func (e ApiError) Response() *types.BasicResponse {
	return &types.BasicResponse{
		StatusCode: e.Code,
		StatusMsg:  e.Message,
	}
}
