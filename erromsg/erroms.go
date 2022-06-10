package erromsg

type Eros struct {
	Code    int32
	Message string
}

//OK 正常返回代码以及值
//Fail 错误返回值
var OK = &Eros{Code: 0, Message: "OK"}
var Fail = &Eros{Code: -1, Message: "Error"}

var (
	//视频类
	ErrVideoUpload           = &Eros{Code: 10001, Message: "视频上传失败"}
	ErrCreateVideoRecordFail = &Eros{Code: 10002, Message: "数据库新增视频记录失败"}
	ErrQueryVideosFail       = &Eros{Code: 10003, Message: "查询视频信息失败"}
	//数据库类
	ErrDataBase            = &Eros{Code: 10101, Message: "数据库错误"}
	ErrQueryUserInfoFail   = &Eros{Code: 10102, Message: "查询用户信息错误"}
	ErrQueryUserLoginFail  = &Eros{Code: 10103, Message: "查询用户登录信息错误"}
	ErrCreateUserFail      = &Eros{Code: 10104, Message: "创建用户信息失败"}
	ErrCreateUserLoginFail = &Eros{Code: 10105, Message: "创建用户登录信息失败"}
	//Token类
	//用户类
	ErrPassWordWrong       = &Eros{Code: 10401, Message: "密码错误"}
	ErrEncryptPassWordFail = &Eros{Code: 10402, Message: "密码加密失败"}
	ErrQueryUserNameFail   = &Eros{Code: 10403, Message: "获取用户名失败"}
	//点赞类
	//评论类
)
