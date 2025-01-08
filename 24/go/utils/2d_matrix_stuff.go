package utils

import "fmt"

type Direction int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

func (d *Direction) TurnClockwise() {
	switch *d {
	case UP:
		*d = RIGHT
	case DOWN:
		*d = LEFT
	case LEFT:
		*d = UP
	case RIGHT:
		*d = DOWN
	}
}

type Coord2D struct {
	L int
	C int
}

func (c *Coord2D) Set(line, col int) {
	c.L = line
	c.C = col
}

func (c Coord2D) Next(d Direction) Coord2D {
	switch d {
	case UP:
		return Coord2D{c.L - 1, c.C}
	case DOWN:
		return Coord2D{c.L + 1, c.C}
	case LEFT:
		return Coord2D{c.L, c.C - 1}
	case RIGHT:
		return Coord2D{c.L, c.C + 1}
	}
	panic("??????????")
}

func (c *Coord2D) Move(d Direction) {
	switch d {
	case UP:
		c.L--
	case DOWN:
		c.L++
	case LEFT:
		c.C--
	case RIGHT:
		c.C++
	}
}

type Matrix2D [][]string

func (m Matrix2D) Get(c Coord2D) string {
	if c.C < 0 || c.C >= len(m[0]) || c.L < 0 || c.L >= len(m) {
		return ""
	}
	return m[c.L][c.C]
}

func (m Matrix2D) Set(c Coord2D, s string) {
	if c.C < 0 || c.C >= len(m[0]) || c.L < 0 || c.L >= len(m) {
		return
	}
	m[c.L][c.C] = s
}

// finds the first occurrence of the target and return its coordinate
func (m Matrix2D) Find(target string) Coord2D {
	for lIdx, line := range m {
		for cIdx, s := range line {
			if s == target {
				return Coord2D{lIdx, cIdx}
			}
		}
	}
	return Coord2D{-1, -1}
}

// Check if the coord is within the matrix boundaries
func (m Matrix2D) Includes(c Coord2D) bool {

	if 0 <= c.L && c.L < len(m) {
		if 0 <= c.C && c.C < len(m[c.L]) {
			return true
		}
	}

	return false
}

func (m Matrix2D) ToCopy() Matrix2D {
	mCopy := make(Matrix2D, len(m))
	for idx, line := range m {
		mCopy[idx] = make([]string, len(line))
		copy(mCopy[idx], line)
	}
	return mCopy
}

func (m Matrix2D) String() string {
	final := ""
	for _, line := range m {
		for _, s := range line {
			final += fmt.Sprintf("%v", s)
		}
		final += "\n"
	}
	return final
}
