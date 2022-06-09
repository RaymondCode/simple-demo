package api

const (
	InnerErr          = 10001
	TokenInvalidErr   = 10002
	UserNotExistErr   = 10003
	UserIdNotMatchErr = 10004
	UnKnownActionType = 10005
	RecordNotExistErr = 10006
	UploadFailErr     = 10007
)

var ErrorCodeToMsg = map[int]string{
	InnerErr:          "发生内部错误",
	TokenInvalidErr:   "非法的Token",
	UserNotExistErr:   "用户不存在",
	UserIdNotMatchErr: "用户Id不匹配",
	UnKnownActionType: "非法的操作",
	RecordNotExistErr: "不存在对应的数据",
	UploadFailErr:     "文件上传失败",
}
