package transactions

import "github.com/google/uuid"

type Transaction struct {
	Id        uuid.UUID `json:"id"`
	Amount    float64   `json:"amount" validate:"required"`
	Currency  string    `json:"currency" validate:"required"`
	Createdat string    `json:"createdAt"`
	Status    bool      `json:"status"`
}
