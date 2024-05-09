// Product represents a product entity.
package model

import "github.com/google/uuid"

type Product struct {
	ID          uuid.UUID `json:"product_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
}
