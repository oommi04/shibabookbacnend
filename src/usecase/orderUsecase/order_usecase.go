package orderUsecase

import (
	"context"
	"github.com/oommi04/shibabookbackend/src/domains/invoiceDomain"
	"github.com/oommi04/shibabookbackend/src/domains/orderDomain"
	_invoiceRepository "github.com/oommi04/shibabookbackend/src/repository/invoiceRepository"
	_orderRepository "github.com/oommi04/shibabookbackend/src/repository/orderRepository"
	_productRepository "github.com/oommi04/shibabookbackend/src/repository/productRepository"
	_customerUsecase "github.com/oommi04/shibabookbackend/src/usecase/customerUsecase"
	"github.com/oommi04/shibabookbackend/src/utils/common"
	"time"
)

type OrderUsecaseInterface interface {
	List(ctx context.Context) ([]*orderDomain.Order, error)
	Save(ctx context.Context, info *orderDomain.Order) error
	CheckOut(ctx context.Context, id string) (*orderDomain.Order, error)
}

type orderUsecase struct {
	orderRepo       _orderRepository.OrderRepositoryInterface
	productRepo     _productRepository.ProductRepositoryInterface
	invoiceRepo     _invoiceRepository.InvoiceRepositoryInterface
	customerUsecase _customerUsecase.CustomerUsecaseInterface
	contextTimeout  time.Duration
}

func NewProductUsecase(o _orderRepository.OrderRepositoryInterface, p _productRepository.ProductRepositoryInterface, i _invoiceRepository.InvoiceRepositoryInterface, c _customerUsecase.CustomerUsecaseInterface, timout time.Duration) OrderUsecaseInterface {
	return &orderUsecase{o, p, i, c, timout}
}

func (p *orderUsecase) DiscountHarryBook(ctx context.Context, hs []*orderDomain.ProductAmount) (float32, float32, float32) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	// get all id harrybook
	resps, _ := p.productRepo.ListByNameAndFrom(ctx, "Harry", "HarryShop")
	harryBooksId := []string{}
	for _, item := range resps {
		harryBooksId = append(harryBooksId, item.ID)
	}

	// get harrybook from basket
	var totalPrice float32 = 0
	harryBooksAmount := []int{}
	harryBooksPrice := []float32{}
	for _, item := range hs {
		price := float32(item.Amount) * item.Product.Price
		totalPrice = totalPrice + price

		if common.IncludeString(harryBooksId, item.Product.ID) {
			harryBooksAmount = append(harryBooksAmount, item.Amount)
			harryBooksPrice = append(harryBooksPrice, item.Product.Price)

		}

	}

	// find num of uniq book
	bookUniq := [][]float32{}
	maxAmountHarryBook := common.MaxIntSlice(harryBooksAmount)
	for i := 0; i < maxAmountHarryBook; i++ {
		bookUniqAmountPrice := []float32{}
		for j := range harryBooksAmount {
			if harryBooksAmount[j] > 0 {
				bookUniqAmountPrice = append(bookUniqAmountPrice, harryBooksPrice[j])
				harryBooksAmount[j] = harryBooksAmount[j] - 1
			}
		}
		bookUniq = append(bookUniq, bookUniqAmountPrice)
	}

	//find discount
	var totalDiscount float32 = 0
	for _, item := range bookUniq {

		if len(item) > 1 {
			discountPer := 10
			discountPer = discountPer + len(item) - 2

			price := common.SumFloat32Slice(item)
			priceWithDiscount := float32(price) * float32(discountPer) * 0.01
			totalDiscount = totalDiscount + priceWithDiscount

		}
	}
	net := totalPrice - totalDiscount
	return net, totalPrice, totalDiscount
}

func (p *orderUsecase) List(ctx context.Context) ([]*orderDomain.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	resps, err := p.orderRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	return resps, nil
}

func (p *orderUsecase) Save(ctx context.Context, info *orderDomain.Order, ) error {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	info.Status = "PENDING"

	if info.Customer.IdentificationNumber != "" {
		customer, err := p.customerUsecase.GetByIdentificationNumber(ctx, info.Customer.IdentificationNumber)

		if err != nil {
			return err
		}
		info.Customer = *customer

	}

	info.NET, info.TotalPrice, info.Discount = p.DiscountHarryBook(ctx, info.Products)

	err := p.orderRepo.Save(ctx, info)

	if err != nil {
		return err
	}

	return nil
}

func (p *orderUsecase) CheckOut(ctx context.Context, id string) (*orderDomain.Order, error) {
	ctx, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()

	data, err := p.orderRepo.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	//create invoice
	if data.Customer.ID != "" {
		customer, err := p.customerUsecase.GetById(ctx, data.Customer.ID)
		if err != nil {
			return nil, err
		}
		data.Customer = *customer

		data.Invoice = invoiceDomain.Invoice{
			Path: "google.com",
			Mail: data.Customer.Mail,
		}
		err = p.invoiceRepo.Save(ctx, &data.Invoice)

		if err != nil {
			return nil, err
		}

	} else {
		data.Invoice = invoiceDomain.Invoice{
			Path: "google.com",
			Mail: "oommi04@gmail.com",
		}
	}

	data.Status = "DONE"
	err = p.orderRepo.CheckOut(ctx, data, id)

	if err != nil {
		return nil, err
	}

	return data, nil
}
