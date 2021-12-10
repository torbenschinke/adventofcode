package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	path := flag.String("file", "2021/day9/input.txt", "the commands")
	flag.Parse()

	field, err := readField(*path)
	if err != nil {
		log.Fatalf("cannot read file: %v\n", err)
	}

	fmt.Println(part1(field))
	fmt.Println(part2(field))
}

func part1(field Field) int {
	risk := 0
	for y, row := range field {
		for x, val := range row {
			if field.IsMinima(x, y) {
				risk += val + 1
			}
		}
	}

	return risk
}

func part2(field Field) int {
	var lens []int
	for y, row := range field {
		for x, _ := range row {
			if field.IsMinima(x, y) {
				positions := field.FloodFill(x, y)
				lens = append(lens, len(positions))
			}
		}
	}

	sort.Ints(lens)
	l := len(lens)

	return lens[l-3] * lens[l-2] * lens[l-1]
}

type Pos struct {
	X, Y int
}

type Field [][]int

func (f Field) adjacents(x, y int) []Pos {
	var r []Pos
	if y > 0 {
		r = append(r, Pos{x, y - 1})
	}

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

func (f Field) FloodFill(x, y int) []Pos {
	var r []Pos
	if v := f.Get(x, y); v == -1 || v == 9 {
		return nil
	}

	r = append(r, Pos{x, y})
	f.Set(x, y, -1)

	for _, pos := range f.adjacents(x, y) {
		r = append(r, f.FloodFill(pos.X, pos.Y)...)
	}

	return r
}

func (f Field) Set(x, y, v int) {
	f[y][x] = v
}

func (f Field) Get(x, y int) int {
	return f[y][x]
}

func (f Field) IsMinima(x, y int) bool {
	v := f.Get(x, y)
	var vals []int
	for _, pos := range f.adjacents(x, y) {
		vals = append(vals, f.Get(pos.X, pos.Y))
	}
	
	return min(vals...) > v
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

func min(v ...int) int {
	x := v[0]
	for _, i := range v {
		if i < x {
			x = i
		}
	}
	return x
}
