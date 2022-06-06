package inventory

import "github.com/google/uuid"

type apiRepository struct {
	productAvailability map[uuid.UUID]ProductDetail
}

func NewApiRepository() Repository {
	return &apiRepository{
		productAvailability: map[uuid.UUID]ProductDetail{},
	}
}

func (repo apiRepository) GetProductAvailability(productIds []uuid.UUID) ([]ProductDetail, error) {
	var productDetails []ProductDetail
	for _, productId := range productIds {
		productDetails = append(productDetails, repo.productAvailability[productId])
	}
	return productDetails, nil
}
