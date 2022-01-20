package usecase

import (
	"errors"
	"falabella/internal/domain"
	"falabella/internal/platform/mocks"

	"testing"

	"github.com/stretchr/testify/assert"
)

var mockRepositoryBeer mocks.BeerRepository

func TestFindAll(t *testing.T) {
	arrayBeer := []domain.Beer{}
	tests := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			name: "Success Method FindAll",
			mock: func() {
				mockRepositoryBeer.On("FindAll").Return(arrayBeer, nil)
			},
			wantErr: false,
		},
		{
			name: "Error Method FindAll",
			mock: func() {
				mockRepositoryBeer.On("FindAll").Return(nil, errors.New("error get list beer"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepositoryBeer = mocks.BeerRepository{}
			tt.mock()
			useCase := NewFindBeer(&mockRepositoryBeer)
			_, err := useCase.FindAll()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestFindById(t *testing.T) {
	beer := domain.Beer{}
	tests := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			name: "Success Method FindByID",
			mock: func() {
				mockRepositoryBeer.On("FindByID", 1).Return(&beer, nil)
			},
			wantErr: false,
		},
		{
			name: "Error Method FindByID",
			mock: func() {
				mockRepositoryBeer.On("FindByID", 1).Return(nil, errors.New("error get by id"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepositoryBeer = mocks.BeerRepository{}
			tt.mock()
			useCase := NewFindBeer(&mockRepositoryBeer)
			_, err := useCase.FindById(1)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
