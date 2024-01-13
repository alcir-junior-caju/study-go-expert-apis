package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alcir-junior-caju/study-go-expert-client-apis/internal/dto"
	"github.com/alcir-junior-caju/study-go-expert-client-apis/internal/entity"
	"github.com/alcir-junior-caju/study-go-expert-client-apis/internal/infra/database"
	"github.com/alcir-junior-caju/study-go-expert-client-apis/pkg/valueObject"
	"github.com/go-chi/chi/v5"
)

type ProductHandlerStruct struct {
	ProductDatabase database.ProductInterface
}

func ProductHandler(database database.ProductInterface) *ProductHandlerStruct {
	return &ProductHandlerStruct{
		ProductDatabase: database,
	}
}

// CreateProduct godoc
// @Summary      Create product
// @Description  Create products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateProductInput  true  "product request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /products [post]
// @Security ApiKeyAuth
func (productHandler *ProductHandlerStruct) CreateProduct(writer http.ResponseWriter, request *http.Request) {
	var product dto.CreateProductInput
	errorDecode := json.NewDecoder(request.Body).Decode(&product)
	if errorDecode != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorDecode.Error())
		return
	}
	productCreated, errorCreated := entity.Product(product.Name, product.Price)
	if errorCreated != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorCreated.Error())
		return
	}
	errorDatabase := productHandler.ProductDatabase.Create(productCreated)
	if errorDatabase != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorDatabase.Error())
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

// FindProducts godoc
// @Summary      List products
// @Description  get all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page      query     string  false  "page number"
// @Param        limit     query     string  false  "limit"
// @Param        sort      query     string  false  "sort"
// @Success      200       {array}   entity.ProductStruct
// @Failure      404       {object}  Error
// @Failure      500       {object}  Error
// @Router       /products [get]
// @Security ApiKeyAuth
func (productHandler *ProductHandlerStruct) FindProducts(writer http.ResponseWriter, request *http.Request) {
	page := request.URL.Query().Get("page")
	limit := request.URL.Query().Get("limit")
	sort := request.URL.Query().Get("sort")
	pageInt, errorPage := strconv.Atoi(page)
	if errorPage != nil {
		pageInt = 0
	}
	limitInt, errorLimit := strconv.Atoi(limit)
	if errorLimit != nil {
		limitInt = 0
	}
	products, errorDatabase := productHandler.ProductDatabase.FindAll(pageInt, limitInt, sort)
	if errorDatabase != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorDatabase.Error())
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(products)
}

// GetProduct godoc
// @Summary      Get a product
// @Description  Get a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "product ID" Format(uuid)
// @Success      200  {object}  entity.ProductStruct
// @Failure      404
// @Failure      500  {object}  Error
// @Router       /products/{id} [get]
// @Security ApiKeyAuth
func (productHandler *ProductHandlerStruct) FindProduct(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode("id is required")
		return
	}
	product, errorDatabase := productHandler.ProductDatabase.FindById(id)
	if errorDatabase != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(errorDatabase.Error())
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(product)
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Update a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id        	path      string                  true  "product ID" Format(uuid)
// @Param        request     body      dto.CreateProductInput  true  "product request"
// @Success      200
// @Failure      404
// @Failure      500       {object}  Error
// @Router       /products/{id} [put]
// @Security ApiKeyAuth
func (productHandler *ProductHandlerStruct) UpdateProduct(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode("id is required")
		return
	}
	var product entity.ProductStruct
	errorDecode := json.NewDecoder(request.Body).Decode(&product)
	if errorDecode != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorDecode.Error())
		return
	}
	product.ID, errorDecode = valueObject.ParseID(id)
	if errorDecode != nil {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(errorDecode.Error())
		return
	}
	_, errorFound := productHandler.ProductDatabase.FindById(id)
	if errorFound != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(errorFound.Error())
		return
	}
	errorDatabase := productHandler.ProductDatabase.Update(&product)
	if errorDatabase != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorDatabase.Error())
		return
	}
	writer.WriteHeader(http.StatusOK)
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id        path      string                  true  "product ID" Format(uuid)
// @Success      200
// @Failure      404
// @Failure      500       {object}  Error
// @Router       /products/{id} [delete]
// @Security ApiKeyAuth
func (productHandler *ProductHandlerStruct) DeleteProduct(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode("id is required")
		return
	}
	_, errorFound := productHandler.ProductDatabase.FindById(id)
	if errorFound != nil {
		writer.WriteHeader(http.StatusNotFound)
		json.NewEncoder(writer).Encode(errorFound.Error())
		return
	}
	errorDatabase := productHandler.ProductDatabase.Delete(id)
	if errorDatabase != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(errorDatabase.Error())
		return
	}
	writer.WriteHeader(http.StatusOK)
}
