package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal7(t *testing.T) {
	result := Soal7(10)

	// make sure panjang slicenya itu 10
	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[3], 7)
	assert.NotEqual(t, result[4], 7)

	result = Soal7(11)
	assert.Equal(t, len(result), 11)
	assert.Equal(t, result[2], 5)
	assert.NotEqual(t, result[2], 10)

}
