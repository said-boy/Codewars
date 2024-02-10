package code6kyu

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTakeATenMinutesWalk(t *testing.T) {
	assert.True(t, IsValidWalk([]rune{'n','s','n','s','n','s','n','s','n','s'}))
	assert.False(t, IsValidWalk([]rune{'w','e','w','e','w','e','w','e','w','e','w','e'}))
	assert.False(t, IsValidWalk([]rune{'w'}))
	assert.False(t, IsValidWalk([]rune{'n','n','n','s','n','s','n','s','n','s'}))
	assert.False(t, IsValidWalk([]rune{'e', 'e', 'e', 'e', 'w', 'w', 's', 's', 's', 's'}))
}