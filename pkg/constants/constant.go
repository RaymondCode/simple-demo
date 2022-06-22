package constants

const (
	MySQLDefaultDSN = "dyweadmin:JXU2MTNGJXU5NzUyJXU5Rjk5JXU2MzA3JXU1RjE1JXU0RjYw@tcp(xinzf0520.top:20607)/douyin?charset=utf8&parseTime=True&loc=Local"
	MySQLLocalDSN   = "root:root@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"

	RedisLocalDSN   = "localhost:6379"
	RedisLocalPWD   = ""
	RedisDefaultDSN = "localhost:6379"
	RedisDefaultPWD = ""
)

var RedisDBList = map[string]int{
	"Default": 0,
}
