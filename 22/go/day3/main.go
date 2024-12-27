// removing the duplicated itens before putting in the 'inBoth' slice could be better, but it would require some testing to be sure

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getPriority(str string) int {
	if str == "" {
		return 0
	}

	codePoint := []byte(str)[0]
	if codePoint > 96 { // lowercase
		return int(codePoint) - 96
	}
	// UPPERCASE
	return int(codePoint) - 38 // -64+26
}

func main() {
	PartTwo()
}

func partOne() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalPriority := 0
	for scanner.Scan() {
		// 🫣
		fullSack := strings.Split(scanner.Text(), "")
		comp1 := fullSack[:len(fullSack)/2]
		comp2 := fullSack[len(fullSack)/2:]
		inBoth := make([]string, 0, len(comp1))
		for _, item1 := range comp1 {
			for _, item2 := range comp2 {
				if item1 == item2 {
					alreadyFound := false
					for _, foundItem := range inBoth {
						if item1 == foundItem {
							alreadyFound = true
							break
						}
					}
					if !alreadyFound {
						inBoth = append(inBoth, item1)
					}
					break
				}
			}
		}
		for _, item := range inBoth {
			totalPriority += getPriority(item)
		}
	}
	fmt.Println(totalPriority)
}

func PartTwo() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	totalPriority := 0
	for scanner.Scan() {
		sack1 := strings.Split(scanner.Text(), "")
		scanner.Scan()
		sack2 := strings.Split(scanner.Text(), "")
		scanner.Scan()
		sack3 := strings.Split(scanner.Text(), "")
		alreadySearched := make([]string, 0, len(sack1))
		badge := ""
		for _, item1 := range sack1 {
			keepSearchingItem := true
			for _, ASItem := range alreadySearched {
				if item1 == ASItem {
					keepSearchingItem = false
					break
				}
			}
			if keepSearchingItem {
				alreadySearched = append(alreadySearched, item1)
				for _, item2 := range sack2 {
					if item2 == item1 {
						for _, item3 := range sack3 {
							if item3 == item1 {
								badge = item1
								break
							}
						}
						break
					}
				}
			}
			if badge != "" {
				break
			}
		}
		totalPriority += getPriority(badge)
	}
	fmt.Println(totalPriority)
}
