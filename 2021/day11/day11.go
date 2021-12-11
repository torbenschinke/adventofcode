package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	path := flag.String("file", "2021/day11/input.txt", "the commands")
	flag.Parse()

	field, err := readField(*path)
	if err != nil {
		log.Fatalf("cannot read file: %v\n", err)
	}

	fmt.Println(part1(field))
}

func part1(field Field) int {
	count := 0
	print(field)
	fmt.Println()

	for i := 0; i < 100; i++ {
		fmt.Printf("step %d\n", i+1)
		inc(field)
		for flash(field) {
			count++
		}
		resetSaturated(field)
		print(field)
		fmt.Println()
	}

	return count
}

func flash(field Field) bool {
	for y, row := range field {
		for x, val := range row {
			if val > 9 {
				for _, pos := range field.adjacents(x, y) {
					adVal := field.Get(pos.X, pos.Y)
					if adVal == -1 {
						continue
					}
					field.Set(pos.X, pos.Y, adVal+1)
				}
				field.Set(x, y, -1)
				return true
			}
		}
	}

	return false
}

func inc(field Field) {
	for y, row := range field {
		for x, val := range row {
			field.Set(x, y, val+1)
		}
	}
}

func print(field Field) {
	for _, row := range field {
		for _, val := range row {
			fmt.Printf("%d", val)
		}
		fmt.Println()
	}
}

func resetSaturated(field Field) {
	for y, row := range field {
		for x, val := range row {
			if val == -1 {
				field.Set(x, y, 0)
			}
		}
	}
}

type Pos struct {
	X, Y int
}

type Field [][]int

func (f Field) adjacents(x, y int) []Pos {
	var r []Pos

	for iy := max(y-1, 0); iy <= min(y+1, 9); iy++ {
		for ix := max(x-1, 0); ix <= min(x+1, 9); ix++ {
			if iy != y || ix != x {
				r = append(r, Pos{ix, iy})
			}
		}
	}

	return r
}

func (f Field) Set(x, y, v int) {
	f[y][x] = v
}

func (f Field) Get(x, y int) int {
	return f[y][x]
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
