package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"inventory-service/internal/use_case"
	"lib"
	"net/http"
)

type ProductService interface {
	CreateProduct(request use_case.CreateProductRequest) (*uuid.UUID, error)
	GetAllProducts() ([]use_case.ProductDetail, error)
}

type Handler struct {
	productService ProductService
}

func NewHandler(productService ProductService) *Handler {
	return &Handler{productService: productService}
}

func (h Handler) createProduct(w http.ResponseWriter, r *http.Request) {
	var request use_case.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), http.StatusBadRequest, nil)
		return
	}
	productId, err := h.productService.CreateProduct(request)

	lib.WriteResponse(w, err, http.StatusCreated, productId)
}

func (h Handler) getProduct(w http.ResponseWriter, r *http.Request) {
	products, err := h.productService.GetAllProducts()
	lib.WriteResponse(w, err, http.StatusOK, products)
}
