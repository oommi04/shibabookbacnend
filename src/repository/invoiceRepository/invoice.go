package invoiceRepository

import (
	"github.com/tkhamsila/shibabookbackend/src/domains/invoiceDomain"
	"context"
)

type InvoiceRepositoryInterface interface {
	Save(ctx context.Context, info *invoiceDomain.Invoice) error
}
