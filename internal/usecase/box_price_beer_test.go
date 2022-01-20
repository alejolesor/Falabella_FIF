package usecase

import (
	"falabella/internal/domain"
	"falabella/internal/platform/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var mockRepositoryCurrency mocks.CurrencyRepository
var mockRepositoryDb mocks.BeerRepository

func TestCalculateBuyBox(t *testing.T) {

	beer := domain.Beer{}
	convert := map[string]float32{}
	tests := []struct {
		name    string
		mock    func()
		wantErr bool
	}{
		{
			name: "Success Method FindByID",
			mock: func() {
				mockRepositoryDb.On("FindByID", 1).Return(&beer, nil)
				mockRepositoryCurrency.On("ConvertToUSD", mock.Anything).Return(convert, nil)
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepositoryCurrency = mocks.CurrencyRepository{}
			mockRepositoryDb = mocks.BeerRepository{}
			tt.mock()
			useCase := NewBoxPriceBeer(&mockRepositoryCurrency, &mockRepositoryDb)
			_, err := useCase.Calculate(1, "COP", 6)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
