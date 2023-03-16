package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 10.0)

	assert.NoError(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
	assert.NotEmpty(t, p.CreatedAt)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 10.0, p.Price)
}

func TestProductWenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10.0)

	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Product", 0.0)

	assert.Nil(t, p)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWenPriceIsLessThanZero(t *testing.T) {
	p, err := NewProduct("Product", -10.0)

	assert.Nil(t, p)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductValidate(t *testing.T) {
	p, err := NewProduct("Product", 10.0)

	assert.NoError(t, err)
	assert.NotNil(t, p)

	err = p.Validate()

	assert.NoError(t, err)
}
