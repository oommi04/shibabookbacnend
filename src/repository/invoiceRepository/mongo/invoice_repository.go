package mongo

import (
	"github.com/oommi04/shibabookbackend/src/domains/invoiceDomain"
	_invoiceRepository "github.com/oommi04/shibabookbackend/src/repository/invoiceRepository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
)

type Invocie struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Path          string             `bson:"path,omitempty"`
	Mail         string             `bson:"mail,omitempty"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}

type invoiceRepository struct {
	Collection *mongo.Collection
}

func NewInvoiceRepository(d *mongo.Database) _invoiceRepository.InvoiceRepositoryInterface {
	collection := d.Collection("invoices")
	return &invoiceRepository{collection}
}

func (p *invoiceRepository) Save(ctx context.Context, info *invoiceDomain.Invoice) error {

	data := Invocie{
		Path: info.Path,
		Mail: info.Mail,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	insertResult, err := p.Collection.InsertOne(ctx, data)

	if err != nil {
		return err
	}

	info.ID = insertResult.InsertedID.(primitive.ObjectID).Hex()

	return nil
}