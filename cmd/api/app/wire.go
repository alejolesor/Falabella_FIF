package app

import (
	"falabella/internal/platform/server/handler"
	"falabella/internal/platform/storage/mysql"
	configDB "falabella/internal/platform/storage/mysql/config"
	"falabella/internal/platform/storage/mysql/dao"
	"falabella/internal/platform/storage/rest"
	"falabella/internal/usecase"
)

func DependencyInjection() *handler.Beer {
	db := configDB.NewConnectionDB()

	dao := dao.NewBeerDao(db)
	beerRepository := mysql.NewBeerRepository(dao)
	currencyRepository := rest.NewCurrencyRepository()

	findBeer := usecase.NewFindBeer(beerRepository)
	createBeer := usecase.NewCreateBeer(beerRepository)
	boxPriceBeer := usecase.NewBoxPriceBeer(currencyRepository, beerRepository)
	beers := handler.NewBeer(findBeer, createBeer, boxPriceBeer)

	return beers
}
