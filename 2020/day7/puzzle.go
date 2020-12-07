// Package day7 solves the according puzzle from https://adventofcode.com/2020/day/7.
package day7

import (
	"fmt"
	"strconv"
	"strings"
)

// Bag is a slice of bag names. This started a bit awkward, so the design is a bit ugly, due to
// redundant bag names.
type Bag []string

// NewBag creates a new bag with other bag names.
func NewBag(slices ...[]string) Bag {
	tmp := make(Bag, 0, 10)
	for _, slice := range slices {
		tmp = append(tmp, slice...)
	}

	return tmp
}

// Contains returns true if this bag Contains at least one other bag with the given name.
func (b Bag) Contains(name string) bool {
	for _, s := range b {
		if s == name {
			return true
		}
	}

	return false
}

// Rules is a custom type for named Bags and the root of our rules.
type Rules map[string]Bag

// ParseRules is the factory for reading that rules text format.
func ParseRules(lines string) (Rules, error) {
	res := Rules{}

	for _, line := range strings.Split(lines, "\n") {
		name := line[0 : strings.Index(line, "bags")-1]
		values := strings.Split(line[strings.Index(line, "contain")+8:], ",")

		var params []string

		if values[0] != "no other bags." {
			for _, value := range values {
				value = strings.TrimSpace(value)
				numSep := strings.Index(value, " ")

				count, err := strconv.Atoi(value[:numSep])
				if err != nil {
					return nil, fmt.Errorf("unable to parse num from rule: %w", err)
				}

				params = append(params, put(count, value[numSep+1:strings.Index(value, "bag")-1])...)
			}
		}

		res[name] = NewBag(params)
	}

	return res, nil
}

// put is a helper which I have created because I originally thought it would make some things easier, but that was
// just a stupid idea.
func put(n int, str string) []string {
	tmp := make([]string, 0, n)
	for i := 0; i < n; i++ {
		tmp = append(tmp, str)
	}

	return tmp
}

// CountBagsWhichMayContainAtLeast is actually a java-style name, well its late already. However, nice
// thing about this implementation is, that it is the imperative way of a recursion and would not suffer
// from a stack overflow in other languages.
func (r Rules) CountBagsWhichMayContainAtLeast(n string) int {
	visited := map[string]bool{}

	var seeds []string
	seeds = append(seeds, n)

	for len(seeds) > 0 {
		seed := seeds[len(seeds)-1]
		seeds = seeds[:len(seeds)-1]

		for name, bag := range r {
			if bag.Contains(seed) {
				seeds = append(seeds, name)
				visited[name] = true
			}
		}
	}

	return len(visited)
}

// CountTotalBagsFor is a simple recursive counter.
func (r Rules) CountTotalBagsFor(n string) int {
	count := 0
	bag := r[n]

	count += len(bag)
	for _, bagName := range bag {
		count += r.CountTotalBagsFor(bagName)
	}

	return count
}
