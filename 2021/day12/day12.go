package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

func main() {
	path := flag.String("file", "2021/day12/input.txt", "the commands")
	flag.Parse()

	idx, err := parseNodes(*path)
	if err != nil {
		log.Fatalf("cannot read file: %v\n", err)
	}

	fmt.Println(part1(idx["start"]))
}

func part1(start *Node) int {
	count := 0
	find(start, &count)
	return count
}

func find(p *Node, paths *int) bool {
	if !p.big() && p.visited > 0 {
		return false
	}

	if p.name == "end" {
		*paths++
		return true
	}

	p.visited++
	any := false
	for _, node := range p.next {
		if find(node, paths) {
			any = true
		}
	}

	p.visited--
	return any
}

type Node struct {
	name    string
	visited int
	next    []*Node
}

func (n *Node) String() string {
	return n.name
}

func (n *Node) big() bool {
	return unicode.IsUpper(rune(n.name[0]))
}

type NodeIndex map[string]*Node

func parseNodes(fname string) (idx NodeIndex, err error) {
	buf, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %w", err)
	}

	idx = NodeIndex{}
	for _, line := range strings.Split(string(buf), "\n") {
		tokens := strings.Split(line, "-")
		from, to := tokens[0], tokens[1]
		nF, nT := idx[from], idx[to]
		if nF == nil {
			nF = &Node{name: from}
			idx[from] = nF
		}

		if nT == nil {
			nT = &Node{name: to}
			idx[to] = nT
		}

		nF.next = append(nF.next, nT)
		nT.next = append(nT.next, nF)
	}

	return idx, nil
}
