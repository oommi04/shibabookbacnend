package customerUsecase

import (
	"context"
	"github.com/oommi04/shibabookbackend/src/domains/customerDomain"
	_customerRepository "github.com/oommi04/shibabookbackend/src/repository/customerRepository"
	"time"
)

type CustomerUsecaseInterface interface {
	GetByIdentificationNumber(ctx context.Context, idn string) (*customerDomain.Customer, error)
	Register(ctx context.Context, info *customerDomain.Customer) error
	GetById(ctx context.Context, id string) (*customerDomain.Customer, error)
}

type customerUsecase struct {
	customerRepo   _customerRepository.CustomerRepositoryInterface
	contextTimeout time.Duration
}

func New(p _customerRepository.CustomerRepositoryInterface, timout time.Duration) CustomerUsecaseInterface {
	return &customerUsecase{p, timout}
}

func (p *customerUsecase) GetByIdentificationNumber(ctx context.Context, idn string) (*customerDomain.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	resp, err := p.customerRepo.GetByIdentificationNumber(ctx, idn)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *customerUsecase) GetById(ctx context.Context, id string) (*customerDomain.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	resp, err := p.customerRepo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *customerUsecase) Register(ctx context.Context, info *customerDomain.Customer) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	checkCustomerExist, err := p.GetByIdentificationNumber(ctx, info.IdentificationNumber)

	if err != customerDomain.ErrorIdentificationNotFound && checkCustomerExist.IdentificationNumber != "" {
		return customerDomain.ErrorIdentificationNumberWasUsed
	}

	err = p.customerRepo.Save(ctx, info)

	if err != nil {
		return err
	}

	return nil
}
