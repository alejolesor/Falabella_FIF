package usecase

import (
	"falabella/internal/domain"
)

type FindBeer struct {
	repository domain.BeerRepository
}

func NewFindBeer(r domain.BeerRepository) *FindBeer {
	return &FindBeer{
		repository: r,
	}
}

func (b *FindBeer) FindAll() ([]domain.Beer, error) {
	beer, err := b.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return beer, nil
}

func (b *FindBeer) FindById(id int) (*domain.Beer, error) {
	beer, err := b.repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return beer, nil
}
