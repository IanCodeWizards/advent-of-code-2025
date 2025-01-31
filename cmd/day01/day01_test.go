package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay01(t *testing.T) {
	assert := assert.New(t)
	input := "3   4\r\n4   3\r\n2   5\r\n1   3\r\n3   9\r\n3   3"

	t.Run("part 1", func(t *testing.T) {
		expected := 11
		actual := Part1(input)

		assert.Equal(actual, expected)
	})

	t.Run("part 2", func(t *testing.T) {
		expected := 31
		actual := Part2(input)

		assert.Equal(actual, expected)
	})
}
