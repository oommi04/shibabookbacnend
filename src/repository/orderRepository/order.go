package orderRepository

import (
	"context"
	"github.com/oommi04/shibabookbackend/src/domains/orderDomain"
)

type OrderRepositoryInterface interface {
	List(ctx context.Context) ([]*orderDomain.Order, error)
	Save(ctx context.Context, info *orderDomain.Order) error
	CheckOut(ctx context.Context, info *orderDomain.Order, id string) error
	GetById(ctx context.Context, id string) (*orderDomain.Order, error)
}
