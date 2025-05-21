package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCalculateTaxAndSave(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockRepository(ctrl)

	// Espera que SaveTax(10.0) seja chamado 2 vezes, retornando nil
	mockRepo.EXPECT().SaveTax(10.0).Return(nil).Times(2)

	// Espera que SaveTax(0.0) seja chamado e retorne erro
	mockRepo.EXPECT().SaveTax(0.0).Return(errors.New("error saving tax"))

	// Primeira chamada
	err := CalculateTaxAndSave(1000.0, mockRepo)
	assert.NoError(t, err)

	// Segunda chamada
	err = CalculateTaxAndSave(1000.0, mockRepo)
	assert.NoError(t, err)

	// Terceira chamada, erro
	err = CalculateTaxAndSave(0.0, mockRepo)
	assert.Error(t, err)
}
