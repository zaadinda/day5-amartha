package logic_01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal3(t *testing.T) {
	result := Soal3(9)

	// check the length of result
	assert.Equal(t, len(result), 9)

	// check specific values
	assert.Equal(t, 3, result[0])    // indeks pertama harusnya 3
	assert.Equal(t, 15, result[4])   // indeks kelimaharusnya 15
	assert.NotEqual(t, 9, result[3]) // indeks keempat harusnya 12

	// ensure function consistency
	result2 := Soal3(9)
	assert.Equal(t, result, result2)
}
