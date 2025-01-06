package main

import (
	"aoc_24/utils"
	"fmt"
	"os"
	"strings"
	"sync"
)

// ⬅➡
func findHorizontal(puzzle [][]string) int {
	count := 0
	for _, line := range puzzle {
		for idx := range line {
			if idx+4 > len(line) {
				break
			}

			joined := strings.Join(line[idx:idx+4], "")

			if joined == "XMAS" || joined == "SAMX" {
				count++
			}
		}
	}
	return count
}

// ⬇⬆
func findVertical(puzzle [][]string) int {
	count := 0
	elements := make([]string, 4)

	for colIdx := range puzzle[0] {
		for lineIdx := range puzzle {
			if lineIdx+3 >= len(puzzle) {
				break
			}
			elements[0] = puzzle[lineIdx][colIdx]
			elements[1] = puzzle[lineIdx+1][colIdx]
			elements[2] = puzzle[lineIdx+2][colIdx]
			elements[3] = puzzle[lineIdx+3][colIdx]
			joined := strings.Join(elements, "")
			if joined == "XMAS" || joined == "SAMX" {
				count++
			}
		}
	}
	return count
}

// ↘↖
func findDiagonal1(puzzle [][]string) int {
	count := 0
	elements := make([]string, 4)

	for colIdx := range puzzle[0] {
		if colIdx+3 >= len(puzzle[0]) {
			break
		}
		for lineIdx := range puzzle {
			if lineIdx+3 >= len(puzzle) {
				break
			}
			elements[0] = puzzle[lineIdx][colIdx]
			elements[1] = puzzle[lineIdx+1][colIdx+1]
			elements[2] = puzzle[lineIdx+2][colIdx+2]
			elements[3] = puzzle[lineIdx+3][colIdx+3]
			joined := strings.Join(elements, "")
			if joined == "XMAS" || joined == "SAMX" {
				count++
			}
		}
	}
	return count
}

// ↗↙
func findDiagonal2(puzzle [][]string) int {
	count := 0
	elements := make([]string, 4)

	for colIdx := len(puzzle[0]) - 1; colIdx >= 0; colIdx-- {
		if colIdx-3 < 0 {
			break
		}
		for lineIdx := range puzzle {
			if lineIdx+3 >= len(puzzle) {
				break
			}
			elements[0] = puzzle[lineIdx][colIdx]
			elements[1] = puzzle[lineIdx+1][colIdx-1]
			elements[2] = puzzle[lineIdx+2][colIdx-2]
			elements[3] = puzzle[lineIdx+3][colIdx-3]
			joined := strings.Join(elements, "")
			if joined == "XMAS" || joined == "SAMX" {
				count++
			}
		}
	}
	return count
}

func D4P1() {
	content, _ := os.ReadFile("./inputs/d4.txt")

	puzzle := utils.Map(strings.Split(string(content), "\r\n"), func(s string) []string {
		return strings.Split(s, "")
	})

	var wg sync.WaitGroup
	c := make(chan int, 4)

	wg.Add(4)

	go func() {
		defer wg.Done()
		c <- findHorizontal(puzzle)
	}()

	go func() {
		defer wg.Done()
		c <- findVertical(puzzle)
	}()

	go func() {
		defer wg.Done()
		c <- findDiagonal1(puzzle)
	}()

	go func() {
		defer wg.Done()
		c <- findDiagonal2(puzzle)
	}()

	wg.Wait()

	close(c)

	sum := 0
	for res := range c {
		sum += res
	}

	fmt.Println(sum)

}

// --------------------------------------------- PART 2 --------------------------------------------- //

func D4P2() {
	content, _ := os.ReadFile("./inputs/d4.txt")

	puzzle := utils.Map(strings.Split(string(content), "\r\n"), func(s string) []string {
		return strings.Split(s, "")
	})

	sum := 0

	l1 := make([]string, 3)
	l2 := make([]string, 3)
	l3 := make([]string, 3)

	for lIdx := range puzzle {
		if lIdx+2 >= len(puzzle) {
			break
		}

		for cIdx := range puzzle[lIdx] {
			if cIdx+3 > len(puzzle[lIdx]) {
				break
			}

			copy(l1, puzzle[lIdx][cIdx:cIdx+3])
			l1[1] = "."

			copy(l2, puzzle[lIdx+1][cIdx:cIdx+3])
			l2[0] = "."
			l2[2] = "."

			copy(l3, puzzle[lIdx+2][cIdx:cIdx+3])
			l3[1] = "."

			// four possibilities and the middle line is aways the same
			/*
				M.S M.M S.S S.M
				.A. .A. .A. .A.
				M.S S.S M.M S.M
			*/

			joined1 := strings.Join(l1, "")
			joined2 := strings.Join(l2, "")
			joined3 := strings.Join(l3, "")
			if joined2 == ".A." {
				if joined1 == "M.S" && joined3 == "M.S" || joined1 == "M.M" && joined3 == "S.S" || joined1 == "S.S" && joined3 == "M.M" || joined1 == "S.M" && joined3 == "S.M" {
					sum++
				}
			}
		}
	}

	fmt.Println(sum)
}
