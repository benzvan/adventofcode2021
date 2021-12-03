package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"

	"fmt"
	"log"
	"os"
)

type submarine struct {
	depth    int
	distance int
	aim      int
}

type move struct {
	direction string
	distance  int
}

func main() {
	input, err := parseInput(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	sub := submarine{}
	for _, move := range input {
		sub.move(move.direction, move.distance)
	}
	fmt.Println(sub.depth * sub.distance)
}

func parseInput(reader io.Reader) (input []move, err error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		distance, err := strconv.Atoi(line[1])
		if err != nil {
			return nil, err
		}
		input = append(input, move{
			direction: line[0],
			distance:  distance,
		})

	}

	err = scanner.Err()
	return
}

func (s *submarine) move(direction string, distance int) {
	switch direction {
	case "forward":
		s.distance = s.distance + distance
		s.depth = s.depth + (s.aim * distance)
	case "up":
		s.aim = s.aim - distance
	case "down":
		s.aim = s.aim + distance
	}
}
