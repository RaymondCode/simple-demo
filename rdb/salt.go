package rdb

import (
	"github.com/warthecatalyst/douyin/common"
)

func GetAllSalts() []string {
	return rdb.SMembers(common.KeySalt).Val()
}

func GetRandomSalt() []byte {
	return []byte(rdb.SRandMemberN(common.KeySalt, 1).Val()[0])
}
