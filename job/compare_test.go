package job

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TesCompare(t *testing.T) {
	a := []byte("a very long string a")
	b := []byte("a very long string b")
	c := []byte("short")

	assert.False(t, CompareBytes(a, b))
	assert.False(t, CompareBytes(a, c))
	assert.True(t, CompareBytes(c, c))
	assert.True(t, CompareBytes(a, a))
	assert.True(t, CompareBytes([]byte(""), []byte{}))
}
