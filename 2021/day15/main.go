package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	path := flag.String("file", "2021/day15/testinput.txt", "the commands")
	flag.Parse()

	field, err := readField(*path)
	if err != nil {
		log.Fatalf("cannot read file: %v\n", err)
	}

	fmt.Println(part1(field))
}

func part1(field Field) int {
	src, dst := Pos{0, 0}, Pos{len(field) - 1, len(field[len(field)-1]) - 1}
	totalCost := 0
	for {
		pos, cost := estimate(field, src, dst, totalCost, 20)
		fmt.Println(pos, cost)
		if pos == dst {
			break
		}
		break

	}
	return totalCost
}

func estimate(field Field, src, dst Pos, cost, lookAhead int) (p Pos, pCost int) {
	if lookAhead == 0 || src == dst {
		return src, cost
	}

	visited := field.Get(src.X, src.Y)
	if visited == -1 {
		return src, -1
	}

	field.Set(src.X, src.Y, -1)

	var bestPos Pos
	var bestPosCost int
	for _, pos := range field.adjacents(src.X, src.Y) {
		nextPos, nextCost := estimate(field, pos, dst, cost+visited, lookAhead-1)
		if nextCost == -1 {
			continue
		}

		if bestPosCost == 0 || bestPosCost > nextCost {
			bestPos = nextPos
			bestPosCost = nextCost
		}
	}

	field.Set(src.X, src.Y, visited)

	if bestPosCost == 0 {
		return src, -1
	}

	//fmt.Println("found", bestPos, bestPosCost+cost)

	return bestPos, bestPosCost + cost
}

type Pos struct {
	X, Y int
}

type Field [][]int

func (f Field) adjacents(x, y int) []Pos {
	var r []Pos
	/*if y > 0 { //never go back????
		r = append(r, Pos{x, y - 1})
	}*/

	if y < len(f)-1 {
		r = append(r, Pos{x, y + 1})
	}

	if x > 0 {
		r = append(r, Pos{x - 1, y})
	}

	if x < len(f[y])-1 {
		r = append(r, Pos{x + 1, y})
	}

	return r
}

func (f Field) Get(x, y int) int {
	return f[y][x]
}

func (f Field) Set(x, y, v int) {
	f[y][x] = v
}

func readField(fname string) (field Field, err error) {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	for _, s := range strings.Split(string(buf), "\n") {
		var row []int
		for _, v := range s {
			row = append(row, int(v)-'0')
		}
		field = append(field, row)
	}

	return field, nil
}
