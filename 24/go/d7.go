package main

import (
	"aoc_24/utils"
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func generatePermutations(size int) [][]string {
	permutations := [][]string{}

	numPermutations := int(math.Pow(2, float64(size)))

	for i := 0; i < numPermutations; i++ {
		currentPermutation := []string{}

		// Generate the current permutation using binary numbers as a basis: 101 == *+*
		for j := 0; j < size; j++ {
			// Check if the j-th bit of i is set (1)
			if i&(1<<j) != 0 {
				currentPermutation = append(currentPermutation, "*")
			} else {
				currentPermutation = append(currentPermutation, "+")
			}
		}

		// Add the current permutation to the result
		permutations = append(permutations, currentPermutation)
	}
	return permutations
}

func D7P1() {

	file, _ := os.Open("./inputs/d7.txt")
	scn := bufio.NewScanner(file)

	allPermutations := make(map[int][][]string) // key = size of the sequence, value = sequence of possible permutations
	sum := 0
	for scn.Scan() {
		split := strings.Split(scn.Text(), " ")
		target, _ := strconv.Atoi(split[0][:len(split[0])-1])
		operands := utils.Map(split[1:], func(v string) int {
			i, _ := strconv.Atoi(v)
			return i
		})

		perms, ok := allPermutations[len(operands)-1]

		if !ok {
			allPermutations[len(operands)-1] = generatePermutations(len(operands) - 1)
			perms = allPermutations[len(operands)-1]
		}

		for _, perm := range perms {
			res := operands[0]
			for idx, op := range perm {
				switch op {
				case "+":
					res += operands[idx+1]
				case "*":
					res *= operands[idx+1]
				}

			}
			if res == target {
				sum += target
				break
			}
		}
	}
	fmt.Println(sum)

}

// --------------------------------------------- PART 2 --------------------------------------------- //

func generatePermutationsP2(size int) [][]string {
	permutations := [][]string{}

	numPermutations := int(math.Pow(3, float64(size)))

	for i := 0; i < numPermutations; i++ {
		currentPermutation := []string{}
		ternary := big.NewInt(int64(i)).Text(3)

		// padding so all ternary numbers have the same lenght as the biggest one
		for len(ternary) < len(big.NewInt(int64(numPermutations-1)).Text(3)) {
			ternary = "0" + ternary
		}

		for _, char := range ternary {
			switch char {
			case '0':
				currentPermutation = append(currentPermutation, "+")
			case '1':
				currentPermutation = append(currentPermutation, "*")
			case '2':
				currentPermutation = append(currentPermutation, "||")
			}
		}
		permutations = append(permutations, currentPermutation)
	}

	return permutations
}

func D7P2() {

	file, _ := os.Open("./inputs/d7.txt")
	scn := bufio.NewScanner(file)

	allPermutations := make(map[int][][]string) // key = size of the sequence, value = sequence of possible permutations
	sum := 0
	for scn.Scan() {
		split := strings.Split(scn.Text(), " ")
		target, _ := strconv.Atoi(split[0][:len(split[0])-1])
		operands := utils.Map(split[1:], func(v string) int {
			i, _ := strconv.Atoi(v)
			return i
		})

		perms, ok := allPermutations[len(operands)-1]

		if !ok {
			allPermutations[len(operands)-1] = generatePermutationsP2(len(operands) - 1)
			perms = allPermutations[len(operands)-1]
		}

		for _, perm := range perms {
			res := operands[0]
			for idx, op := range perm {
				switch op {
				case "+":
					res += operands[idx+1]

				case "*":
					res *= operands[idx+1]

				case "||":
					res, _ = strconv.Atoi(fmt.Sprintf("%v%v", res, operands[idx+1]))

				}

			}
			if res == target {
				sum += target
				break
			}
		}
	}
	fmt.Println(sum)

}
