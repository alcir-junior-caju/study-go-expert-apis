package database

import "github.com/alcir-junior-caju/study-go-expert-client-apis/internal/entity"

type UserInterface interface {
	Create(user *entity.UserStruct) error
	FindByEmail(email string) (*entity.UserStruct, error)
}

type ProductInterface interface {
	Create(product *entity.ProductStruct) error
	FindAll(page, limit int, sort string) ([]entity.ProductStruct, error)
	FindById(id string) (*entity.ProductStruct, error)
	Update(product *entity.ProductStruct) error
	Delete(id string) error
}
