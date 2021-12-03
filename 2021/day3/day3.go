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

	g := gamma(lines)
	e := epsilon(g)
	fmt.Println(g, e, g*e)
}

func gamma(lines []string) int {
	ones := countOnes(lines)
	var v int16
	for i, count := range ones {
		if count > len(lines)/2 {
			v |= 1 << (11 - i)
		}
	}
	return int(v)
}

func epsilon(gamma int) int {
	mask := (1 << 11) - 1
	return ^gamma & mask
}

func countOnes(lines []string) (r [12]int) {
	for _, line := range lines {
		for i, v := range line {
			if v == '1' {
				r[i]++
			}
		}
	}

	return
}
