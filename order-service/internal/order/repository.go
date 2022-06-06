package order

type Repository interface {
	SaveOrder(order *Order) error
}
