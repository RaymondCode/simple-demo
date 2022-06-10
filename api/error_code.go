package api

//不同的error对应的errorCode，以及对应的message(可选)
const (
	//视频保存、上传、截取等错误
	UploadFailErr      = 10001
	SavingFailErr      = 10002
	InputVideoCheckErr = 10003

	//数据库和用户错误
	InnerErr          = 10101
	TokenInvalidErr   = 10102
	UserNotExistErr   = 10103
	UserIdNotMatchErr = 10104
	UnKnownActionType = 10105
	RecordNotExistErr = 10106

	//输入逻辑错误
	InputFormatCheckErr = 10208
	LogicErr            = 10209
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
