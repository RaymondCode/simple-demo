package service

import (
	"github.com/life-studied/douyin-simple/dao"
)

func IsTokenExists(Token string) bool {
	return dao.QueryToken
}
