package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isSafe(report []int) bool {

	if len(report) <= 1 {
		return true
	}

	direction := 0 // -1 == descending ||| 1 == ascending
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]

		if diff == 0 {
			return false
		}

		if direction == 0 {
			if diff > 0 {
				direction = 1
			} else {
				direction = -1
			}
		}

		if direction == 1 && diff < 0 {
			return false
		} else if direction == -1 && diff > 0 {
			return false
		}

		diff = int(math.Abs(float64(diff)))
		if diff > 3 {
			return false
		}
	}

	return true
}

func deleteElement(s []int, TargetIdx int) []int {
	res := make([]int, 0, len(s)-1)
	for i, v := range s {
		if i == TargetIdx {
			continue
		}
		res = append(res, v)
	}
	return res
}

func D2P1() {
	file, _ := os.Open("./inputs/d2.txt")
	defer file.Close()

	scn := bufio.NewScanner(file)

	sum := 0
	for scn.Scan() {
		report := make([]int, 0, 64)
		for _, lvl_str := range strings.Split(scn.Text(), " ") {
			i, _ := strconv.Atoi(lvl_str)
			report = append(report, i)
		}
		if isSafe(report) {
			sum++
		}
	}
	fmt.Println(sum)
}

func D2P2() {
	file, _ := os.Open("./inputs/d2.txt")
	defer file.Close()

	scn := bufio.NewScanner(file)

	sum := 0
	for scn.Scan() {
		report := make([]int, 0, 64)
		for _, lvl_str := range strings.Split(scn.Text(), " ") {
			i, _ := strconv.Atoi(lvl_str)
			report = append(report, i)
		}

		if isSafe(report) {
			sum++
			continue
		}

		for i := range report {
			new_report := deleteElement(report, i)
			if isSafe(new_report) {
				sum++
				break
			}
		}
	}
	fmt.Println(sum)
}
