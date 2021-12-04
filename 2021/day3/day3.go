package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	path := flag.String("file", "2021/day3/input.txt", "the commands")
	flag.Parse()

	buf, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatalf("cannot read file: %v\n", err)
	}

	lines := strings.Split(string(buf), "\n")
	report := loadReport(lines)

	g := gamma(report)
	e := epsilon(g)
	fmt.Printf("gamma = %d, epsilon = %d, power = %d\n", g, e, g*e)

	oxy := oxy(report)
	co2 := co2(report)
	fmt.Printf("oxygen = %d, co2 = %d, life support = %d\n", oxy, co2, oxy*co2)
}

type Pattern []bool

func (p Pattern) int() int {
	var v int
	for i, b := range p {
		if b {
			v |= 1 << (len(p) - 1 - i)
		}
	}

	return v
}
func (p Pattern) count() (ones, zeros int) {
	for _, b := range p {
		if b {
			ones++
		} else {
			zeros++
		}
	}

	return
}

type Report []Pattern

func (r Report) width() int { return len(r[0]) }

func (r Report) vert(idx int) Pattern {
	var p Pattern
	for _, pattern := range r {
		p = append(p, pattern[idx])
	}

	return p
}

func (r Report) filter(predicate func(Pattern) bool) Report {
	var c Report
	for _, pattern := range r {
		if predicate(pattern) {
			c = append(c, pattern)
		}
	}

	return c
}

func search(r Report, p func(ones, zeros int) bool) int {
	pos := 0
	for len(r) > 1 {
		ones, zeros := r.vert(pos).count()
		highOrLow := p(ones, zeros)

		r = r.filter(Select(highOrLow, pos))
		pos++
	}

	return r[0].int()
}

func oxy(r Report) int {
	return search(r, func(ones, zeros int) bool { return ones >= zeros })
}

func co2(r Report) int {
	return search(r, func(ones, zeros int) bool { return ones < zeros })
}

func Select(high bool, idx int) func(Pattern) bool {
	return func(pattern Pattern) bool {
		if high {
			return pattern[idx]
		} else {
			return !pattern[idx]
		}
	}
}

func loadReport(lines []string) Report {
	var r Report
	for _, line := range lines {
		var p Pattern
		for _, b := range line {
			p = append(p, b != '0')
		}

		r = append(r, p)
	}

	return r
}

func gamma(r Report) int {
	var v int
	for i := 0; i < r.width(); i++ {
		ones, zeros := r.vert(i).count()
		if ones > zeros {
			v |= 1 << (r.width() - 1 - i)
		}
	}
	return v
}

func epsilon(gamma int) int {
	mask := (1 << 11) - 1
	return ^gamma & mask
}
