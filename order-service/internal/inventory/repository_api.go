package inventory

import "github.com/google/uuid"

type ApiRepository struct {
	productAvailability map[uuid.UUID]ProductDetail
}

func NewApiRepository() *ApiRepository {
	return &ApiRepository{
		productAvailability: map[uuid.UUID]ProductDetail{},
	}
}

func (repo ApiRepository) GetProductAvailability(productIds []uuid.UUID) ([]ProductDetail, error) {
	var productDetails []ProductDetail
	for _, productId := range productIds {
		productDetails = append(productDetails, repo.productAvailability[productId])
	}
	return productDetails, nil
}
