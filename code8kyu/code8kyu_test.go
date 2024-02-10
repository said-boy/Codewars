package code8kyu

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvertValues(t *testing.T) {
	assert.Equal(t, []int{-1,-2,-3,-4,-5}, Invert([]int{1,2,3,4,5}))
	assert.Equal(t, []int{-1,2,-3,4,-5}, Invert([]int{1,-2,3,-4,5}))
	assert.Equal(t, []int(nil), Invert(nil))
	assert.Equal(t, []int{0}, Invert([]int{0}))
}

func TestStringToNumber(t *testing.T) {
	assert.Equal(t, 1234, StringToNumber("1234"))
	assert.Equal(t, 605, StringToNumber("605"))
	assert.Equal(t, 1405, StringToNumber("1405"))
	assert.Equal(t, -7, StringToNumber("-7"))
}

func TestNumberToString(t *testing.T) {
	assert.Equal(t, "67", NumberToString(67))
	assert.Equal(t, "79585", NumberToString(79585))
	assert.Equal(t, "79585", NumberToString(79585))
	assert.Equal(t, "3", NumberToString(3))
	assert.Equal(t, "-1", NumberToString(-1))
}
