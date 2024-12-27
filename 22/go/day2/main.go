package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Result int

const (
	LOST Result = iota
	WON
	DRAW
)

func main() {
	partOne()
	partTwo()
}

func getRoundScore(you string, result Result) int {
	score := 0
	switch result {
	case DRAW:
		score += 3
	case WON:
		score += 6
	}
	switch you {
	case "X":
		score += 1
	case "Y":
		score += 2
	case "Z":
		score += 3
	}
	return score
}

func getResult(other, you string) Result {
	if other == "A" {
		if you == "Y" {
			return WON
		}
		if you == "X" {
			return DRAW
		}
	}
	if other == "B" {
		if you == "Z" {
			return WON
		}
		if you == "Y" {
			return DRAW
		}
	}
	if other == "C" {
		if you == "X" {
			return WON
		}
		if you == "Z" {
			return DRAW
		}
	}
	return LOST
}

func partOne() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	score := 0

	for scanner.Scan() {
		roundPlay := strings.Split(scanner.Text(), " ")
		score += getRoundScore(roundPlay[1], getResult(roundPlay[0], roundPlay[1]))
	}
	fmt.Println(score)
}

func getMove(other string, result Result) string {
	if other == "A" {
		if result == WON {
			return "Y"
		}
		if result == DRAW {
			return "X"
		}
		return "Z"
	}
	if other == "B" {
		if result == WON {
			return "Z"
		}
		if result == DRAW {
			return "Y"
		}
		return "X"
	}
	if other == "C" {
		if result == WON {
			return "X"
		}
		if result == DRAW {
			return "Z"
		}
		return "Y"
	}
	return ""
}

func partTwo() {
	resultMap := map[string]Result{
		"X": LOST,
		"Y": DRAW,
		"Z": WON,
	}

	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	score := 0
	for scanner.Scan() {
		roundPlay := strings.Split(scanner.Text(), " ")
		result := resultMap[roundPlay[1]]
		score += getRoundScore(getMove(roundPlay[0], result), result)
	}
	fmt.Println(score)
}
