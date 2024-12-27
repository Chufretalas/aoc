package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"golang.org/x/exp/slices"
)

func main() {
	partOne()
}

func partOne() {
	start := time.Now()
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	mostCalories := 0
	elfCalories := 0

	for scanner.Scan() {
		if linetext := scanner.Text(); linetext != "" {
			if v, err := strconv.Atoi(linetext); err == nil {
				elfCalories += v
			}
		} else {
			if elfCalories > mostCalories {
				mostCalories = elfCalories
			}
			elfCalories = 0
		}
	}
	if elfCalories > mostCalories {
		mostCalories = elfCalories
	}
	fmt.Println(mostCalories)
	fmt.Printf("Exc. time: %v us", time.Now().UnixMicro()-start.UnixMicro())
}

func partTwo() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	topThree := make([]int, 0, 3)
	elfCalories := 0

	for scanner.Scan() {
		if linetext := scanner.Text(); linetext != "" {
			if v, err := strconv.Atoi(linetext); err == nil {
				elfCalories += v
			}
		} else {
			if len(topThree) < 3 {
				topThree = append(topThree, elfCalories)
				slices.Sort(topThree)
			} else {
				idx := slices.IndexFunc(topThree, func(v int) bool {
					if elfCalories > v {
						return true
					}
					return false
				})

				if idx != -1 {
					topThree[idx] = elfCalories
					slices.Sort(topThree)
				}
			}

			elfCalories = 0
		}
	}
	idx := slices.IndexFunc(topThree, func(v int) bool {
		if elfCalories > v {
			return true
		}
		return false
	})

	if idx != -1 {
		topThree[idx] = elfCalories
		slices.Sort(topThree)
	}

	sum := 0

	for _, v := range topThree {
		sum += v
	}
	fmt.Println(sum)
}
