package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal9(t *testing.T) {
	result := Soal9(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[9], 6)
	assert.NotEqual(t, result[10], 8)

	result = Soal8(11)
	assert.Equal(t, len(result), 11)
	assert.Equal(t, result[9], 9)
	assert.NotEqual(t, result[10], 8)

}
