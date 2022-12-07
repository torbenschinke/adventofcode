package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Type int

const (
	CD Type = iota
	LS
	DIR
	FILE
)

type Instruction string

func (l Instruction) Kind() Type {
	switch l[2] {
	case 'c':
		return CD
	case 'l':
		return LS
	case 'r':
		return DIR
	default:
		return FILE
	}
}

func (l Instruction) ChangeDir(f *File) *File {
	name := l[len("$ cd "):]
	if name == ".." {
		return f.Parent
	}
	child := f.ByName(string(name))
	if child == nil {
		child = &File{
			Name: string(name),
			Dir:  true,
		}
		f.Files = append(f.Files, child)
	}

	return child
}

func (l Instruction) AppendDir(f *File) {
	name := l[len("dir "):]
	f.Files = append(f.Files, &File{
		Name:   string(name),
		Dir:    true,
		Parent: f,
	})
}

func (l Instruction) AppendFile(f *File) {
	c := &File{Parent: f}
	if _, err := fmt.Fscanf(bytes.NewBufferString(string(l)), "%d %s\n", &c.Size, &c.Name); err != nil {
		panic(err)
	}

	f.Files = append(f.Files, c)
}

type File struct {
	Name   string
	Size   int
	Dir    bool
	Files  []*File
	Parent *File
}

func (f *File) Total() int {
	sum := f.Size
	for _, file := range f.Files {
		sum += file.Total()
	}
	return sum
}

func (f *File) Visit(fun func(f *File)) {
	fun(f)
	for _, file := range f.Files {
		if file.Dir {
			file.Visit(fun)
		}
	}
}

func (f *File) ByName(n string) *File {
	for _, file := range f.Files {
		if file.Name == n {
			return file
		}
	}
	return nil
}

func ParseTree(root *File, lines []string) {
	if len(lines) == 0 {
		return
	}

	inst := Instruction(lines[0])
	switch inst.Kind() {
	case CD:
		ParseTree(inst.ChangeDir(root), lines[1:])
	case LS:
		ParseTree(root, lines[1:])
	case FILE:
		inst.AppendFile(root)
		ParseTree(root, lines[1:])
	case DIR:
		inst.AppendDir(root)
		ParseTree(root, lines[1:])
	}
}

func part1() {
	root := &File{Name: "/", Dir: true}
	ParseTree(root, strings.Split(input, "\n")[1:])

	sum := 0
	root.Visit(func(f *File) {
		if f.Total() <= 100000 {
			sum += f.Total()
		}
	})
	fmt.Printf("part1: %d\n", sum)
}

func part2() {
	root := &File{Name: "/", Dir: true}
	ParseTree(root, strings.Split(input, "\n")[1:])
	free := 70000000 - root.Total()
	requiredFree := 30000000

	var smallestDir *File
	root.Visit(func(f *File) {
		if f.Total()+free > requiredFree {
			if smallestDir == nil || smallestDir.Total() > f.Total() {
				smallestDir = f
			}
		}
	})
	fmt.Printf("part2: %d\n", smallestDir.Total())
}

func main() {
	part1()
	part2()
}
