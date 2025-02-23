package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal8(t *testing.T) {
	result := Soal8(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[3], 8)
	assert.NotEqual(t, result[4], 10)

	result = Soal8(11)
	assert.Equal(t, len(result), 11)
	assert.Equal(t, result[9], 6)
	assert.NotEqual(t, result[10], 8)

}
