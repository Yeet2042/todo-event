package application

import (
	"errors"

	"todoe/domain/task/port"
)

var (
	ErrInvalidTitle  = errors.New("title must not be empty")
	ErrInvalidStatus = errors.New("invalid status")
)

type Service struct {
	repo      port.Repository
	publisher port.Publisher
}

//var _ port.UseCase = (*Service)(nil)

func NewService(repo port.Repository, publisher port.Publisher) *Service {
	return &Service{repo: repo, publisher: publisher}
}
