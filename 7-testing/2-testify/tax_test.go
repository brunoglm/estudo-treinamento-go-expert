package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax := CalculateTax(1000.0)
	assert.Equal(t, 10.0, tax)
}
