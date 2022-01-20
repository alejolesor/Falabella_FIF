package mysql

import (
	"falabella/internal/domain"
	"falabella/internal/platform/storage/mysql/dao"
	"falabella/internal/platform/storage/mysql/model"
)

type BeerRepository struct {
	dao dao.IDao
}

func NewBeerRepository(d dao.IDao) *BeerRepository {
	return &BeerRepository{
		dao: d,
	}
}

func (b BeerRepository) FindAll() ([]domain.Beer, error) {
	beerList := []domain.Beer{}
	// execute dao
	beerModel, err := b.dao.Get()
	if err != nil {
		return nil, err
	}
	for _, beer := range beerModel {
		beerDomain := domain.NewBeer(beer.ID, beer.Name, beer.Brewery, beer.Country, beer.Price, beer.Currency)
		beerList = append(beerList, *beerDomain)
	}

	return beerList, nil
}

func (b BeerRepository) FindByID(id int) (*domain.Beer, error) {
	beer, err := b.dao.GetById(id)
	if err != nil {
		return nil, err
	}
	return domain.NewBeer(beer.ID, beer.Name, beer.Brewery, beer.Country, beer.Price, beer.Currency), nil
}

func (b BeerRepository) CreateBeer(beer *domain.Beer) (int, error) {
	id, err := b.dao.Create(model.NewModelBeer(beer.Name(), beer.Brewery(), beer.Country(), beer.Price(), beer.Currency()))
	if err != nil {
		return 0, nil
	}
	return id, nil
}

func (b BeerRepository) FindByName(name string) (bool, error) {
	result, err := b.dao.GetByName(name)
	if err != nil {
		return false, err
	}
	return result, nil

}
