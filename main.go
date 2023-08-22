package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{5,2,5,6,3,2}
	res := Solution(input)
	fmt.Println(res)
}

/*
	Given an input of unsorted integers, return all number ranges of the input without duplicates.
	A range is represented as either "n...m" where n-1 and m+1 are not included in the input, or just "n" where n-1 or n+1 are not included in the input

	For my own design, this input will only accept integers that fit within a 32-bit int (or 64 depending on the platform), both postitive and negative

	Big-O evaluations:
	O(NlogN) for time, due to sorting
	O(N) for space, for the return value, but if we exlude the return value in the evaluation of space, the complexity is constant O(1)
*/
func Solution(input []int) []string {

	// sort the input array using the sort package - O(nlogn), where n is the length of the input array
	sort.Ints(input)

	// O(n) space for the result, where n is the length of the array. O(n) because worst case if each input in the array contains "n" but not "n+1" or "n-1"
	result := []string{}

	// Use sliding window technique to build number ranges
	left, right := 0, 0
	s := ""

	// O(n) for this loop, where n is the length of input
	for right < len(input) {

		if left == right {

			// build a new range
			s = fmt.Sprintf("%d", input[right])
			right++

		} else if input[right] > input[right-1] + 1 {
			// complete a range
			
			if input[right-1] != input[left] {
				// only create an n...m range if n != m
				s = fmt.Sprintf("%s...%d", s, input[right-1])
			}

			result = append(result, s)
			
			// reset the window
			s = ""
			left = right

		} else {
			right++
		}
	}

	// Final range processing
	if len(s) > 0 {

		if input[right-1] != input[left] {
			// only create an n...m range if n != m
			s = fmt.Sprintf("%s...%d", s, input[right-1])
		}
		result = append(result, s)
	}

	return result
}
