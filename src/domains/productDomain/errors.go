package productDomain

import "errors"

var (
	ErrorProductExist                   = errors.New("unable create product. cause product still exist")
	ErrorProductNotFoundByNameAndFrom                   = errors.New("unable get information product. cause name and from not found")
)
