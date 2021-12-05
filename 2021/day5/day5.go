package main

import (
	"flag"
	"fmt"
	"image"
	"io"
	"log"
	"os"
)

func main() {
	path := flag.String("file", "2021/day5/input.txt", "the commands")
	flag.Parse()

	lines, err := parse(*path)
	if err != nil {
		log.Fatalf("cannot parse file: %v\n", err)
	}

	fmt.Println(count(plot(lines, func(line Line) bool {
		return line.vertical() || line.horizontal()
	})))

	fmt.Println(count(plot(lines, func(Line) bool { return true })))
}

func count(img *image.Gray) int {
	hits := 0
	for x := 0; x < img.Stride; x++ {
		for y := 0; y < img.Stride; y++ {
			if img.GrayAt(x, y).Y > 1 {
				hits++
			}
		}
	}

	return hits
}

func plot(lines []Line, p func(Line) bool) *image.Gray {
	img := image.NewGray(image.Rect(0, 0, 1000, 1000))
	for _, line := range lines {
		if p(line) {
			line.Each(func(v Vec2) {
				c := img.GrayAt(v.x, v.y)
				c.Y++
				img.SetGray(v.x, v.y, c)
			})
		}
	}

	return img
}

func parse(fname string) (lines []Line, err error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}

	defer file.Close()

	for i := 0; ; i++ {
		var line Line
		_, err = fmt.Fscanf(file, "%d,%d -> %d,%d\n", &line.start.x, &line.start.y, &line.end.x, &line.end.y)
		if err == io.EOF {
			return lines, nil
		}

		if err != nil {
			return nil, fmt.Errorf("cannot parse line %d: %w", i, err)
		}

		lines = append(lines, line)
	}
}

type Vec2 struct {
	x, y int
}

type Line struct {
	start, end Vec2
}

func (l Line) vertical() bool {
	return l.start.x == l.end.x
}

func (l Line) horizontal() bool {
	return l.start.y == l.end.y
}

func (l Line) Each(f func(Vec2)) {
	dx := l.end.x - l.start.x
	dy := l.end.y - l.start.y
	n := max(abs(dx), abs(dy))

	x := l.start.x
	y := l.start.y
	for i := 0; i <= n; i++ {
		f(Vec2{x, y})
		x += dx / n
		y += dy / n
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
