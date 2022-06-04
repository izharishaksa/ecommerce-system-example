package main

import (
	"encoding/json"
	"inventory-service/internal/use_case"
	"lib"
	"net/http"
)

type ProductService interface {
	CreateProduct(request use_case.CreateProductRequest) error
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
	err = h.productService.CreateProduct(request)
	lib.WriteResponse(w, err, http.StatusCreated, nil)
}

func (h Handler) getProduct(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
