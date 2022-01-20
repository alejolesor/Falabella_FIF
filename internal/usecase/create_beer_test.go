package usecase

import (
	"errors"
	"falabella/internal/domain"
	"falabella/internal/platform/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockRepositoryBeerCreate mocks.BeerRepository

func TestCreateBeer(t *testing.T) {

	tests := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			name: "Success Method FindByName",
			mock: func() {
				mockRepositoryBeerCreate.On("FindByName", mock.Anything).Return(false, nil)
				mockRepositoryBeerCreate.On("CreateBeer", mock.Anything).Return(1, nil)
			},
			wantErr: false,
		},
		{
			name: "Error Method CreateBeer",
			mock: func() {
				mockRepositoryBeerCreate.On("FindByName", mock.Anything).Return(false, nil)
				mockRepositoryBeerCreate.On("CreateBeer", mock.Anything).Return(0, errors.New("error create"))
			},
			wantErr: true,
		},
		{
			name: "Error Method FindByName",
			mock: func() {
				mockRepositoryBeerCreate.On("FindByName", mock.Anything).Return(false, errors.New("error find by name"))
				mockRepositoryBeerCreate.On("CreateBeer", mock.Anything).Return(1, nil)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepositoryBeerCreate = mocks.BeerRepository{}
			tt.mock()
			useCase := NewCreateBeer(&mockRepositoryBeerCreate)
			_, err := useCase.CreateBeer(domain.NewBeer(1, "Aguila", "Bavaria", "Colombia", 2500, "COP"))
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
