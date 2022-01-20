package domain

type BeerRepository interface {
	FindAll() ([]Beer, error)
	FindByID(id int) (*Beer, error)
	CreateBeer(beer *Beer) (int, error)
	FindByName(name string) (bool, error)
}

type Beer struct {
	id       int
	name     string
	brewery  string
	country  string
	price    float32
	currency string
}

func NewBeer(id int, name string, brewery string, country string, price float32, currency string) *Beer {
	return &Beer{
		id:       id,
		name:     name,
		brewery:  brewery,
		country:  country,
		price:    price,
		currency: currency,
	}
}

func (b Beer) ID() int {
	return b.id
}

func (b *Beer) Name() string {
	return b.name
}

func (b *Beer) Brewery() string {
	return b.brewery
}

func (b *Beer) Country() string {
	return b.country
}

func (b *Beer) Price() float32 {
	return b.price
}

func (b *Beer) Currency() string {
	return b.currency
}
