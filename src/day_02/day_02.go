package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	movements, _ := readMovements("src/day_02/input.txt")

	fmt.Printf("PART 1: %d \n", part1(movements))
	fmt.Printf("PART 2: %d \n", part2(movements))
}

func part1(movements []Direction) int {
	var x = 0
	var y = 0

	for _, move := range movements {
		if move.direction == "up" {
			y -= move.step
		}
		if move.direction == "down" {
			y += move.step
		}

		if move.direction == "forward" {
			x += move.step
		}

	}
	return x * y
}
func part2(movements []Direction) int {
	var x = 0
	var y = 0
	var aim = 0

	for _, move := range movements {
		if move.direction == "up" {
			aim -= move.step
		}

		if move.direction == "down" {
			aim += move.step
		}

		if move.direction == "forward" {
			x += move.step
			y += aim * move.step
		}

	}
	return x * y
}

// readMovements reads a whole file into memory
// and returns a slice of its lines.
func readMovements(path string) ([]Direction, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data []Direction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, string(' '))
		var direction = split[0]
		var step, _ = strconv.Atoi(split[1])
		data = append(data, Direction{direction: direction, step: step})
	}
	return data, scanner.Err()
}

type Direction struct {
	direction string
	step      int
}
