package main

import (
	"aoc_24/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Before int
	After  int
}

func (r Rule) Contains(n int) bool {
	return r.Before == n || r.After == n
}

func getRuleSubset(rules []Rule, n int) []Rule {
	return utils.Filter(rules, func(r Rule) bool {
		return r.Contains(n)
	})
}

func isSorted(pages []int, rules []Rule) bool {
	for idx, target := range pages {
		rulesSubset := getRuleSubset(rules, target)

		for compIdx, comp := range pages {
			if compIdx < idx {
				if utils.Any(rulesSubset, func(r Rule) bool {
					return r.Before == target && r.After == comp
				}) {
					return false
				}
			} else if idx < compIdx {
				if utils.Any(rulesSubset, func(r Rule) bool {
					return r.Before == comp && r.After == target
				}) {
					return false
				}
			}
		}
	}
	return true
}

func D5P1() {
	file, _ := os.Open("./inputs/d5.txt")

	scn := bufio.NewScanner(file)

	rulesEnded := false

	var updates [][]int
	var rules []Rule

	for scn.Scan() {
		if scn.Text() == "" {
			rulesEnded = true
			continue
		}

		if rulesEnded {
			updates = append(updates, utils.Map(strings.Split(scn.Text(), ","), func(s string) int {
				i, _ := strconv.Atoi(s)
				return i
			}))
		} else {
			split := strings.Split(scn.Text(), "|")
			i1, _ := strconv.Atoi(split[0])
			i2, _ := strconv.Atoi(split[1])
			rules = append(rules, Rule{Before: i1, After: i2})
		}
	}

	sum := 0
	for _, pages := range updates {
		isCorrect := true
		for idx, target := range pages {
			if !isCorrect {
				break
			}

			rulesSubset := getRuleSubset(rules, target)

			for compIdx, comp := range pages {
				if compIdx < idx {
					if utils.Any(rulesSubset, func(r Rule) bool {
						return r.Before == target && r.After == comp
					}) {
						isCorrect = false
						break
					}
				} else if idx < compIdx {
					if utils.Any(rulesSubset, func(r Rule) bool {
						return r.Before == comp && r.After == target
					}) {
						isCorrect = false
						break
					}
				}
			}
		}

		if isCorrect {
			sum += pages[len(pages)/2]
		}
	}

	fmt.Println(sum)
}

// --------------------------------------------- PART 2 --------------------------------------------- //

func D5P2() {
	file, _ := os.Open("./inputs/d5.txt")

	scn := bufio.NewScanner(file)

	rulesEnded := false

	var updates [][]int
	var rules []Rule

	for scn.Scan() {
		if scn.Text() == "" {
			rulesEnded = true
			continue
		}

		if rulesEnded {
			updates = append(updates, utils.Map(strings.Split(scn.Text(), ","), func(s string) int {
				i, _ := strconv.Atoi(s)
				return i
			}))
		} else {
			split := strings.Split(scn.Text(), "|")
			i1, _ := strconv.Atoi(split[0])
			i2, _ := strconv.Atoi(split[1])
			rules = append(rules, Rule{Before: i1, After: i2})
		}
	}

	var badUpdates [][]int
	for _, pages := range updates {
		isCorrect := true
		for idx, target := range pages {
			if !isCorrect {
				break
			}

			rulesSubset := getRuleSubset(rules, target)

			for compIdx, comp := range pages {
				if compIdx < idx {
					if utils.Any(rulesSubset, func(r Rule) bool {
						return r.Before == target && r.After == comp
					}) {
						badUpdates = append(badUpdates, pages)
						isCorrect = false
						break
					}
				} else if idx < compIdx {
					if utils.Any(rulesSubset, func(r Rule) bool {
						return r.Before == comp && r.After == target
					}) {
						badUpdates = append(badUpdates, pages)
						isCorrect = false
						break
					}
				}
			}
		}
	}

	sum := 0
	for _, pages := range badUpdates {
		sorted := make([]int, len(pages))
		copy(sorted, pages)

		// swap everytime a rule is broken and pray that it gets sorted at some point
		for !isSorted(sorted, rules) {

			for idx, target := range sorted {
				rulesSubset := getRuleSubset(rules, target)
				changed := false
				for compIdx, comp := range sorted {
					if compIdx < idx {
						if utils.Any(rulesSubset, func(r Rule) bool {
							return r.Before == target && r.After == comp
						}) {
							changed = true
							sorted[idx] = comp
							sorted[compIdx] = target
							break
						}
					} else if idx < compIdx {
						if utils.Any(rulesSubset, func(r Rule) bool {
							return r.Before == comp && r.After == target
						}) {
							changed = true
							sorted[idx] = comp
							sorted[compIdx] = target
							break
						}
					}
				}

				if changed {
					break
				}
			}
		}

		sum += sorted[len(sorted)/2]
	}

	fmt.Println(sum)
}
