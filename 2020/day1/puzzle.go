// Package day1 solves https://adventofcode.com/2020/day/1.
package day1

import (
	"fmt"
	"strconv"
)

// FindBySum2 returns the first values a and b which when added, are equal to sum.
// If no such sum is found, the returned values are -1.
func FindBySum2(nums []int, sum int) (a, b int) {
	for _, a := range nums {
		for _, b := range nums {
			if a+b == sum {
				return a, b
			}
		}
	}

	return -1, -1
}

// FindBySum3 returns the first values a, b and c which when added, are equal to sum.
// If no such sum is found, the returned values are -1. We may consider sorting
// the numbers to cancel checking early and avoid O(n^3).
func FindBySum3(nums []int, sum int) (a, b, c int) {
	for _, a := range nums {
		for _, b := range nums {
			for _, c := range nums {
				if a+b+c == sum {
					return a, b, c
				}
			}
		}
	}

	return -1, -1, -1
}

// MustParseNums panics, if in cannot be parsed.
func MustParseNums(in string) []int {
	r, err := ParseNums(in)
	if err != nil {
		panic(err)
	}

	return r
}

// ParseNums tries to parse the given numbers, one per line.
func ParseNums(in string) ([]int, error) {
	tmp := make([]rune, 0, 4)

	var nums []int

	for _, r := range in {
		if r == '\n' {
			i, err := strconv.Atoi(string(tmp))
			if err != nil {
				return nums, fmt.Errorf("cannot parse integer in line: %w", err)
			}

			nums = append(nums, i)
			tmp = tmp[:0]

			continue
		}

		tmp = append(tmp, r)
	}

	return nums, nil
}
