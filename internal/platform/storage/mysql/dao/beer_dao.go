package dao

import (
	"database/sql"
	"falabella/internal/platform/storage/mysql/model"
)

const (
	selectBeer     = "SELECT id, name, brewery, country, price, currency FROM beer.BeerItem;"
	insertBeer     = "INSERT INTO beer.BeerItem (name, brewery, country, price, currency) VALUES(?, ?, ?, ?, ?);"
	selectBeerById = "SELECT id, name, brewery, country, price, currency FROM beer.BeerItem WHERE id  = ?;"
	selectByName   = "SELECT id, name, brewery, country, price, currency FROM beer.BeerItem WHERE name = ?;"
)

type BeerDao struct {
	db *sql.DB
}

func NewBeerDao(db *sql.DB) *BeerDao {
	return &BeerDao{
		db: db,
	}
}

func (bd *BeerDao) Get() ([]model.Beer, error) {
	listBeers := []model.Beer{}
	beerResult, err := bd.db.Query(selectBeer)
	if err != nil {
		return nil, err
	}

	for beerResult.Next() {
		var valuesBeer model.Beer
		err = beerResult.Scan(&valuesBeer.ID, &valuesBeer.Name, &valuesBeer.Brewery, &valuesBeer.Country, &valuesBeer.Price, &valuesBeer.Currency)
		if err != nil {
			return nil, err
		}
		listBeers = append(listBeers, valuesBeer)
	}

	return listBeers, nil
}

func (bd *BeerDao) GetById(id int) (*model.Beer, error) {
	valuesBeer := model.Beer{}
	stmt, err := bd.db.Prepare(selectBeerById)
	if err != nil {
		return nil, err
	}

	rowBeer, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	for rowBeer.Next() {
		err = rowBeer.Scan(&valuesBeer.ID, &valuesBeer.Name, &valuesBeer.Brewery, &valuesBeer.Country, &valuesBeer.Price, &valuesBeer.Currency)
		if err != nil {
			return nil, err
		}
	}

	return &valuesBeer, nil
}

func (bd *BeerDao) Create(beer *model.Beer) (int, error) {
	stmt, err := bd.db.Prepare(insertBeer)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(beer.Name, beer.Brewery, beer.Country, beer.Price, beer.Currency)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (bd *BeerDao) GetByName(name string) (bool, error) {
	valuesBeer := model.Beer{}
	stmt, err := bd.db.Prepare(selectByName)
	if err != nil {
		return false, err
	}

	rowBeer, err := stmt.Query(name)
	if err != nil {
		return false, err
	}

	for rowBeer.Next() {
		err = rowBeer.Scan(&valuesBeer.ID, &valuesBeer.Name, &valuesBeer.Brewery, &valuesBeer.Country, &valuesBeer.Price, &valuesBeer.Currency)
		if err != nil {
			return false, err
		}
	}

	if valuesBeer.ID != 0 {
		return true, nil
	}

	return false, nil

}
