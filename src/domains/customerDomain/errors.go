package customerDomain

import "errors"

var (
	ErrorIdentificationNumberWasUsed                 = errors.New("unable create customer. cause identification number was used")
	ErrorIdentificationNotFound              = errors.New("unable get infomation customer. cause identification not found")
	ErrorCustomerIdNotFound                   = errors.New("unable get infomation customer. cause _id not found")
)
