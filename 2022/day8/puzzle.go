package main

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input string

var directions = []Vec{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

type Tree struct {
	Height int
}

type TreeMap [][]Tree

func (m TreeMap) Tree(x, y int) *Tree {
	return &(m[y][x])
}

func (m TreeMap) Width() int { return len(m[0]) }

func (m TreeMap) Height() int { return len(m) }

func (m TreeMap) InBounds(x, y int) bool {
	return x < m.Width() && x >= 0 && y < m.Height() && y >= 0
}

func (m TreeMap) Each(f func(x, y int, t *Tree)) {
	for x := 0; x < m.Width(); x++ {
		for y := 0; y < m.Height(); y++ {
			f(x, y, m.Tree(x, y))
		}
	}
}

type Vec struct {
	X, Y int
}

func (m TreeMap) Trace(x, y, height int, v Vec) (Vec, bool) {
	x, y = x+v.X, y+v.Y
	if !m.InBounds(x, y) {
		return Vec{x - v.X, y - v.Y}, true
	}
	nextTree := m.Tree(x, y)
	if nextTree.Height >= height {
		return Vec{x, y}, false
	}

	return m.Trace(x, y, height, v)
}

func Parse(input string) TreeMap {
	var m TreeMap
	for _, line := range strings.Split(input, "\n") {
		var row []Tree
		for _, r := range line {
			row = append(row, Tree{Height: int(r - '0')})
		}
		m = append(m, row)
	}
	return m
}

func main() {
	trees := Parse(input)
	outsideScore := 0
	scenicScore := 1.0
	trees.Each(func(x, y int, t *Tree) {
		score := 1.0
		wasVisible := false
		for _, dir := range directions {
			pos, hitBound := trees.Trace(x, y, t.Height, dir)
			wasVisible = wasVisible || hitBound
			dx, dy := math.Abs(float64(x-pos.X)), math.Abs(float64(y-pos.Y))
			score *= dx + dy
		}

		if wasVisible {
			outsideScore++
		}
		if scenicScore < score {
			scenicScore = score
		}
	})

	fmt.Printf("part1: %d\n", outsideScore)
	fmt.Printf("part2: %d\n", int(scenicScore))
}
