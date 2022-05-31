package repository

type Group struct {
	VideoRepository
	UserRepository
	RelationRepository
	// VideoRepository
	// ...
}

var GroupApp = new(Group)
