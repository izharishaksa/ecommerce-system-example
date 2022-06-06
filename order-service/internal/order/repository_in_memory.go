package order

type InMemoryRepository struct {
	orders map[string]*Order
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		orders: make(map[string]*Order),
	}
}

func (repo InMemoryRepository) SaveOrder(order *Order) error {
	//TODO implement me
	panic("implement me")
}
