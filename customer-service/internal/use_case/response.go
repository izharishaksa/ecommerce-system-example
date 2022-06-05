package use_case

import "github.com/google/uuid"

type CustomerDetail struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Balance float64   `json:"balance"`
}
