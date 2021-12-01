package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	depths, _ := readLines("src/day_01/input.txt")
	fmt.Printf("PART 1: %d \n", part1(depths))
	fmt.Printf("PART 2: %d \n", part2(depths))
}

func part1(depths []int) int {
	var prev = 0
	var inc = 0

	for _, depth := range depths {

		if prev != 0 && depth > prev {
			inc++
		}

		prev = depth
	}

	return inc
}

func part2(depths []int) int {
	var inc = 0
	var prev []int

	for _, depth := range depths {
		if len(prev) < 3 {
			prev = append(prev, depth)
			continue
		}

		var current = append(prev, depth)[1:]

		if findArraySum(current) > findArraySum(prev) {
			inc++
		}

		prev = current
	}

	return inc
}

func findArraySum(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res += arr[i]
	}
	return res
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, num)
	}
	return lines, scanner.Err()
}
