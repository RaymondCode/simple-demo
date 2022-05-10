package controller

import (
	"github.com/NoCLin/douyin-backend-go/model"
)

var DemoVideos = []model.VideoResponse{
	{
		Video: model.Video{
			Id:       1,
			Author:   DemoUser.User,
			PlayUrl:  "https://www.w3schools.com/html/movie.mp4",
			CoverUrl: "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		},
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []model.CommentResponse{
	{
		Comment: model.Comment{
			Id:         1,
			User:       DemoUser.User,
			Content:    "Test Comment",
			CreateDate: "05-01",
		},
	},
}

var DemoUser = model.UserInfo{
	User: model.User{
		Id:   1,
		Name: "TestUser",
	},
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}

var usersLoginInfo = map[string]model.UserInfo{
	"zhangleidouyin": {
		User: model.User{
			Id:   1,
			Name: "zhanglei",
		},
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}
