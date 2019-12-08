package quote

import (
	"github.com/eucleciojosias/codenation-challenge/pkg/entity"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) FindByActor(actor string) ([]*entity.Quote, error) {
	return s.repo.FindByActor(actor)
}

func (s *Service) FindAll() ([]*entity.Quote, error) {
	return s.repo.FindAll()
}
