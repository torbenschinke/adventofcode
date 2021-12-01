package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := flag.String("file", "2021/day1/input.txt", "the measurements")
	flag.Parse()

	nums, err := parseNums(*path)
	if err != nil {
		log.Println(fmt.Errorf("cannot parse nums from %s: %w", *path, err))
		os.Exit(1)
	}

	fmt.Printf("part 1: %d\n", countIncs(nums))
	fmt.Printf("part 2: %d\n", countSlideWindowIncs(nums, 3))
}

func parseNums(path string) ([]int, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	var nums []int
	for i, str := range strings.Split(string(buf), "\n") {
		v, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("invalid number in line %d: %w", i, err)
		}

		nums = append(nums, v)
	}

	return nums, nil
}

func countIncs(nums []int) int {
	count := 0
	for i, num := range nums {
		if i > 0 && num > nums[i-1] {
			count++
		}
	}

	return count
}

func sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

func countSlideWindowIncs(nums []int, width int) int {
	incs := 0
	for i := 0; i < len(nums)-width; i++ {
		a := sum(nums[i : i+width])
		b := sum(nums[i+1 : i+width+1])
		if b > a {
			incs++
		}
	}

	return incs
}
