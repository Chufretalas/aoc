package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func splitPairs(line string) ([]int, []int) {
	pairs := strings.Split(line, ",")
	elf1 := make([]int, 0, 2)
	elf2 := make([]int, 0, 2)

	for _, numStr := range strings.Split(pairs[0], "-") {
		num, _ := strconv.Atoi(numStr)
		elf1 = append(elf1, num)
	}

	for _, numStr := range strings.Split(pairs[1], "-") {
		num, _ := strconv.Atoi(numStr)
		elf2 = append(elf2, num)
	}

	return elf1, elf2
}

func main() {
	PartOne()
	PartTwo()
}

func PartOne() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		elf1, elf2 := splitPairs(line)
		if elf1[0] >= elf2[0] && elf1[1] <= elf2[1] || elf2[0] >= elf1[0] && elf2[1] <= elf1[1] {
			count += 1
		}
	}

	fmt.Println(count)
}

func PartTwo() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		elf1, elf2 := splitPairs(line)
		if elf1[1] >= elf2[0] && elf1[1] <= elf2[1] || elf1[0] >= elf2[0] && elf1[0] <= elf2[1] || elf2[1] >= elf1[0] && elf2[1] <= elf1[1] || elf2[0] >= elf1[0] && elf2[0] <= elf1[1] {
			count += 1
		}
	}

	fmt.Println(count)
}
