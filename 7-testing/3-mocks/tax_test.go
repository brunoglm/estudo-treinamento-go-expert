package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type TaxRepositoryMock struct {
	mock.Mock
}

func (m *TaxRepositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0)
}

func TestCalculateTaxAndSave(t *testing.T) {
	repo := &TaxRepositoryMock{}
	repo.On("SaveTax", 10.0).Return(nil)
	// repo.On("SaveTax", 10.0).Return(nil).Once()
	repo.On("SaveTax", 0.0).Return(errors.New("error saving tax"))
	// repo.On("SaveTax", mock.Anything).Return(errors.New("error saving tax"))

	err := CalculateTaxAndSave(1000.0, repo)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(1000.0, repo)
	assert.Nil(t, err)

	err = CalculateTaxAndSave(0.0, repo)
	assert.Error(t, err, "error saving tax")

	repo.AssertExpectations(t)
	// repo.AssertNumberOfCalls(t, "SaveTax", 2)
}
