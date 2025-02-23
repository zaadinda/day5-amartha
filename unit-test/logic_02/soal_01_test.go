package logic_02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal1(t *testing.T) {
	result := Soal1(10)

	assert.Equal(t, len(result), 10)
	assert.Equal(t, result[0], 10)
	// BELUM DILANJUTIN

}
