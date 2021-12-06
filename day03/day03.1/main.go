package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(getGamma(input))
}

func parseInput(reader io.Reader) ([]int64, error) {
	scanner := bufio.NewScanner(reader)
	var result []int64

	for scanner.Scan() {
		line := scanner.Text()
		if len(result) == 0 {
			result = make([]int64, len(line))
		}

		for pos, bit := range line {
			bitval, err := strconv.Atoi(string(bit))
			if err != nil {
				return nil, err
			}

			switch bitval {
			case 1:
				result[pos]++
			case 0:
				result[pos]--
			}

		}

	}
	return result, nil

}

func getGamma(bits []int64) uint64 {
	var gamma string
	var epsilon string
	for _, bit := range bits {
		if bit > 0 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaval, _ := strconv.ParseUint(gamma, 2, 64)
	epsival, _ := strconv.ParseUint(epsilon, 2, 64)

	return gammaval * epsival
}
