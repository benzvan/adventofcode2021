package main

import (
	"sort"
)

func readDiagnostic(input []int) []int {
	sort.Ints(input)
	mostCommon := input[len(input)-1]
	leastCommon := ^mostCommon
	return []int{
		mostCommon,
		leastCommon,
	}
}

func getCommons(input []uint8) []uint8 {
	sort.Slice(input, func(i, j int) bool { return i < j })
	mostCommon := input[len(input)-1]
	leastCommon := 1 ^ mostCommon
	return []uint8{
		mostCommon,
		leastCommon,
	}
}
