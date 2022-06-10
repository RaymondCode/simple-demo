package rdb

import "strconv"

func AddToken(userId int64, token string) {
	rdb.Set(strconv.FormatInt(userId, 10), token, 0)
}

func GetToken(userId int64) string {
	return rdb.Get(strconv.FormatInt(userId, 10)).Val()
}
