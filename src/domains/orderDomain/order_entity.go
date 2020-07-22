package orderDomain

import (
	"github.com/oommi04/backendtest/src/domains/customerDomain"
	"github.com/oommi04/backendtest/src/domains/invoiceDomain"
	"github.com/oommi04/backendtest/src/domains/productDomain"
	"github.com/oommi04/backendtest/src/domains/staffDomain"
	"time"
)

type ProductAmount struct {
	Product productDomain.Product `json:"product,omitempty"`
	Amount int `json:"amount,omitempty"`
}


type Order struct {
	ID          string                    `json:"_id,omitempty"`
	Status      string                    `json:"status,omitempty"`
	Description string                    `json:"description,omitempty"`
	Products    []*ProductAmount `json:"products,omitempty"`
	Staff      staffDomain.Staff        `json:"Staff,omitempty"`
	Customer   customerDomain.Customer  `json:"customer,omitempty"`
	Invoice invoiceDomain.Invoice `json:"invoice,omitempty"`
	TotalPrice float32 `json:"totalPrice,omitempty"`
	Discount float32 `json:"discount,omitempty"`
	DateTime *time.Time `json:"time,omitempty"`
	NET float32 `json:"net,omitempty"`
	//paymentId
}