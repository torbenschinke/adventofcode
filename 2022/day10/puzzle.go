package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type CPU struct {
	RegX       int
	Program    []Instruction
	PC         int
	CycleCount int
}

func (c *CPU) Done() bool { return c.PC >= len(c.Program) }

func (c *CPU) SignalStrength() int { return c.CycleCount * c.RegX }

func (c *CPU) Cycle() bool {
	if c.Done() {
		return false
	}

	ins := c.Program[c.PC]
	if !ins.Exec(c) {
		c.PC++
	}
	c.CycleCount++
	return c.Done()
}

type Instruction interface {
	Exec(cpu *CPU) (hasNext bool)
}

type NoOp struct{}

func (NoOp) Exec(cpu *CPU) bool { return false }

type AddX struct {
	X  int
	CC int
}

func (ins *AddX) Exec(cpu *CPU) bool {
	ins.CC++
	if ins.CC <= 1 {
		return true
	}
	cpu.RegX += ins.X
	return false
}

type CRT [6][40]byte

func (c *CRT) Cycle(cpu *CPU) {
	row := ((cpu.CycleCount - 1) / 40) % 6
	col := (cpu.CycleCount - 1) % 40
	pix := byte('.')
	if col >= cpu.RegX-1 && col <= cpu.RegX+1 {
		pix = '#'
	}

	c[row][col] = pix
}
func (c *CRT) Print() {
	for row := range c {
		for col := range c[row] {
			fmt.Printf("%c", c[row][col])
		}
		fmt.Printf("\n")
	}
}

func parseProgram(input string) []Instruction {
	var res []Instruction
	for _, line := range strings.Split(input, "\n") {
		switch line {
		case "noop":
			res = append(res, NoOp{})
		default:
			x, err := strconv.Atoi(line[len("addx "):])
			if err != nil {
				panic(err)
			}
			res = append(res, &AddX{X: x})
		}
	}

	return res
}

func main() {
	crt, cpu := CRT{}, CPU{RegX: 1, Program: parseProgram(input), CycleCount: 1}
	sum := 0
	for !cpu.Done() {
		crt.Cycle(&cpu)
		cpu.Cycle()
		if cc := cpu.CycleCount; cc == 20 || cc == 60 || cc == 100 || cc == 140 || cc == 180 || cc == 220 {
			sum += cpu.SignalStrength()
		}
	}

	fmt.Printf("part 1: %d\npart2:\n", sum)
	crt.Print() // still hard to read: PLPAFBCL
}
