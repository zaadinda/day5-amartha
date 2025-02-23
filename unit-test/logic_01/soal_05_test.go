package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal5(t *testing.T) {
	result := Soal5(10)
	assert.Equal(t, result[2], 16)
	assert.NotEqual(t, result[3], 17)
}
