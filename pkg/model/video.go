package model

import "gorm.io/gorm"

type Video struct {
	ID            int64          `json:"id,omitempty" gorm:"primaryKey;comment:短视频id;autoIncrement;unique_index:create_time_index"`
	AuthorID      int64          `json:"-" gorm:"not null;comment:作者id;unique_index:create_time_index"`
	PlayUrl       string         `json:"play_url,omitempty" gorm:"size:50;not null;comment:短视频url;unique_index:create_time_index"`
	CoverUrl      string         `json:"cover_url,omitempty" gorm:"size:50;not null;comment:封面url;unique_index:create_time_index"`
	FavoriteCount int64          `json:"favorite_count,omitempty" gorm:"not null;default:0;comment:点赞数;unique_index:create_time_index"`
	CommentCount  int64          `json:"comment_count,omitempty" gorm:"not null;default:0;comment:评论数;unique_index:create_time_index"`
	CreatedAt     int            `json:"created_at" gorm:"comment:投递时间;unique_index:create_time_index;not null"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index;comment:删除标记位;unique_index:create_time_index"`

	//作者
	Author *User `json:"author" gorm:"ForeignKey:AuthorID;"`
	//是否点赞
	IsFavorite bool `json:"is_favorite,omitempty" gorm:"-"`
}

// Video表 CRUD
type VideoCtl interface {
	//创建
	Create(video *Video) (int64, error)
	//删除
	Delete(video *Video) error

	//根据视频id获取视频
	GetVideoById(id int) (*Video, error)
	//获取视频列表by发布者id,按发布时间倒序排列
	GetVideoByAuthorId(authorID int) ([]*Video, error)
	//获取用户总获赞数
	GetFavoriteCount(authorID int) (int64, error)

	//查询所有by发布时间，倒序排列
	FindAllByCreateAt() ([]*Video, error)
	//查询所有by发布时间 ：{lastTime为0时获取最新的视频,否则获取指定时间之前的视频,size表示数量,最多为30}。视频列表按发布时间倒序排列
	GetVideoList(latestTime int64, size int) ([]*Video, error)
	//查询所有by点赞数，倒序排列
	FindAllByFavoriteCount() ([]*Video, error)
	//查询所有by评论数，倒序排列
	FindAllByCommentCount() ([]*Video, error)
}
