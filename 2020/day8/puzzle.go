// Package day8 solves the according puzzle from https://adventofcode.com/2020/day/8.
package day8

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	errInfiniteLoop = fmt.Errorf("infinite loop")
	errSIGSEGV      = fmt.Errorf("SIGSEGV")
)

const (
	acc = iota + 1
	jmp
	nop
)

type Ins struct {
	Mark   byte
	Opcode byte
	Value  int16 // int32 would nearly double the padding within a slice
}

func ParseInstructions(str string) ([]Ins, error) {
	var res []Ins
	for _, line := range strings.Split(str, "\n") {
		value, err := strconv.Atoi(line[4:])
		if err != nil {
			return nil, fmt.Errorf("invalid opcode value: %w", err)
		}
		switch line[0] {
		case 'n':
			res = append(res, Ins{0, nop, int16(value)})
		case 'a':
			res = append(res, Ins{0, acc, int16(value)})
		case 'j':
			res = append(res, Ins{0, jmp, int16(value)})
		default:
			return nil, fmt.Errorf("invalid opcode: %s", line)
		}
	}

	return res, nil
}

type Interpreter struct {
	Ins []Ins
	PC  int
	Acc int
}

func (p *Interpreter) Execute() (int, error) {
	for {
		if p.PC < 0 || p.PC >= len(p.Ins) {
			return p.Acc, errSIGSEGV
		}

		ins := p.Ins[p.PC]
		if ins.Mark > 0 {
			return p.Acc, errInfiniteLoop
		}

		p.Ins[p.PC].Mark = 1
		switch ins.Opcode {
		case nop:
			p.PC++
		case jmp:
			p.PC += int(ins.Value)
		case acc:
			p.PC++
			p.Acc += int(ins.Value)
		default:
			panic("illegal state")
		}

		// termination condition
		if p.PC == len(p.Ins) {
			return p.Acc, nil
		}
	}
}

func (p *Interpreter) Reset() {
	p.PC = 0
	p.Acc = 0
	for i := range p.Ins {
		p.Ins[i].Mark = 0
	}
}

func (p *Interpreter) FixLoop() int {
	for i, in := range p.Ins {
		undoOpcode := in.Opcode
		switch in.Opcode {
		case jmp:
			p.Ins[i].Opcode = nop
		case nop:
			p.Ins[i].Opcode = jmp
		}

		p.Reset()
		res, err := p.Execute()
		p.Ins[i].Opcode = undoOpcode

		if err != nil {
			continue
		}

		return res
	}

	panic("no solution found: " + strconv.Itoa(p.Acc))
}
