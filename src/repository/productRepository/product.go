package productRepository

import (
	"context"
	"github.com/oommi04/shibabookbackend/src/domains/productDomain"
)

type ProductRepositoryInterface interface {
	List(ctx context.Context) ([]*productDomain.Product, error)
	Save(ctx context.Context, info *productDomain.Product) error
	GetByNameAndFrom(ctx context.Context, name string, from string) (*productDomain.Product, error)
	ListByNameAndFrom(ctx context.Context, name string, from string) ([]*productDomain.Product, error)
}
