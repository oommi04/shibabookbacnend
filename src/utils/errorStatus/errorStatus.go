package errorStatus

import (
	"github.com/oommi04/backendtest/src/domains/doscg"
	"github.com/oommi04/backendtest/src/external/harryShop"
	"net/http"
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err {
	case doscg.ErrorUnableCreateRequest:
	case doscg.ErrorAPIKeyInvalid:
	case doscg.ErrorUnableRequestGoogleDirection:
	case harryShop.ErrorUnableCreateRequest:
	case harryShop.ErrorUnableRequestGetHarryBook:
		return http.StatusInternalServerError
	case doscg.ErrorUnableFindDirectionStartAndDestination:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
	return http.StatusInternalServerError
}
