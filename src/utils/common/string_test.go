package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtils_IncludeString_Include(t *testing.T) {
	strs := []string{"a","b","c","d"}
	s := "b"
	resp := IncludeString(strs, s)
	expectResp := true
	assert.Equal(t, expectResp, resp)
}

func TestUtils_IncludeString_Not_Include(t *testing.T) {
	strs := []string{"a","b","c","d"}
	s := "e"
	resp := IncludeString(strs, s)
	expectResp := false
	assert.Equal(t, expectResp, resp)
}
