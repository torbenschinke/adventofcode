package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	path := flag.String("file", "2021/day2/input.txt", "the commands")
	flag.Parse()

	buf, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatalf("cannot read file: %v\n", err)
	}

	lines := strings.Split(string(buf), "\n")
	x, y, err := part1(lines)
	if err != nil {
		log.Fatalf("failed part 1: %v\n", err)
	}

	fmt.Printf("x=%d, y=%d => %d\n", x, y, x*y)

	x, y, err = part2(lines)
	if err != nil {
		log.Fatalf("failed part 2: %v\n", err)
	}

	fmt.Printf("x=%d, y=%d => %d\n", x, y, x*y)
}

func parse(lines []string, f func(cmd byte, val int)) error {
	for i, line := range lines {
		cmd := line[0]
		val, err := strconv.Atoi(string(line[len(line)-1]))
		if err != nil {
			return fmt.Errorf("invalid num in line: %d: %w", i, err)
		}

		f(cmd, val)
	}

	return nil
}

func part1(lines []string) (x, y int, err error) {
	err = parse(lines, func(cmd byte, val int) {
		switch cmd {
		case 'f':
			x += val
		case 'd':
			y += val
		case 'u':
			y -= val
		}
	})

	return
}

func part2(lines []string) (x, y int, err error) {
	a := 0
	err = parse(lines, func(cmd byte, val int) {
		switch cmd {
		case 'f':
			x += val
			y += a * val
		case 'd':
			a += val
		case 'u':
			a -= val
		}
	})

	return
}
