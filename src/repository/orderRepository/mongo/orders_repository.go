package mongo

import (
	"context"
	"github.com/oommi04/shibabookbackend/src/domains/customerDomain"
	"github.com/oommi04/shibabookbackend/src/domains/invoiceDomain"
	"github.com/oommi04/shibabookbackend/src/domains/orderDomain"
	"github.com/oommi04/shibabookbackend/src/domains/productDomain"
	_orderRepository "github.com/oommi04/shibabookbackend/src/repository/orderRepository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ProductAmount struct {
	ProductID primitive.ObjectID `bson:"productId,omitempty"`
	Amount    int                `bson:"amount,omitempty"`
}

type Order struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Status      string             `bson:"status,omitempty"`
	Description string             `bson:"description,omitempty"`
	ProductsID  []*ProductAmount   `bson:"productsId,omitempty"`
	//StaffID primitive.ObjectID `json:"staffId,omitempty"`
	CustomerID primitive.ObjectID `bson:"customerId,omitempty"`
	TotalPrice float32                `bson:"totalPrice,omitempty"`
	Discount float32	`bson:"discount,omitempty"`
	InvoiceID  primitive.ObjectID `bson:"invoiceId,omitempty"`
	DateTime   *time.Time          `bson:"time,omitempty"`
	NET float32 `bson:"net,omitempty"`
	CreatedAt  time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt  time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

type orderRepository struct {
	Collection *mongo.Collection
}

func NewProductRepository(d *mongo.Database) _orderRepository.OrderRepositoryInterface {
	collection := d.Collection("orders")
	return &orderRepository{collection}
}

func (p *orderRepository) List(ctx context.Context) ([]*orderDomain.Order, error) {
	var datas []Order

	cursor, err := p.Collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &datas); err != nil {
		return nil, err
	}

	resps := []*orderDomain.Order{}

	for _, data := range datas {

		products := []*orderDomain.ProductAmount{}
		for _,pItem := range data.ProductsID {
			products = append(products, &orderDomain.ProductAmount{
				Amount: pItem.Amount,
				Product: productDomain.Product{
					ID: pItem.ProductID.Hex(),
				},
			})
		}

		invoice := invoiceDomain.Invoice{
			ID: data.InvoiceID.Hex(),
		}

		resp := &orderDomain.Order{
			ID: data.ID.Hex(),
			TotalPrice:  data.TotalPrice,
			Description: data.Description,
			Status:      data.Status,
			DateTime:    data.DateTime,
			Products: products,
			Invoice: invoice,
			Discount: data.Discount,
			NET: data.NET,
			//StaffID: staffId,
		}
		resps = append(resps, resp)
	}

	return resps, nil
}

func (p *orderRepository) GetById(ctx context.Context,id string) (*orderDomain.Order, error) {
	var data Order

	_id, _ := primitive.ObjectIDFromHex(id)

	err := p.Collection.FindOne(ctx, bson.M{"_id":_id}).Decode(&data)

	if err == mongo.ErrNoDocuments {
		return nil, orderDomain.ErrorOrderIdNotFound
	}

	if err != nil  {
		return nil, err
	}

	products := []*orderDomain.ProductAmount{}
	for _,pItem := range data.ProductsID {
		products = append(products, &orderDomain.ProductAmount{
			Amount: pItem.Amount,
			Product: productDomain.Product{
				ID: pItem.ProductID.Hex(),
			},
		})
	}

	invoice := invoiceDomain.Invoice{
		ID: data.InvoiceID.Hex(),
	}

	customer := customerDomain.Customer{
		ID: data.CustomerID.Hex(),
	}

	resp := orderDomain.Order{
		ID: data.ID.Hex(),
		TotalPrice:  data.TotalPrice,
		Description: data.Description,
		Status:      data.Status,
		DateTime:    data.DateTime,
		Products: products,
		Customer: customer,
		Invoice: invoice,
		Discount: data.Discount,
		NET: data.NET,
		//StaffID: staffId,
	}


	return &resp, nil
}

func (p *orderRepository) Save(ctx context.Context, info *orderDomain.Order) error {
	customerId, _ := primitive.ObjectIDFromHex(info.Customer.ID)

	productsId := []*ProductAmount{}

	for _, product := range info.Products {
		pId, _ := primitive.ObjectIDFromHex(product.Product.ID)
		productsId = append(productsId, &ProductAmount{
			pId,
			product.Amount,
		})
	}

	data := Order{
		TotalPrice:  info.TotalPrice,
		Description: info.Description,
		Status:      info.Status,
		DateTime:    info.DateTime,
		ProductsID:  productsId,
		Discount: info.Discount,
		NET: info.NET,
		//StaffID: staffId,
		CustomerID: customerId,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	insertResult, err := p.Collection.InsertOne(ctx, data)

	if err != nil {
		return err
	}

	info.ID = insertResult.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

func (p *orderRepository) CheckOut(ctx context.Context, info *orderDomain.Order, id string) error {
	_id, _ := primitive.ObjectIDFromHex(id)

	customerId, _ := primitive.ObjectIDFromHex(info.Customer.ID)
	invoiceId, _ := primitive.ObjectIDFromHex(info.Invoice.ID)

	productsId := []*ProductAmount{}

	for _, product := range info.Products {
		pId, _ := primitive.ObjectIDFromHex(product.Product.ID)
		productsId = append(productsId, &ProductAmount{
			pId,
			product.Amount,
		})
	}

	data := Order{
		TotalPrice:  info.TotalPrice,
		Description: info.Description,
		Status:      info.Status,
		DateTime:    info.DateTime,
		ProductsID:  productsId,
		Discount: info.Discount,
		NET: info.NET,
		InvoiceID: invoiceId,
		//StaffID: staffId,
		CustomerID: customerId,
		UpdatedAt:  time.Now(),
	}

	_, err := p.Collection.ReplaceOne(ctx, bson.M{"_id": _id}, data)

	if err != nil {
		return err
	}

	info.ID = id

	return nil
}
