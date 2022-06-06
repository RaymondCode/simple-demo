package common

const (
	UserInfoTable    = "t_douyin_users"
	VideoInfoTable   = "t_douyin_videos"
	LikeInfoTable    = "t_douyin_favourites"
	FollowInfoTable  = "t_douyin_follows"
	CommentInfoTable = "t_douyin_comments"
)

const (
	UserId   = "user_id"
	UserName = "user_name"
)

const (
	KeySalt = "salt"
)

const (
	LikeOn  = 1
	LikeOff = 2

	FollowOn  = 1
	FollowOff = 2

	CommentOn  = 1
	CommentOff = 2

	PublishOn  = 1
	PublishOff = 2

	NameOn  = 1
	NameOff = 2
)

const (
	MaxVideoCount = 30
)
