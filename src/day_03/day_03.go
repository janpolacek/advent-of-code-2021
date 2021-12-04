package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	measurements, _ := read("src/day_03/input.txt")

	fmt.Printf("PART 1: %d \n", part1(measurements))
	fmt.Printf("PART 2: %d \n", part2(measurements))
}

func part1(measurements [][]int) int64 {

	var sum = make([]int, len(measurements[0]))

	for _, bites := range measurements {
		for bIndex, bValue := range bites {
			sum[bIndex] = sum[bIndex] + bValue
		}
	}

	var gamaBiteMeasurement = make([]int, len(sum))
	var epsBiteMeasurement = make([]int, len(sum))

	for sIndex, sValue := range sum {
		if sValue > len(measurements)/2 {
			gamaBiteMeasurement[sIndex] = 1
		}

		if sValue < len(measurements)/2 {
			epsBiteMeasurement[sIndex] = 1
		}
	}

	var gama = binArrayToInt(gamaBiteMeasurement)
	var eps = binArrayToInt(epsBiteMeasurement)

	return gama * eps
}

func binArrayToInt(binArray []int) int64 {
	var _string = ""
	for _, b := range binArray {
		_string = strings.Join([]string{_string, strconv.Itoa(b)}, "")
	}

	parseInt, _ := strconv.ParseInt(_string, 2, 64)

	return parseInt
}

func part2(measurements [][]int) int64 {
	ox := measurements
	co2 := measurements

	for i := 0; i < len(measurements[0]); i++ {
		ox = filterOx(ox, i)
	}
	var oxValue = binArrayToInt(ox[0])

	for i := 0; i < len(measurements[0]); i++ {
		co2 = filterCo2(co2, i)
	}

	var co2Value = binArrayToInt(co2[0])

	return oxValue * co2Value
}

func filterOx(measurements [][]int, index int) [][]int {

	if len(measurements) == 1 {
		return measurements
	}

	var bitSum = 0

	for _, measurement := range measurements {
		bitSum += measurement[index]
	}

	var mostPopular = 0
	if bitSum*2 >= len(measurements) {
		mostPopular = 1
	}

	var filtered = make([][]int, 0, len(measurements))
	for _, measurement := range measurements {
		if measurement[index] == mostPopular {
			filtered = append(filtered, measurement)
		}
	}

	return filtered
}

func filterCo2(measurements [][]int, index int) [][]int {

	if len(measurements) == 1 {
		return measurements
	}

	var bitSum = 0

	for _, measurement := range measurements {
		bitSum += measurement[index]
	}

	var leastPopular = 0
	if bitSum*2 < len(measurements) {
		leastPopular = 1
	}

	var filtered = make([][]int, 0, len(measurements))
	for _, measurement := range measurements {
		if measurement[index] == leastPopular {
			filtered = append(filtered, measurement)
		}
	}

	return filtered
}

// read reads a whole file into memory
// and returns a slice of its lines.
func read(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var data [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var bites []int

		for _, i := range strings.Split(line, "") {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			bites = append(bites, j)
		}

		data = append(data, bites)
	}
	return data, scanner.Err()
}
