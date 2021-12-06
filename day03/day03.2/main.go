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

	fmt.Println(input)
}

func parseInput(reader io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(reader)
	result := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		linebits := []int{}
		for _, bit := range line {
			bitval, err := strconv.Atoi(string(bit))
			if err != nil {
				return nil, err
			}
			linebits = append(linebits, bitval)
		}
		result = append(result, linebits)

	}
	return result, nil

}
