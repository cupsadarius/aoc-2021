package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	file "aoc/internal/utils/file"
)

func TestFileOpen(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		lines := file.ReadFile("../inputs/input.txt")
		assert.True(t, len(lines) == 10)
	})
}



