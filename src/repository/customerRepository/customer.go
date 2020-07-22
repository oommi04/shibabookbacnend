package customerRepository

import (
	"github.com/tkhamsila/shibabookbackend/src/domains/customerDomain"
	"context"
)

type CustomerRepositoryInterface interface {
	GetByIdentificationNumber(ctx context.Context, idn string) (*customerDomain.Customer, error)
	Save(ctx context.Context, info *customerDomain.Customer) error
	GetById(ctx context.Context,id string) (*customerDomain.Customer, error)
}