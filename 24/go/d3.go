package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// which section is expected of the parser
type PARTS int

const (
	MUL PARTS = iota
	Opening
	N1
	N2
)

func D3P1() {
	file, _ := os.Open("./inputs/d3.txt")
	defer file.Close()

	scn := bufio.NewScanner(file)
	scn.Split(bufio.ScanRunes)

	mulBuffer := []byte{0, 0, 0}
	numBuffer1 := ""
	numBuffer2 := ""
	expecting := MUL

	sum := 0
	for scn.Scan() {
		switch expecting {
		case MUL:
			mulBuffer[0] = mulBuffer[1]
			mulBuffer[1] = mulBuffer[2]
			mulBuffer[2] = scn.Bytes()[0]
			if mulBuffer[0] == 'm' && mulBuffer[1] == 'u' && mulBuffer[2] == 'l' {
				expecting = Opening
				mulBuffer = []byte{0, 0, 0}
			}

		case Opening:
			if scn.Bytes()[0] == '(' {
				expecting = N1
			} else {
				expecting = MUL
			}

		case N1:
			if '0' <= scn.Bytes()[0] && scn.Bytes()[0] <= '9' {
				numBuffer1 += scn.Text()
			} else {
				// must be a comma to end the number
				if numBuffer1 == "" {
					// it needs to be ate least one number for the first parameter
					// back to the start
					expecting = MUL
					break
				}

				if scn.Bytes()[0] == ',' {
					// success
					expecting = N2
				} else {
					// something other than a comma appeared after the number
					numBuffer1 = ""
					expecting = MUL
				}
			}

		case N2:
			if '0' <= scn.Bytes()[0] && scn.Bytes()[0] <= '9' {
				numBuffer2 += scn.Text()
			} else {
				// must be a ')' to end the number
				if numBuffer2 == "" {
					// it needs to be ate least one number for the second parameter
					// back to the start
					numBuffer1 = ""
					expecting = MUL
					break
				}

				if scn.Bytes()[0] == ')' {
					// success
					n1, _ := strconv.Atoi(numBuffer1)
					n2, _ := strconv.Atoi(numBuffer2)

					sum += n1 * n2
				}

				// something other than a ')' appeared after the number
				numBuffer1 = ""
				numBuffer2 = ""
				expecting = MUL

			}

		}
	}
	fmt.Println(sum)
}

func D3P1_alt() {

	input, _ := os.ReadFile("./inputs/d3.txt")

	re, _ := regexp.Compile(`mul\((?:(\d+),(\d+))\)`)

	match := re.FindAllStringSubmatch(string(input), -1)

	sum := 0
	for _, group := range match {
		n1, _ := strconv.Atoi(group[1])
		n2, _ := strconv.Atoi(group[2])

		sum += n1 * n2
	}
	fmt.Println(sum)
}

func D3P2() {
	file, _ := os.Open("./inputs/d3.txt")
	defer file.Close()

	scn := bufio.NewScanner(file)
	scn.Split(bufio.ScanRunes)

	enable := true
	doBuffer := make([]byte, 7)

	mulBuffer := []byte{0, 0, 0}
	numBuffer1 := ""
	numBuffer2 := ""
	expecting := MUL

	sum := 0
	for scn.Scan() {
		if enable {
			switch expecting {
			case MUL:
				mulBuffer[0] = mulBuffer[1]
				mulBuffer[1] = mulBuffer[2]
				mulBuffer[2] = scn.Bytes()[0]
				if mulBuffer[0] == 'm' && mulBuffer[1] == 'u' && mulBuffer[2] == 'l' {
					expecting = Opening
					mulBuffer = []byte{0, 0, 0}
				}

			case Opening:
				if scn.Bytes()[0] == '(' {
					expecting = N1
				} else {
					expecting = MUL
				}

			case N1:
				if '0' <= scn.Bytes()[0] && scn.Bytes()[0] <= '9' {
					numBuffer1 += scn.Text()
				} else {
					// must be a comma to end the number
					if numBuffer1 == "" {
						// it needs to be ate least one number for the first parameter
						// back to the start
						expecting = MUL
						break
					}

					if scn.Bytes()[0] == ',' {
						// success
						expecting = N2
					} else {
						// something other than a comma appeared after the number
						numBuffer1 = ""
						expecting = MUL
					}
				}

			case N2:
				if '0' <= scn.Bytes()[0] && scn.Bytes()[0] <= '9' {
					numBuffer2 += scn.Text()
				} else {
					// must be a ')' to end the number
					if numBuffer2 == "" {
						// it needs to be ate least one number for the second parameter
						// back to the start
						numBuffer1 = ""
						expecting = MUL
						break
					}

					if scn.Bytes()[0] == ')' {
						// success
						n1, _ := strconv.Atoi(numBuffer1)
						n2, _ := strconv.Atoi(numBuffer2)

						sum += n1 * n2
					}

					// something other than a ')' appeared after the number
					numBuffer1 = ""
					numBuffer2 = ""
					expecting = MUL

				}
			}
		}

		for i := 0; i < len(doBuffer)-1; i++ {
			doBuffer[i] = doBuffer[i+1]
		}
		doBuffer[6] = scn.Bytes()[0]

		if string(doBuffer[3:]) == "do()" {
			enable = true
		} else if string(doBuffer) == "don't()" {
			enable = false
		}
	}
	fmt.Println(sum)
}
