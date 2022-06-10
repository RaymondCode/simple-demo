package controller

import "github.com/BaiZe1998/douyin-simple-demo/dto"

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser1,
		PlayUrl:       "http://niuyefan.oss-cn-beijing.aliyuncs.com/douyin/bear.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 2,
		CommentCount:  2,
		IsFavorite:    false,
	},
}

// var DemoComments = []Comment{
// 	{
// 		Id:         1,
// 		User:       DemoUser1,
// 		Content:    "Test Comment",
// 		CreateDate: "05-01",
// 	},
// 	{
// 		Id:         2,
// 		User:       DemoUser1,
// 		Content:    "Test Commen22",
// 		CreateDate: "05-012",
// 	},
// }

var DemoUser1 = dto.User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
var DemoUser = User{
	Id:            "1",
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
