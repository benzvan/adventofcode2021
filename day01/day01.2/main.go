package main

import (
	"bufio"
	"errors"
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
	input, err = sumByThrees(input)
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

func sumByThrees(input []int) (output []int, err error) {
	if len(input) < 3 {
		return output, errors.New("input less than 3 values")
	}
	for i := 0; i < len(input)-2; i++ {
		output = append(output, (input[i] + input[i+1] + input[i+2]))
	}
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
