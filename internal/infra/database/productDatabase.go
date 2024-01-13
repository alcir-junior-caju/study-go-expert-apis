package database

import (
	"github.com/alcir-junior-caju/study-go-expert-client-apis/internal/entity"
	"gorm.io/gorm"
)

type ProductDatabaseStruct struct {
	Connection *gorm.DB
}

func ProductDatabase(connection *gorm.DB) *ProductDatabaseStruct {
	return &ProductDatabaseStruct{Connection: connection}
}

func (productDatabase *ProductDatabaseStruct) Create(product *entity.ProductStruct) error {
	return productDatabase.Connection.Create(product).Error
}

func (productDatabase *ProductDatabaseStruct) FindAll(page, limit int, sort string) ([]entity.ProductStruct, error) {
	var products []entity.ProductStruct
	var errorProduct error
	if sort == "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		errorProduct = productDatabase.Connection.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		errorProduct = productDatabase.Connection.Order("created_at " + sort).Find(&products).Error
	}
	if errorProduct != nil {
		return nil, errorProduct
	}
	return products, nil
}

func (productDatabase *ProductDatabaseStruct) FindById(id string) (*entity.ProductStruct, error) {
	var product entity.ProductStruct
	if errorProduct := productDatabase.Connection.First(&product, "id = ?", id).Error; errorProduct != nil {
		return nil, errorProduct
	}
	return &product, nil
}

func (productDatabase *ProductDatabaseStruct) Update(product *entity.ProductStruct) error {
	_, errorProduct := productDatabase.FindById(product.ID.String())
	if errorProduct != nil {
		return errorProduct
	}
	return productDatabase.Connection.Save(product).Error
}

func (productDatabase *ProductDatabaseStruct) Delete(id string) error {
	product, errorProduct := productDatabase.FindById(id)
	if errorProduct != nil {
		return errorProduct
	}
	return productDatabase.Connection.Delete(product).Error
}
