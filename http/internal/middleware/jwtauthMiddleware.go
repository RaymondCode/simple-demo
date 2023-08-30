package middleware

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"tikstart/common/utils"
	"tikstart/http/internal/config"
	"tikstart/http/internal/types"
)

type JwtAuthMiddleware struct {
	Config config.Config
}

func NewJwtAuthMiddleware(c config.Config) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{
		Config: c,
	}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.URL.Query().Get("token")

		if token == "" {
			httpx.WriteJson(w, http.StatusUnauthorized, &types.BasicResponse{
				StatusCode: 40101,
				StatusMsg:  "没有提供token",
			})
			return
		}
		_, err := utils.ParseToken(token, m.Config.JwtAuth.Secret)
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				httpx.WriteJson(w, http.StatusUnauthorized, &types.BasicResponse{
					StatusCode: 40102,
					StatusMsg:  "token已过期",
				})
			} else {
				httpx.WriteJson(w, http.StatusUnauthorized, &types.BasicResponse{
					StatusCode: 40103,
					StatusMsg:  "token无效",
				})
			}
			return
		}
		next(w, r)
	}
}
