package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal4(t *testing.T) {
	result := Soal4(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[4], 12)
	assert.NotEqual(t, result[3], 9)

	result2 := Soal4(10)
	assert.Equal(t, result, result2)

}
