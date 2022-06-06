package rest

import (
	"encoding/json"
	"inventory-service/internal/use_case"
	"lib"
	"net/http"
)

type Handler struct {
	inventoryService use_case.InventoryServiceClient
}

func NewHandler(inventoryService use_case.InventoryServiceClient) *Handler {
	return &Handler{inventoryService: inventoryService}
}

func (h Handler) createProduct(w http.ResponseWriter, r *http.Request) {
	var request use_case.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	productId, err := h.inventoryService.CreateProduct(request)

	lib.WriteResponse(w, err, productId)
}

func (h Handler) getProduct(w http.ResponseWriter, _ *http.Request) {
	products, err := h.inventoryService.GetAllProducts()
	lib.WriteResponse(w, err, products)
}

func (h Handler) addProductStock(w http.ResponseWriter, r *http.Request) {
	var request use_case.AddStockRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	err = h.inventoryService.AddStock(request)
	lib.WriteResponse(w, err, nil)
}
