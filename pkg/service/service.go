package service

import "github.com/MarkStroykow/any-app/pkg/repository"

type Authotization interface {
}

type AppCart interface {
}

type AppItem interface {
}

type Service struct {
	Authotization
	AppCart
	AppItem
}

func NewService(rep *repository.Repository) *Service {
	return &Service{}
}
