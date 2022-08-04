package order

const (
	Placed   Status = "placed"
	Created  Status = "created"
	Rejected Status = "rejected"

	PlacedEvent   EventType = "order_placed"
	RejectedEvent EventType = "order_rejected"
	CreatedEvent  EventType = "order_created"
)
