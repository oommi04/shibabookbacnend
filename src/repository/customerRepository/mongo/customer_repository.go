package mongo

import (
	"context"
	"github.com/oommi04/shibabookbackend/src/domains/customerDomain"
	_customerRepository "github.com/oommi04/shibabookbackend/src/repository/customerRepository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Customer struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	Name                 string             `bson:"name,omitempty"`
	LastName             string             `bson:"lastName,omitempty"`
	IdentificationNumber string             `bson:"identificationNumber,omitempty"`
	Point                int                `bson:"point,omitempty"`
	Mail                 string             `bson:"mail,omitempty"`
	CreatedAt            time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt            time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

type customerRepository struct {
	Collection *mongo.Collection
}

func NewCustomerRepository(d *mongo.Database) _customerRepository.CustomerRepositoryInterface {
	collection := d.Collection("customers")
	return &customerRepository{collection}
}

func (p *customerRepository) GetByIdentificationNumber(ctx context.Context, idn string) (*customerDomain.Customer, error) {
	var data Customer

	err := p.Collection.FindOne(ctx, bson.D{{"identificationNumber", idn}}).Decode(&data)

	if err == mongo.ErrNoDocuments {
		return nil, customerDomain.ErrorIdentificationNotFound
	}

	if err != nil {
		return nil, err
	}

	resp := customerDomain.Customer{
		ID:                   data.ID.Hex(),
		Name:                 data.Name,
		LastName:             data.LastName,
		IdentificationNumber: data.IdentificationNumber,
		Point:                data.Point,
		Mail:                 data.Mail,
	}

	return &resp, nil
}

func (p *customerRepository) GetById(ctx context.Context, id string) (*customerDomain.Customer, error) {
	var data Customer

	_id, _ := primitive.ObjectIDFromHex(id)

	err := p.Collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&data)

	if err == mongo.ErrNoDocuments {
		return nil, customerDomain.ErrorCustomerIdNotFound
	}

	if err != nil {
		return nil, err
	}

	resp := customerDomain.Customer{
		ID:                   data.ID.Hex(),
		Name:                 data.Name,
		LastName:             data.LastName,
		IdentificationNumber: data.IdentificationNumber,
		Point:                data.Point,
		Mail:                 data.Mail,
	}

	return &resp, nil
}

func (p *customerRepository) Save(ctx context.Context, info *customerDomain.Customer) error {

	data := Customer{
		Name:                 info.Name,
		LastName:             info.LastName,
		IdentificationNumber: info.IdentificationNumber,
		Point:                info.Point,
		Mail:                 info.Mail,
		CreatedAt:            time.Now(),
		UpdatedAt:            time.Now(),
	}

	_, err := p.Collection.InsertOne(ctx, data)

	if err != nil {
		return err
	}

	return nil
}
