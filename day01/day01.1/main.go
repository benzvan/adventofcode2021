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
	fmt.Println(countIncreases(input))
}

func parseInput(reader io.Reader) (input []int, err error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		input = append(input, num)
	}

	err = scanner.Err()
	return
}

func countIncreases(input []int) (total int) {
	for index := 0; index < len(input)-1; index++ {
		if input[index] < input[index+1] {
			total++
		}
	}
	return
}
