package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(test *testing.T) {
	product, error := Product("Product 1", 10)
	assert.Nil(test, error)
	assert.NotNil(test, product)
	assert.NotEmpty(test, product.ID)
	assert.Equal(test, "Product 1", product.Name)
	assert.Equal(test, 10.0, product.Price)
}

func TestProductWhenNameIsEmpty(test *testing.T) {
	product, error := Product("", 10)
	assert.Nil(test, product)
	assert.Equal(test, ErrorNameIsRequired, error)
}

func TestProductWhenPriceIsZero(test *testing.T) {
	product, error := Product("Product 1", 0)
	assert.Nil(test, product)
	assert.Equal(test, ErrorPriceIsRequired, error)
}

func TestProductWhenPriceIsNegative(test *testing.T) {
	product, error := Product("Product 1", -1)
	assert.Nil(test, product)
	assert.Equal(test, ErrorInvalidPrice, error)
}

func TestProductValidate(test *testing.T) {
	product, error := Product("Product 1", 10)
	assert.Nil(test, error)
	assert.NotNil(test, product)
	assert.Nil(test, product.Validate())
}
