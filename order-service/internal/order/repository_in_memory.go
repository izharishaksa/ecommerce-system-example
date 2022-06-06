package order

type inMemoryRepository struct {
	orders map[string]*Order
}

func NewInMemoryRepository() Repository {
	return &inMemoryRepository{
		orders: make(map[string]*Order),
	}
}

func (repo inMemoryRepository) SaveOrder(order *Order) error {
	repo.orders[order.Id.String()] = order
	return nil
}
