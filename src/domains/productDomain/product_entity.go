package productDomain

import "time"

type Product struct {
	ID            string     `json:"_id,omitempty"`
	Name          string     `json:"name,omitempty"`
	Price         float32    `json:"price,omitempty"`
	QuantityStock string     `json:"quantityStock,omitempty"`
	Description   string     `json:"description,omitempty"`
	From          string     `json:"from,omitempty"`
	Status        string     `json:"status,omitempty"`
	DateTime      *time.Time `json:"dateTime,omitempty"`
	Image         string     `json:"image,omitempty"`
	BarCode       string     `json:"barCode,omitempty"`
}
