package schema

import (
	"fmt"
	"tikstart/http/internal/types"
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
