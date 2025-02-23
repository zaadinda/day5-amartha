package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal6(t *testing.T) {
	result := Soal6(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[3], 7)
	assert.NotEqual(t, result[3], 10)

	result = Soal6(15)
	assert.NotEqual(t, result[3], 10)

}
