// Package day9 solves the according puzzle from https://adventofcode.com/2020/day/9.
package day9

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func MustParseInts(str string) []int {
	var res []int
	for _, line := range strings.Split(str, "\n") {
		val, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Errorf("unable to parse number line: %w", err))
		}

		res = append(res, val)
	}

	return res
}

func FindFirstNonXMAS(xmas []int, windowSize int) int {
NextValue:
	for i := windowSize; i < len(xmas); i++ {
		for a := i - windowSize; a < i; a++ {
			for b := i - windowSize; b < i; b++ {
				if a == b {
					continue
				}

				if xmas[i] == xmas[a]+xmas[b] {
					continue NextValue
				}
			}

		}

		return xmas[i] // we can ever reach this either if everything is fine (not possible) or if no sum exists
	}

	return -1
}

func FindFirstContiguous(xmas []int, target int) int {
Next:
	for i := range xmas {
		sum := xmas[i]
		for x := i + 1; x < len(xmas); x++ {
			sum += xmas[x]
			if sum == target {
				return minMaxSum(xmas[i : x+1])
			}

			if sum > target { // don't do unnecessary checks
				continue Next
			}
		}
	}

	return -1
}

func minMaxSum(slice []int) int {
	min := math.MaxInt32
	max := math.MinInt32

	for _, v := range slice {
		if min > v {
			min = v
		} else if max < v { // unclear if max and min can be the same, but puzzle don't has it
			max = v
		}
	}

	return min + max
}
