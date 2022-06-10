package api

//不同的error对应的errorCode，以及对应的message(可选)
const (
	InnerErr            = 10001
	TokenInvalidErr     = 10002
	UserNotExistErr     = 10003
	UserIdNotMatchErr   = 10004
	UnKnownActionType   = 10005
	RecordNotExistErr   = 10006
	UploadFailErr       = 10007
	InputFormatCheckErr = 10008
	LogicErr            = 10009
	InputVideoCheckErr  = 10010
)

var ErrorCodeToMsg = map[int]string{
	InnerErr:            "发生数据库错误",
	TokenInvalidErr:     "非法的Token",
	UserNotExistErr:     "用户不存在",
	UserIdNotMatchErr:   "用户Id不匹配",
	UnKnownActionType:   "非法的操作",
	RecordNotExistErr:   "不存在对应的数据",
	UploadFailErr:       "文件上传失败",
	InputFormatCheckErr: "参数格式错误",
	LogicErr:            "逻辑错误",
	InputVideoCheckErr:  "上传视频有误",
}
