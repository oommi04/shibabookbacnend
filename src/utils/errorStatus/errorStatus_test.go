package errorStatus

import (
	"github.com/oommi04/backendtest/src/domains/doscg"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestUtils_ErrorStatus_500(t *testing.T) {
	err := doscg.ErrorAPIKeyInvalid
	resp := GetStatusCode(err)
	expectResp := http.StatusInternalServerError
	assert.Equal(t, expectResp, resp)
}

func TestUtils_ErrorStatus_200(t *testing.T) {
	resp := GetStatusCode(nil)
	expectResp := http.StatusOK
	assert.Equal(t, expectResp, resp)
}

func TestUtils_ErrorStatus_404(t *testing.T) {
	err := doscg.ErrorUnableFindDirectionStartAndDestination
	resp := GetStatusCode(err)
	expectResp := http.StatusNotFound
	assert.Equal(t, expectResp, resp)
}
