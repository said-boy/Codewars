package code8kyu

// Link Codewars : https://www.codewars.com/kata/5899dc03bc95b1bf1b0000ad/train/go

import (
	"math"
)

func Invert(arr []int) []int {
	result := []int{}
	if arr == nil {
		return arr
	}

	for _, v := range arr {
		if v == 0 {
			result = append(result, v)
		}

		if v < 0 {
			result = append(result, int(math.Abs(float64(v))))
		}

		if v > 0 {
			positive := int(math.Abs(float64(v)))
			result = append(result, positive - (positive*2))
		}

	}

	return result
}
