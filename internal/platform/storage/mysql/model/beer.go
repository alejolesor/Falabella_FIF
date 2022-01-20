package model

// Beer ...
type Beer struct {
	ID       int     `db:"id" gorm:"column:id"`
	Name     string  `db:"name" gorm:"column:name"`
	Brewery  string  `db:"brewery" gorm:"column:brewery"`
	Country  string  `db:"country" gorm:"column:country"`
	Price    float32 `db:"price" gorm:"column:price"`
	Currency string  `db:"currency" gorm:"column:currency"`
}

func NewModelBeer(name, brewery, country string, price float32, currency string) *Beer {
	return &Beer{
		Name:     name,
		Brewery:  brewery,
		Country:  country,
		Price:    price,
		Currency: currency,
	}
}
