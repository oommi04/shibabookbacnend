package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtils_StringToInt(t *testing.T) {
	resp := StringToInt("5")
	expectResp := 5
	assert.Equal(t, expectResp, resp)
}

func TestUtils_MaxIntSlice(t *testing.T) {
	input := []int{1, 3, 4, 2, 0, 10, 2}
	resp := MaxIntSlice(input)
	expectResp := 10
	assert.Equal(t, expectResp, resp)
}

func TestUtils_SumFloat32Slice(t *testing.T) {
	input := []float32{1, 3, 4, 2, 0, 10, 2}
	resp := SumFloat32Slice(input)
	var expectResp float32 = 22
	assert.Equal(t, expectResp, resp)
}
