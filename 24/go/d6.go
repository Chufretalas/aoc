package main

import (
	u "aoc_24/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func D6P1() {

	file, _ := os.Open("./inputs/d6.txt")
	scn := bufio.NewScanner(file)

	m := u.Matrix2D{}
	for scn.Scan() {
		m = append(m, strings.Split(scn.Text(), ""))
	}

	nextDirection := u.UP
	guard := m.Find("^")

	for m.Includes(guard) {
		v := m.Get(guard.Next(nextDirection))

		if v == "#" {
			nextDirection.TurnClockwise()
			continue
		}

		m.Set(guard, "X")
		guard.Move(nextDirection)
	}

	sum := 0
	for _, line := range m {
		for _, s := range line {
			if s == "X" {
				sum++
			}
		}
	}

	fmt.Println(sum)
}

// --------------------------------------------- PART 2 --------------------------------------------- //

func D6P2() {

	file, _ := os.Open("./inputs/d6.txt")
	scn := bufio.NewScanner(file)

	original := u.Matrix2D{}
	for scn.Scan() {
		original = append(original, strings.Split(scn.Text(), ""))
	}

	nextDirection := u.UP
	originalGuard := original.Find("^")
	guard := originalGuard

	sum := 0
	for lIdx, line := range original {
		for cIdx := range line {
			m := original.ToCopy()
			guard = originalGuard
			nextDirection = u.UP
			if m.Get(u.Coord2D{L: lIdx, C: cIdx}) == "." {
				m.Set(u.Coord2D{L: lIdx, C: cIdx}, "#")
			}
			isLooping := false
			for m.Includes(guard) {
				v := m.Get(guard.Next(nextDirection))

				if v == "#" {
					nextDirection.TurnClockwise()
					continue
				}

				switch nextDirection {
				case u.UP:
					if m.Get(guard) == "U" {
						isLooping = true
						break
					}
					m.Set(guard, "U")

				case u.DOWN:
					if m.Get(guard) == "D" {
						isLooping = true
						break
					}
					m.Set(guard, "D")

				case u.LEFT:
					if m.Get(guard) == "L" {
						isLooping = true
						break
					}
					m.Set(guard, "L")

				case u.RIGHT:
					if m.Get(guard) == "R" {
						isLooping = true
						break
					}
					m.Set(guard, "R")
				}

				if isLooping {
					sum++
					break
				}
				guard.Move(nextDirection)
			}
		}
	}

	fmt.Println(sum)
}
