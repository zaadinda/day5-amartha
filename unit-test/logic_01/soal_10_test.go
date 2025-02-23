package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal10(t *testing.T) {
	result := Soal10(10)
	assert.Equal(t, result[2], 4)
}
