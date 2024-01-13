package entity

import (
	"errors"
	"time"

	"github.com/alcir-junior-caju/study-go-expert-client-apis/pkg/valueObject"
)

var (
	ErrorIDIsRequired    = errors.New("id is required")
	ErrorInvalidID       = errors.New("invalid id")
	ErrorNameIsRequired  = errors.New("name is required")
	ErrorPriceIsRequired = errors.New("price is required")
	ErrorInvalidPrice    = errors.New("invalid price")
)

type ProductStruct struct {
	ID        valueObject.IDType `json:"id"`
	Name      string             `json:"name"`
	Price     float64            `json:"price"`
	CreatedAt time.Time          `json:"created_at"`
}

func Product(name string, price float64) (*ProductStruct, error) {
	product := &ProductStruct{
		ID:        valueObject.ID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	errorProduct := product.Validate()
	if errorProduct != nil {
		return nil, errorProduct
	}
	return product, nil
}

func (product *ProductStruct) Validate() error {
	if product.ID.String() == "" {
		return ErrorIDIsRequired
	}
	if _, errorValidID := valueObject.ParseID(product.ID.String()); errorValidID != nil {
		return ErrorInvalidID
	}
	if product.Name == "" {
		return ErrorNameIsRequired
	}
	if product.Price == 0 {
		return ErrorPriceIsRequired
	}
	if product.Price < 0 {
		return ErrorInvalidPrice
	}
	return nil
}
