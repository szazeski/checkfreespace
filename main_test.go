package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Round(t *testing.T) {
	input := 1.2345
	expected := 1.2
	actual := roundOneDecimal(input)

	assert.Equal(t, expected, actual)
}
