package errorStatus

import (
	"github.com/stretchr/testify/assert"
	"github.com/tkhamsila/backendtest/src/domains/doscg"
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
