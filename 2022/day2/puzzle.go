package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1()
	part2()
}

type Turn byte

const (
	Rock Turn = 'X' + iota
	Paper
	Scissor
)

func (t Turn) Fight(other Turn) int {
	if t.Paper() && other.Paper() || t.Rock() && other.Rock() || t.Scissors() && other.Scissors() {
		return 3
	}

	if t.Rock() && other.Scissors() || t.Scissors() && other.Paper() || t.Paper() && other.Rock() {
		return 6
	}

	return 0
}

func (t Turn) Choose(other Turn) Turn {
	for _, picked := range []Turn{Rock, Paper, Scissor} {
		if picked.Fight(t) == other.RequiredScore() {
			return picked
		}
	}
	panic("illegal state")
}

func (t Turn) RequiredScore() int {
	switch t {
	case 'X':
		return 0
	case 'Y':
		return 3
	case 'Z':
		return 6
	}
	panic("illegal state")
}

func (t Turn) Rock() bool     { return t == 'A' || t == 'X' }
func (t Turn) Paper() bool    { return t == 'B' || t == 'Y' }
func (t Turn) Scissors() bool { return t == 'C' || t == 'Z' }

func part1() {
	score := 0
	for _, line := range strings.Split(input, "\n") {
		other, me := Turn(line[0]), Turn(line[2])
		score += int(me - 'X' + 1)
		score += me.Fight(other)
	}

	fmt.Printf("part 1: %d\n", score)
}

func part2() {
	score := 0
	for _, line := range strings.Split(input, "\n") {
		other, me := Turn(line[0]), Turn(line[2])
		picked := other.Choose(me)
		score += int(picked - 'X' + 1)
		score += picked.Fight(other)
	}

	fmt.Printf("total score: %d\n", score)
}
