// Package day1 solves https://adventofcode.com/2020/day/1.
package day1

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
// If no such sum is found, the returned values are -1.
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
