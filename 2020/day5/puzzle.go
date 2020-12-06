// Package day5 solves the according puzzle from https://adventofcode.com/2020/day/5.
package day5

import (
	"math"
	"sort"
	"strings"
)

// A Pass looks like FBFBBFFRLR and has 10 bytes, using only the characters F, B, R and L.
type Pass string

// Row returns the row of the seat in the aircraft.
func (p Pass) Row() int {
	return search(p[:len(p)-3], 0, 127)
}

// Column returns the seat column in the aircraft.
func (p Pass) Column() int {
	return search(p[len(p)-3:], 0, 7)
}

// ID returns the seat ID.
func (p Pass) ID() int {
	return p.Row()*8 + p.Column()
}

// search performs a recursive "binary search" to resolve a row or seat.
func search(next Pass, min, max int) int {
	delta := int(math.Floor((float64(max) - float64(min)) / 2)) //nolint:gomnd

	switch next[0] {
	case 'L':
		fallthrough
	case 'F': // lower half
		max = min + delta

		if len(next) == 1 {
			return max
		}
	case 'R':
		fallthrough
	case 'B': // upper half
		min += delta

		if len(next) == 1 {
			return max
		}
	}

	return search(next[1:], min, max)
}

// HighestSeatID interprets each line of the passes string as a Pass and returns the largest Pass.ID.
func HighestSeatID(passes string) int {
	max := math.MinInt32

	for _, pass := range strings.Split(passes, "\n") {
		id := Pass(pass).ID()
		if id > max {
			max = id
		}
	}

	return max
}

// FindFirstMissingSeatID returns the first Pass.ID which is missing (means that e.g. ID 4 and 6 is available but 5
// is missing).
func FindFirstMissingSeatID(passes string) int {
	allPassIDs := make([]int, 0, 10)

	for _, p := range strings.Split(passes, "\n") {
		pass := Pass(p)
		allPassIDs = append(allPassIDs, pass.ID())
	}

	sort.Ints(allPassIDs)

	lastID := -1
	for _, id := range allPassIDs {
		if lastID == -1 {
			lastID = id

			continue
		}

		if id-lastID > 1 {
			return id - 1
		}

		lastID = id
	}

	return -1
}
