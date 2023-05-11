package util

import (
	"strconv"
)

func StringToFloat64(input string) float64 {
	output, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic(err)
	}
	return output
}

func StringToPointerInt64(input *string) *int64 {
	output, err := strconv.ParseInt(*input, 10, 64)
	if err != nil {
		panic(err)
	}
	return &output
}
