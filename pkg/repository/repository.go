package repository

type Authotization interface {
}

type AppCart interface {
}

type AppItem interface {
}

type Repository struct {
	Authotization
	AppCart
	AppItem
}

func NewRepository() *Repository {
	return &Repository{}
}
