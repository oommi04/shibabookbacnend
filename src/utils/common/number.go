package common

import (
	"sort"
	"strconv"
)

func StringToInt(str string) int {
	i1, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return i1
}

func MaxIntSlice(v []int) int {
	sort.Ints(v)
	return v[len(v)-1]
}

func SumFloat32Slice(v []float32) float32 {
	var sum float32 = 0
	for _, item := range v {
		sum = sum + item
	}
	return sum
}
