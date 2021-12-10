package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	path := flag.String("file", "2021/day10/input.txt", "the commands")
	flag.Parse()

	lines, err := readLines(*path)
	if err != nil {
		log.Fatalf("cannot read file: %v\n", err)
	}

	fmt.Println(part1(lines)) //243939
	fmt.Println(part2(lines)) //2421222841
}

func readLines(fname string) (lines []string, err error) {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	return strings.Split(string(buf), "\n"), nil
}

func part1(lines []string) int {
	score := 0
	for _, line := range lines {
		score += corruptionScore(line)
	}

	return score
}

func part2(lines []string) int {
	var allScores []int
	for _, line := range lines {
		if corruptionScore(line) != 0 {
			continue
		}

		var stack Stack
		for _, r := range line {
			openIdx := bytes.IndexByte(OpenBraces, byte(r))
			if openIdx != -1 {
				stack.Push(byte(r))
			} else {
				_, _ = stack.Pop()
			}
		}

		totalScore := 0
		for len(stack.open) > 0 {
			openBrace, _ := stack.Pop()
			points := bytes.IndexByte(OpenBraces, openBrace) + 1
			totalScore *= 5
			totalScore += points
		}

		allScores = append(allScores, totalScore)
	}

	sort.Ints(allScores)
	return allScores[len(allScores)/2]
}

func corruptionScore(line string) int {
	var stack Stack
	for _, r := range line {
		openIdx := bytes.IndexByte(OpenBraces, byte(r))
		if openIdx != -1 {
			stack.Push(byte(r))
		} else {
			openChar, ok := stack.Pop()
			if !ok {
				panic("cannot pop")
			}
			closeIdx := bytes.IndexByte(OpenBraces, openChar)
			expected := CloseBraces[closeIdx]
			if byte(r) != expected {
				fmt.Printf("expected %s, but found %s instead\n", string(expected), string(r))
				return Points[bytes.IndexByte(CloseBraces, byte(r))]
			}
		}
	}
	return 0
}

var OpenBraces = []byte{'(', '[', '{', '<'}
var CloseBraces = []byte{')', ']', '}', '>'}
var Points = []int{3, 57, 1197, 25137}

type Stack struct {
	open []byte
}

func (s *Stack) Push(c byte) {
	s.open = append(s.open, c)
}

func (s *Stack) Pop() (byte, bool) {
	if len(s.open) == 0 {
		return 0, false
	}
	c := s.open[len(s.open)-1]
	s.open = s.open[:len(s.open)-1]
	return c, true
}
