package errorStatus

import (
	"github.com/oommi04/shibabookbackend/src/domains/customerDomain"
	"github.com/oommi04/shibabookbackend/src/domains/orderDomain"
	"github.com/oommi04/shibabookbackend/src/external/harryShop"
	"net/http"
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case harryShop.ErrorUnableCreateRequest:
	case harryShop.ErrorUnableRequestGetHarryBook:
		return http.StatusInternalServerError
	case customerDomain.ErrorIdentificationNotFound:
	case customerDomain.ErrorIdentificationNotFound:
	case orderDomain.ErrorOrderIdNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
