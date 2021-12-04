package days

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const day2Input string = "inputs/02.txt"

func Day2Part1() int {
	file, _ := os.Open(day2Input)
	defer file.Close()

	depth := 0
	position := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		direction := parts[0]
		amount, _ := strconv.Atoi(parts[1])
		switch direction {
		case "forward":
			position += amount
		case "up":
			depth -= amount
		case "down":
			depth += amount
		}
	}

	return depth * position
}

func Day2Part2() int {
	file, _ := os.Open(day2Input)
	defer file.Close()

	depth := 0
	position := 0
	aim := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		direction := parts[0]
		amount, _ := strconv.Atoi(parts[1])
		switch direction {
		case "forward":
			position += amount
			depth += aim * amount
		case "up":
			aim -= amount
		case "down":
			aim += amount
		}
	}

	return depth * position
}
