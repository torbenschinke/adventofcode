package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/bits"
	"os"
)

func main() {
	path := flag.String("file", "2021/day8/input.txt", "the commands")
	flag.Parse()

	entries, err := parse(*path)
	if err != nil {
		log.Fatalf("cannot parse file: %v\n", err)
	}

	fmt.Println(count(entries))
	fmt.Println(resolve(entries)) //1084606
}

func resolve(entries []Entry) int {
	sum := 0
	for _, entry := range entries {
		sum += entry.Display()
	}

	return sum
}

func count(entries []Entry) int {
	count := 0
	what := []int{2, 3, 4, 7}
	for _, entry := range entries {
		for _, i := range what {
			for _, digit := range entry.Digits {
				if bits.OnesCount(uint(digit)) == i {
					count++
				}
			}
		}
	}

	return count
}

func parse(fname string) (entries []Entry, err error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}

	defer file.Close()

	for i := 0; ; i++ {
		var patterns [10]Pattern
		var display [4]Pattern
		_, err = fmt.Fscanf(file, "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s\n",
			&patterns[0], &patterns[1], &patterns[2], &patterns[3], &patterns[4], &patterns[5], &patterns[6], &patterns[7], &patterns[8], &patterns[9],
			&display[0], &display[1], &display[2], &display[3])

		if err == io.EOF {
			return entries, nil
		}

		if err != nil {
			return nil, fmt.Errorf("cannot parse line %d: %w", i, err)
		}

		e := Entry{}
		for i, pattern := range patterns {
			e.Patterns[i] = pattern.Int()
		}

		for i, pattern := range display {
			e.Digits[i] = pattern.Int()
		}
		entries = append(entries, e)
	}
}

type Entry struct {
	Patterns [10]int
	Digits   [4]int
}

const (
	A = 1 << iota
	B
	C
	D
	E
	F
	G
)

var segmentSet = []int{A, B, C, D, E, F, G}

func (e Entry) Display() int {
	_1 := e.pickByLen(2)[0]
	_4 := e.pickByLen(4)[0]
	_7 := e.pickByLen(3)[0]
	_8 := e.pickByLen(7)[0]

	// 3 is the only digit of in total 3 numbers with 5 segments (2,3,5)
	var _3 int
	for _, pattern := range e.pickByLen(5) {
		if pattern&_1 == _1 {
			_3 = pattern
			break
		}
	}

	// 9 is digit 3 | 4
	_9 := _3 | _4

	// 6 segments are either 0,6,9 but we already know 9 so same procedure as with 3
	var _0, _6 int
	for _, pattern := range e.pickByLen(6) {

		if pattern == _9 {
			continue
		}

		if pattern&_1 == _1 {
			_0 = pattern
		} else {
			_6 = pattern
		}
	}

	// missing 2 and 5 (5 segments) => 1+5 = 9
	var _2, _5 int
	for _, pattern := range e.pickByLen(5) {
		if pattern == _3 {
			continue
		}
		if pattern|_1 == _9 {
			_5 = pattern
		} else {
			_2 = pattern
		}
	}

	nums := []int{_0, _1, _2, _3, _4, _5, _6, _7, _8, _9}
	sum := 0
	for _, digit := range e.Digits {
		sum *= 10
		for i, num := range nums {
			if num == digit {
				sum += i
			}
		}

	}

	return sum
}

func (e Entry) pickByLen(l int) []int {
	var r []int
	for _, pattern := range e.Patterns {
		if bits.OnesCount(uint(pattern)) == l {
			r = append(r, pattern)
		}
	}

	return r
}

type Pattern string

func (p Pattern) Int() int {
	v := 0
	for _, i := range p {
		v |= segmentSet[i-'a']
	}

	return v
}
