package main

import (
	"inventory-service/internal/use_case"
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
	w.Write([]byte("ok"))
}
