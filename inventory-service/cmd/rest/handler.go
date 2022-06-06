package rest

import (
	"encoding/json"
	"inventory-service/internal/use_case"
	"lib"
	"net/http"
)

type Handler interface {
	CreateProduct(http.ResponseWriter, *http.Request)
	GetProduct(http.ResponseWriter, *http.Request)
	AddProductStock(http.ResponseWriter, *http.Request)
}

type handler struct {
	inventoryService use_case.InventoryService
}

func NewHandler(inventoryService use_case.InventoryService) Handler {
	return &handler{inventoryService: inventoryService}
}

func (h handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var request use_case.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	productId, err := h.inventoryService.CreateProduct(request)

	lib.WriteResponse(w, err, productId)
}

func (h handler) GetProduct(w http.ResponseWriter, _ *http.Request) {
	products, err := h.inventoryService.GetAllProducts()
	lib.WriteResponse(w, err, products)
}

func (h handler) AddProductStock(w http.ResponseWriter, r *http.Request) {
	var request use_case.AddStockRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	err = h.inventoryService.AddStock(request)
	lib.WriteResponse(w, err, nil)
}
