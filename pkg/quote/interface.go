package quote

import "github.com/eucleciojosias/codenation-challenge/pkg/entity"

type Reader interface {
	FindByActor(actor string) ([]*entity.Quote, error)
	FindAll() ([]*entity.Quote, error)
}

type Repository interface {
	Reader
}
