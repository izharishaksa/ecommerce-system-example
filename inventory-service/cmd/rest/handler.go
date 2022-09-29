package rest

import (
	"encoding/json"
	"github.com/google/uuid"
	"inventory-service/internal/use_case"
	"lib"
	"net/http"
)

type InventoryService interface {
	CreateProduct(request use_case.CreateProductRequest) (*uuid.UUID, error)
	GetAllProducts() ([]use_case.ProductDetail, error)
	AddStock(request use_case.AddStockRequest) error
}

type handlerImpl struct {
	inventoryService InventoryService
}

func NewHandler(inventoryService InventoryService) *handlerImpl {
	return &handlerImpl{inventoryService: inventoryService}
}

func (h handlerImpl) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var request use_case.CreateProductRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	productId, err := h.inventoryService.CreateProduct(request)

	lib.WriteResponse(w, err, productId)
}

func (h handlerImpl) GetProduct(w http.ResponseWriter, _ *http.Request) {
	products, err := h.inventoryService.GetAllProducts()
	lib.WriteResponse(w, err, products)
}

func (h handlerImpl) AddProductStock(w http.ResponseWriter, r *http.Request) {
	var request use_case.AddStockRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		lib.WriteResponse(w, lib.NewErrBadRequest(err.Error()), nil)
		return
	}
	err = h.inventoryService.AddStock(request)
	lib.WriteResponse(w, err, nil)
}
