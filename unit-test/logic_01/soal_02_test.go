package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal2(t *testing.T) {
	result := Soal2(10)

	assert.Equal(t, len(result), 9)
	assert.Equal(t, result[3], 7)
	assert.NotEqual(t, result[3], 9)

	result = Soal2(13)
	assert.NotEqual(t, result[3], 9)

}
