package service

type Group struct {
	FavoriteService
	FeedService
	UserService
	// VideoService
	// ...
}

var GroupApp = new(Group)
