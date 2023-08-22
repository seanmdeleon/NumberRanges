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

func TestBigNumberRange(t *testing.T) {

	input := []int{}
	for i := -20000; i <= 555000; i++ {
		input = append(input, i)
	}

	resultRanges := Solution(input)

	assert.Equal(t, []string{"-20000...555000"}, resultRanges)
}

func TestBigNumberRangeEndingWithSingleNumberRange(t *testing.T) {

	input := []int{}
	for i := -20000; i <= 555000; i++ {
		input = append(input, i)
	}

	input = append(input, 600000)
	resultRanges := Solution(input)

	assert.Equal(t, []string{"-20000...555000", "600000"}, resultRanges)
}

func TestSingleNumberRangesEndingInBigNumberRange(t *testing.T) {

	input := []int{}

	for i := -30000; i < -29990; i=i+2 {
		input = append(input, i)
	}

	for i := -20000; i <= 555000; i++ {
		input = append(input, i)
	}

	resultRanges := Solution(input)

	assert.Equal(t, []string{"-30000", "-29998", "-29996", "-29994", "-29992", "-20000...555000"}, resultRanges)
}