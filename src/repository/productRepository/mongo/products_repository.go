package mongo

import (
	"context"
	"github.com/oommi04/shibabookbackend/src/domains/productDomain"
	_productRepository "github.com/oommi04/shibabookbackend/src/repository/productRepository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Product struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name,omitempty"`
	Price         float32            `bson:"price,omitempty"`
	QuantityStock string             `bson:"quantityStock,omitempty"`
	Description   string             `bson:"description,omitempty"`
	From          string             `bson:"from,omitempty"`
	Status        string             `bson:"status,omitempty"`
	DateTime      *time.Time         `bson:"dateTime,omitempty"`
	Image         string             `bson:"image,omitempty"`
	BarCode       string             `bson:"barCode,omitempty"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

type productRepository struct {
	Collection *mongo.Collection
}

func New(d *mongo.Database) _productRepository.ProductRepositoryInterface {
	collection := d.Collection("products")
	return &productRepository{collection}
}

func (p *productRepository) List(ctx context.Context) ([]*productDomain.Product, error) {
	var datas []Product

	cursor, err := p.Collection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &datas); err != nil {
		return nil, err
	}

	resps := []*productDomain.Product{}

	for _, data := range datas {
		resp := &productDomain.Product{
			ID:            data.ID.Hex(),
			Name:          data.Name,
			Price:         data.Price,
			QuantityStock: data.QuantityStock,
			Description:   data.Description,
			From:          data.From,
			Status:        data.Status,
			DateTime:      data.DateTime,
			Image:         data.Image,
		}
		resps = append(resps, resp)
	}

	return resps, nil
}

func (p *productRepository) ListByNameAndFrom(ctx context.Context, name string, from string) ([]*productDomain.Product, error) {
	var datas []Product

	cursor, err := p.Collection.Find(ctx, bson.D{{"name", bson.D{{"$regex", name}}}, {"from", from}})

	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &datas); err != nil {
		return nil, err
	}

	resps := []*productDomain.Product{}

	for _, data := range datas {
		resp := &productDomain.Product{
			ID:            data.ID.Hex(),
			Name:          data.Name,
			Price:         data.Price,
			QuantityStock: data.QuantityStock,
			Description:   data.Description,
			From:          data.From,
			Status:        data.Status,
			DateTime:      data.DateTime,
			Image:         data.Image,
			BarCode:       data.BarCode,
		}
		resps = append(resps, resp)
	}

	return resps, nil
}

func (p *productRepository) GetByNameAndFrom(ctx context.Context, name string, from string) (*productDomain.Product, error) {
	var data Product

	err := p.Collection.FindOne(ctx, bson.D{{"name", name}, {"from", from}}).Decode(&data)

	if err == mongo.ErrNoDocuments {
		return nil, productDomain.ErrorProductNotFoundByNameAndFrom
	}

	if err != nil {
		return nil, err
	}

	resp := productDomain.Product{
		data.ID.Hex(),
		data.Name,
		data.Price,
		data.QuantityStock,
		data.Description,
		data.From,
		data.Status,
		data.DateTime,
		data.Image,
		data.BarCode,
	}

	return &resp, nil
}

func (p *productRepository) Save(ctx context.Context, info *productDomain.Product) error {

	data := Product{
		Name:          info.Name,
		Price:         info.Price,
		QuantityStock: info.QuantityStock,
		Description:   info.Description,
		From:          info.From,
		Status:        info.Status,
		DateTime:      info.DateTime,
		Image:         info.Image,
		BarCode:       info.BarCode,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	_, err := p.Collection.InsertOne(ctx, data)

	if err != nil {
		return err
	}

	return nil
}
