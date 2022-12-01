package main

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	if err := realMain(); err != nil {
		panic(err)
	}
}

type Elf []int

func ParseElf(raw []string) (Elf, error) {
	var e Elf
	for _, s := range raw {
		calories, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		e = append(e, calories)
	}

	return e, nil
}

func (e Elf) Sum() int {
	s := 0
	for _, i := range e {
		s += i
	}

	return s
}

func realMain() error {
	var elfs []Elf
	for num, rawElf := range strings.Split(input, "\n\n") {
		elf, err := ParseElf(strings.Split(rawElf, "\n"))
		if err != nil {
			return fmt.Errorf("cannot parse elf #%d: %w", num, err)
		}

		elfs = append(elfs, elf)
	}

	sort.Slice(elfs, func(i, j int) bool {
		return elfs[i].Sum() < elfs[j].Sum()
	})

	fmt.Printf("largest: %d\n", elfs[len(elfs)-1].Sum())

	top3Sum := 0
	for i := 1; i <= 3; i++ {
		top3Sum += elfs[len(elfs)-i].Sum()
	}

	fmt.Printf("top 3 sum: %d\n", top3Sum)

	return nil
}
