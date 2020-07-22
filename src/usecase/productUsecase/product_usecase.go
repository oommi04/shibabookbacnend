package productUsecase

import (
	"context"
	"github.com/tkhamsila/shibabookbackend/src/domains/productDomain"
	_productRepository "github.com/tkhamsila/shibabookbackend/src/repository/productRepository"
	"time"
)

type ProductUsecaseInterface interface {
	List(ctx context.Context) ([]*productDomain.Product, error)
	Save(ctx context.Context, info *productDomain.Product) error
}

type productUsecase struct {
	productRepo   _productRepository.ProductRepositoryInterface
	contextTimeout time.Duration
}

func NewProductUsecase(p _productRepository.ProductRepositoryInterface, timout time.Duration) ProductUsecaseInterface {

	return &productUsecase{p, timout}
}

func (p *productUsecase) List(ctx context.Context) ([]*productDomain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	resps, err := p.productRepo.List(ctx)

	if err != nil {
		return nil, err
	}

	return resps, nil
}

func (p *productUsecase) Save(ctx context.Context, info *productDomain.Product) error{
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	checkProductExist, err := p.productRepo.GetByNameAndFrom(ctx, info.Name, info.From)

	if err != nil && err != productDomain.ErrorProductNotFoundByNameAndFrom {
		return err
	}

	if checkProductExist.Name != "" {
		return productDomain.ErrorProductExist
	}

	err = p.productRepo.Save(ctx, info)

	if err != nil {
		return err
	}

	return nil
}
