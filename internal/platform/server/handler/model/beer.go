package model

import "falabella/internal/domain"

type Beer struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Brewery  string  `json:"brewery"`
	Country  string  `json:"country"`
	Price    float32 `json:"price"`
	Currency string  `json:"currency"`
}

func NewBeer(b *domain.Beer) *Beer {
	return &Beer{
		ID:       b.ID(),
		Name:     b.Name(),
		Brewery:  b.Brewery(),
		Country:  b.Country(),
		Price:    b.Price(),
		Currency: b.Currency(),
	}
}
