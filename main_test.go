package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyInput(t *testing.T) {

	input := []int{}

	resultRanges := Solution(input)

	assert.Equal(t, []string{}, resultRanges)
}

func TestFullyDuplicateInput(t *testing.T) {

	input := []int{2, 2, 2, 2, 2, 2, 2}

	resultRanges := Solution(input)

	assert.Equal(t, []string{"2"}, resultRanges)
}

func TestSingleNumberRangesOnly(t *testing.T) {

	input := []int{-1, -14, -2001, 5, 9, 134, 356789}

	resultRanges := Solution(input)

	assert.Equal(t, []string{"-2001", "-14", "-1", "5", "9", "134", "356789"}, resultRanges)
}

func TestMultipleNumberRangesOnly(t *testing.T) {
	input := []int{200, 3, 6, 2001, 2002, 198, 5, 4, 199, -1, -2, -3}

	resultRanges := Solution(input)

	assert.Equal(t, []string{"-3...-1", "3...6", "198...200", "2001...2002"}, resultRanges)
}

func TestBothSingleAndMultipleNumberRanges(t *testing.T) {
	input := []int{200, 3, 6, 2001, 2002, 198, 5, 4, 199, -1, -2, -3, 45, 67, 89, -20000, 256789}

	resultRanges := Solution(input)

	assert.Equal(t, []string{"-20000", "-3...-1", "3...6", "45", "67", "89", "198...200", "2001...2002", "256789"}, resultRanges)
}
