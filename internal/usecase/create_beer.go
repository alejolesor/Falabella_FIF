package usecase

import (
	"errors"
	"falabella/internal/domain"
)

type CreateBeer struct {
	repository domain.BeerRepository
}

func NewCreateBeer(r domain.BeerRepository) *CreateBeer {
	return &CreateBeer{
		repository: r,
	}
}

func (c *CreateBeer) CreateBeer(beer *domain.Beer) (int, error) {
	existBeer, err := c.repository.FindByName(beer.Name())
	if err != nil {
		return 0, err
	}
	if existBeer {
		return 0, errors.New("beer ID already exists")
	}
	id, err := c.repository.CreateBeer(beer)
	if err != nil {
		return 0, err
	}
	return id, nil
}
