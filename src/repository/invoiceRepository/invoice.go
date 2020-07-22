package invoiceRepository

import (
	"github.com/oommi04/shibabookbackend/src/domains/invoiceDomain"
	"context"
)

type InvoiceRepositoryInterface interface {
	Save(ctx context.Context, info *invoiceDomain.Invoice) error
}
