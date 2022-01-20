package usecase

import (
	"errors"
	"falabella/internal/domain"
)

var ErrBeerNotExists = errors.New("el id de la cerveza no existe")

type BoxPriceBeer struct {
	currencyRepository domain.CurrencyRepository
	beerRepository     domain.BeerRepository
}

func NewBoxPriceBeer(currencyRepository domain.CurrencyRepository, beerRepository domain.BeerRepository) *BoxPriceBeer {
	return &BoxPriceBeer{
		currencyRepository: currencyRepository,
		beerRepository:     beerRepository,
	}
}

func (b *BoxPriceBeer) Calculate(beerID int, currency string, quantity int) (float32, error) {
	beer, err := b.beerRepository.FindByID(beerID)
	if err != nil {
		return 0, err
	}
	if beer == nil {
		return 0, ErrBeerNotExists
	}

	currenciesInUSD, err := b.currencyRepository.ConvertToUSD([]string{beer.Currency(), currency})
	if err != nil {
		return 0, err
	}

	total := float32(quantity) * (beer.Price() / currenciesInUSD[beer.Currency()])
	return total * currenciesInUSD[currency], nil
}
