package invoiceRepository

import (
	"context"
	"github.com/oommi04/shibabookbackend/src/domains/invoiceDomain"
)

type InvoiceRepositoryInterface interface {
	Save(ctx context.Context, info *invoiceDomain.Invoice) error
}
