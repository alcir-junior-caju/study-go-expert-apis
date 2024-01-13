package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/alcir-junior-caju/study-go-expert-client-apis/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestProductDatabaseCreate(test *testing.T) {
	database, errorDatabase := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if errorDatabase != nil {
		test.Error("Error to connect with database")
	}
	database.AutoMigrate(&entity.ProductStruct{})
	product, _ := entity.Product("Product 1", 10.00)
	productDatabase := ProductDatabase(database)
	errorCreate := productDatabase.Create(product)
	assert.Nil(test, errorCreate)
	var productFound entity.ProductStruct
	errorFound := database.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(test, errorFound)
	assert.Equal(test, product.ID, productFound.ID)
	assert.Equal(test, product.Name, productFound.Name)
	assert.Equal(test, product.Price, productFound.Price)
}

func TestProductDatabaseFindAll(test *testing.T) {
	database, errorDatabase := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if errorDatabase != nil {
		test.Error("Error to connect with database")
	}
	database.AutoMigrate(&entity.ProductStruct{})
	for i := 1; i < 24; i++ {
		product, errorProduct := entity.Product(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(test, errorProduct)
		database.Create(product)
	}
	productDatabase := ProductDatabase(database)
	products, errorProducts := productDatabase.FindAll(1, 10, "asc")
	assert.NoError(test, errorProducts)
	assert.Len(test, products, 10)
	assert.Equal(test, "Product 1", products[0].Name)
	assert.Equal(test, "Product 10", products[9].Name)

	products, errorProducts = productDatabase.FindAll(2, 10, "asc")
	assert.NoError(test, errorProducts)
	assert.Len(test, products, 10)
	assert.Equal(test, "Product 11", products[0].Name)
	assert.Equal(test, "Product 20", products[9].Name)

	products, errorProducts = productDatabase.FindAll(3, 10, "asc")
	assert.NoError(test, errorProducts)
	assert.Len(test, products, 3)
	assert.Equal(test, "Product 21", products[0].Name)
	assert.Equal(test, "Product 23", products[2].Name)
}

func TestProductDatabaseFindById(test *testing.T) {
	database, errorDatabase := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if errorDatabase != nil {
		test.Error("Error to connect with database")
	}
	database.AutoMigrate(&entity.ProductStruct{})
	product, _ := entity.Product("Product 1", 10.00)
	database.Create(product)
	productDatabase := ProductDatabase(database)
	productFound, errorFound := productDatabase.FindById(product.ID.String())
	assert.Nil(test, errorFound)
	assert.Equal(test, product.ID, productFound.ID)
	assert.Equal(test, product.Name, productFound.Name)
	assert.Equal(test, product.Price, productFound.Price)
}

func TestProductDatabaseUpdate(test *testing.T) {
	database, errorDatabase := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if errorDatabase != nil {
		test.Error("Error to connect with database")
	}
	database.AutoMigrate(&entity.ProductStruct{})
	product, _ := entity.Product("Product 1", 10.00)
	database.Create(product)
	product.Name = "Product 2"
	product.Price = 20.00
	productDatabase := ProductDatabase(database)
	errorUpdate := productDatabase.Update(product)
	assert.Nil(test, errorUpdate)
	var productFound entity.ProductStruct
	errorFound := database.First(&productFound, "id = ?", product.ID).Error
	assert.Nil(test, errorFound)
	assert.Equal(test, product.ID, productFound.ID)
	assert.Equal(test, product.Name, productFound.Name)
	assert.Equal(test, product.Price, productFound.Price)
}

func TestProductDatabaseDelete(test *testing.T) {
	database, errorDatabase := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if errorDatabase != nil {
		test.Error("Error to connect with database")
	}
	database.AutoMigrate(&entity.ProductStruct{})
	product, _ := entity.Product("Product 1", 10.00)
	database.Create(product)
	productDatabase := ProductDatabase(database)
	errorDelete := productDatabase.Delete(product.ID.String())
	assert.Nil(test, errorDelete)
	var productFound entity.ProductStruct
	errorFound := database.First(&productFound, "id = ?", product.ID).Error
	assert.Error(test, errorFound)
}
