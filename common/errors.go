package common

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrUserAlreadyExists = status.New(codes.AlreadyExists, "user already exists")
	ErrUserNotFound      = status.New(codes.NotFound, "user not found")
)
