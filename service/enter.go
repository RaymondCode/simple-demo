package service

type Group struct {
	FavoriteService
	FeedService
	UserService
	PublishService
	//VideoService
	// ...
}

var GroupApp = new(Group)
