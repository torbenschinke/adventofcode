package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Item byte
type Rucksack string

func (r Rucksack) FirstCompartment() Rucksack {
	return r[:len(r)/2]
}

func (r Rucksack) SecondCompartment() Rucksack {
	return r[len(r)/2:]
}

func (r Rucksack) CommonItem() Item {
	for _, runeFromFirst := range r.FirstCompartment() {
		for _, runeFromSecond := range r.SecondCompartment() {
			if runeFromSecond == runeFromFirst {
				return Item(runeFromSecond)
			}
		}
	}

	panic("illegal state")
}

func (r Rucksack) uniqueItems() ItemCounts {
	tmp := ItemCounts{}
	for _, i := range r {
		tmp[Item(i)] = 1
	}

	return tmp
}

type ItemCounts map[Item]int

func (c ItemCounts) Add(other ItemCounts) ItemCounts {
	tmp := ItemCounts{}
	for k, v := range other {
		tmp[k] = v + c[k]
	}
	return tmp
}

func CommonItem(r []Rucksack) Item {
	sums := ItemCounts{}
	for _, rucksack := range r {
		sums = sums.Add(rucksack.uniqueItems())
	}

	for k, v := range sums {
		if v == len(r) {
			return k
		}
	}

	panic("illegal state")
}

func priority(item Item) int {
	if item >= 'a' {
		return int(item - 'a' + 1)
	}

	return int(item-'A') + 27
}

func part1() {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		sum += priority(Rucksack(line).CommonItem())
	}

	fmt.Printf("part 1: %d\n", sum)
}

func part2() {
	sum := 0
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i += 3 {
		var group []Rucksack
		for g := i; g < i+3; g++ {
			group = append(group, Rucksack(lines[g]))
		}
		sum += priority(CommonItem(group))
	}
	fmt.Printf("part 2: %d\n", sum)
}

func main() {
	part1()
	part2()
}
