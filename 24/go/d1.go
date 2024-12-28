package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func D1P1() {
	file, _ := os.Open("./inputs/d1.txt")
	defer file.Close()

	scn := bufio.NewScanner(file)

	l1 := make([]int, 0, 256)
	l2 := make([]int, 0, 256)
	for scn.Scan() {
		split := strings.Split(scn.Text(), "   ")
		i1, _ := strconv.Atoi(split[0])
		i2, _ := strconv.Atoi(split[1])
		l1 = append(l1, i1)
		l2 = append(l2, i2)
	}

	slices.Sort(l1)
	slices.Sort(l2)

	sum := 0
	for i, _ := range l1 {
		diff := l1[i] - l2[i]
		sum += int(math.Abs(float64(diff)))
	}
	fmt.Println(sum)
}

func D1P2() {
	file, _ := os.Open("./inputs/d1.txt")
	defer file.Close()

	scn := bufio.NewScanner(file)

	l1 := make([]int, 0, 256)
	l2 := make([]int, 0, 256)
	for scn.Scan() {
		split := strings.Split(scn.Text(), "   ")
		i1, _ := strconv.Atoi(split[0])
		i2, _ := strconv.Atoi(split[1])
		l1 = append(l1, i1)
		l2 = append(l2, i2)
	}

	sum := 0
	freqs := make(map[int]int)
	for _, i1 := range l1 {
		freq, ok := freqs[i1]
		if !ok {
			freq = 0
			for _, i2 := range l2 {
				if i1 == i2 {
					freq++
				}
			}
			freqs[i1] = freq
		}
		sum += i1 * freq
	}
	fmt.Println(sum)
}
