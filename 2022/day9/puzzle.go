package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"math"
)

//go:embed input.txt
var input []byte

type Motion struct {
	Dir   byte
	Count int
}

func parseMotions(input []byte) []Motion {
	r := bytes.NewReader(input)
	var motions []Motion
	for {
		var m Motion
		if _, err := fmt.Fscanf(r, "%c %d\n", &m.Dir, &m.Count); err == io.EOF {
			break
		}

		motions = append(motions, m)
	}
	return motions
}

type Vec struct{ X, Y int }

func (v Vec) Distance(o Vec) int {
	return int(math.Sqrt(float64((v.X-o.X)*(v.X-o.X) + (v.Y-o.Y)*(v.Y-o.Y))))
}

func (v Vec) Sub(o Vec) Vec { return Vec{v.X - o.X, v.Y - o.Y} }

func (v Vec) Add(o Vec) Vec { return Vec{v.X + o.X, v.Y + o.Y} }

type Node struct {
	Tracker map[Vec]int
	Child   *Node
	Vec
}

func NewRope(length int) *Node {
	n := &Node{}

	if length > 1 {
		n.Child = NewRope(length - 1)
	} else {
		n.Tracker = map[Vec]int{}
	}

	return n
}

func (n *Node) Tail() *Node {
	root := n
	for {
		if root.Child == nil {
			return root
		}
		root = root.Child
	}
}

func (n *Node) Pos(dir byte) Vec {
	v := n.Vec
	switch dir {
	case 'L':
		v.X--
	case 'R':
		v.X++
	case 'U':
		v.Y--
	case 'D':
		v.Y++
	}

	return v
}

func (n *Node) Move(v Vec) {
	n.Vec = v
	if n.Tracker != nil {
		n.Tracker[v] = 1
	}

	if n.Child == nil {
		return
	}

	delta := n.Vec.Distance(n.Child.Vec)
	if delta > 1 {
		dv := v.Sub(n.Child.Vec)
		if math.Abs(float64(dv.X)) > 1 {
			dv.X /= 2
		}
		if math.Abs(float64(dv.Y)) > 1 {
			dv.Y /= 2
		}
		n.Child.Move(n.Child.Vec.Add(dv))
	}
}

func trace(length int) int {
	head := NewRope(length)
	for _, motion := range parseMotions(input) {
		for i := 0; i < motion.Count; i++ {
			p := head.Pos(motion.Dir)
			head.Move(p)
		}
	}

	return len(head.Tail().Tracker) + 1
}

func main() {
	fmt.Printf("part 1: %d\n", trace(2))
	fmt.Printf("part 2: %d\n", trace(10))
}
