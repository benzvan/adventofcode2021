package main

import (
	"bufio"
	"io"
	"sort"
	"strconv"
)

func parseInput(reader io.Reader) (input [][]uint, err error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		for i := range []rune(line) {
			digit, err := strconv.ParseUint(string(line[i]), 2, 1)
			if err != nil {
				return nil, err
			}
			if input == nil {
				input = make([][]uint, len(line))
			}
			input[i] = append(input[i], uint(digit))
		}

	}

	err = scanner.Err()
	return
}

func getCommons(input []uint) []uint {
	sort.Slice(input, func(i, j int) bool { return i < j })
	mostCommon := input[len(input)-1]
	leastCommon := 1 ^ mostCommon
	return []uint{
		mostCommon,
		leastCommon,
	}
}
