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