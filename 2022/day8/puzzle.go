package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Tree struct {
	Height  int
	Visible int
}

type TreeMap [][]Tree

func (m TreeMap) Tree(x, y int) *Tree {
	return &(m[y][x])
}

func (m TreeMap) Width() int {
	return len(m[0])
}
func (m TreeMap) Height() int {
	return len(m)
}

func (m TreeMap) InBounds(x, y int) bool {
	return x < m.Width() && x >= 0 && y < m.Height() && y >= 0
}

func (m TreeMap) Each(f func(t *Tree)) {
	for x := 0; x < m.Width(); x++ {
		for y := 0; y < m.Height(); y++ {
			f(m.Tree(x, y))
		}
	}
}

type Vec struct {
	X, Y int
}

func (m TreeMap) Mark(x, y int, v Vec) {
	if !m.InBounds(x, y) {
		return
	}
	tree := m.Tree(x, y)
	if x == 0 || x == m.Width()-1 || y == 0 || y == m.Height()-1 {
		tree.Visible++
		x, y = x+v.X, y+v.Y
		m.Mark(x, y, v)
		return
	}
	x, y = x+v.X, y+v.Y
	if !m.InBounds(x, y) {
		return
	}
	nextTree := m.Tree(x, y)

	if nextTree.Height < tree.Height {
		tree.Visible++
		m.Mark(x, y, v)
	}
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

func part1() {
	trees := Parse(input)

	// from top/bottom
	for x := 0; x < trees.Width(); x++ {
		trees.Mark(x, 0, Vec{0, 1})
		trees.Mark(x, trees.Height()-1, Vec{0, -1})
	}

	// from left/right
	for y := 0; y < trees.Height(); y++ {
		trees.Mark(0, y, Vec{1, 0})
		trees.Mark(trees.Width()-1, y, Vec{-1, 0})
	}

	sum := 0
	trees.Each(func(t *Tree) {
		if t.Visible > 0 {
			sum++
		}
	})

	fmt.Printf("part1: %d\n", sum)
}

func main() {
	part1()
}
