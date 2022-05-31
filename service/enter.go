package service

type Group struct {
	FavoriteService
	FeedService
	UserService
	PublishService
	RelationService
	//VideoService
	// ...
}

var GroupApp = new(Group)
