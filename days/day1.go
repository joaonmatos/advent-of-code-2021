package days

import (
	"bufio"
	"os"
	"strconv"
)

const day1Input string = "inputs/011.txt"

func Day1Part1() int {
	file, _ := os.Open(day1Input)
	defer file.Close()

	lastValue := -1
	count := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		if lastValue > 0 && value > lastValue {
			count++
		}
		lastValue = value
	}

	return count
}

func Day1Part2() int {
	file, _ := os.Open(day1Input)
	defer file.Close()

	val1, val2, val3 := -1, -1, -1
	count := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		if val1 > 0 && val2 > 0 && val3 > 0 && val2+val3+value > val1+val2+val3 {
			count++
		}
		val1 = val2
		val2 = val3
		val3 = value
	}

	return count
}
