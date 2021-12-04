package days

import (
	"bufio"
	"os"
	"strings"
)

const day3Input string = "inputs/03.txt"

func Day3Part1() uint {
	file, _ := os.Open(day3Input)
	defer file.Close()

	zeroCount := make([]int, 12)
	oneCount := make([]int, 12)
	for i := 0; i < 12; i++ {
		zeroCount[i] = 0
		oneCount[i] = 0
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		chars := strings.Split(line, "")
		for i, char := range chars {
			if char == "0" {
				zeroCount[i] += 1
			} else if char == "1" {
				oneCount[i] += 1
			}
		}
	}

	var number uint
	if oneCount[0] > zeroCount[0] {
		number = 1
	} else {
		number = 0
	}
	for i := 1; i < 12; i++ {
		number = number << 1
		if oneCount[i] > zeroCount[i] {
			number += 1
		}
	}

	return number * (0b111111111111 ^ number)
}

func Day3Part2() uint {
	file, _ := os.Open(day3Input)
	defer file.Close()

	lines := make([]uint, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, asciiBinaryToNumber(line))
	}

	o2Candidates := make([]uint, len(lines))
	copy(o2Candidates, lines)
	for i := 11; i >= 0 && len(o2Candidates) > 1; i-- {
		zeroCount := 0
		oneCount := 0
		for _, number := range o2Candidates {
			if (1<<i)&number == 0 {
				zeroCount++
			} else {
				oneCount++
			}
		}
		var target uint
		if oneCount >= zeroCount {
			target = 1
		} else {
			target = 0
		}
		n := 0
		for _, number := range o2Candidates {
			if (1<<i)&number == target<<i {
				o2Candidates[n] = number
				n++
			}
		}
		o2Candidates = o2Candidates[:n]
	}

	co2Candidates := lines
	for i := 11; i >= 0 && len(co2Candidates) > 1; i-- {
		zeroCount := 0
		oneCount := 0
		for _, number := range co2Candidates {
			if (1<<i)&number == 0 {
				zeroCount++
			} else {
				oneCount++
			}
		}
		var target uint
		if oneCount >= zeroCount {
			target = 0
		} else {
			target = 1
		}
		n := 0
		for _, number := range co2Candidates {
			if (1<<i)&number == target<<i {
				co2Candidates[n] = number
				n++
			}
		}
		co2Candidates = co2Candidates[:n]
	}

	return o2Candidates[0] * co2Candidates[0]
}

func asciiBinaryToNumber(s string) uint {
	var number uint = 0
	for _, val := range strings.Split(s, "") {
		number <<= 1
		if val == "1" {
			number += 1
		}
	}
	return number
}
