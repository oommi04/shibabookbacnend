package customerRepository

import (
	"context"
	"github.com/oommi04/shibabookbackend/src/domains/customerDomain"
)

type CustomerRepositoryInterface interface {
	GetByIdentificationNumber(ctx context.Context, idn string) (*customerDomain.Customer, error)
	Save(ctx context.Context, info *customerDomain.Customer) error
	GetById(ctx context.Context, id string) (*customerDomain.Customer, error)
}
