package dao

import "falabella/internal/platform/storage/mysql/model"

type IDao interface {
	Get() ([]model.Beer, error)
	GetById(id int) (*model.Beer, error)
	GetByName(name string) (bool, error)
	Create(beer *model.Beer) (int, error)
}
