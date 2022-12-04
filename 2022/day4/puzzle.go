package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
)

//go:embed input.txt
var input []byte

type Interval struct {
	From, To int
}

func (i Interval) Contains(o Interval) bool {
	return o.From >= i.From && o.To <= i.To
}

func (i Interval) Overlaps(o Interval) bool {
	return (o.From >= i.From && o.From <= i.To) || (o.To >= i.From && o.To <= i.To)
}

type Pair struct {
	A, B Interval
}

func (p Pair) Contained() bool {
	return p.A.Contains(p.B) || p.B.Contains(p.A)
}

func (p Pair) Overlaped() bool {
	return p.A.Overlaps(p.B) || p.B.Overlaps(p.A)
}

func parsePairs() (pairs []Pair) {
	r := bytes.NewReader(input)
	for {
		var p Pair
		if _, err := fmt.Fscanf(r, "%d-%d,%d-%d\n", &p.A.From, &p.A.To, &p.B.From, &p.B.To); err == io.EOF {
			break
		}

		pairs = append(pairs, p)
	}

	return
}

func main() {
	contained, overlapped := 0, 0
	for _, pair := range parsePairs() {
		if pair.Contained() {
			contained++
		}
		if pair.Overlaped() {
			overlapped++
		}
	}
	fmt.Printf("part1: %d\npart2: %d\n", contained, overlapped)
}
