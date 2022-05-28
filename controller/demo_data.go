package controller

import "github.com/RaymondCode/simple-demo/model"

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}

var user1 = model.User{
	Id:            1,
	Name:          "zhangsan",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}

var user2 = model.User{
	Id:            2,
	Name:          "lisi",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}

var videos = []model.Video{
	{
		Id:            1,
		Author:        user1,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         "测试1",
	},
	{
		Id:            2,
		Author:        user2,
		PlayUrl:       "https://www.w3schools.com/html/mov_bbb.mp4",
		CoverUrl:      "https://www.topgoer.cn/uploads/blog/202111/attach_16b79b08788038a8.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         "测试2",
	},
}
