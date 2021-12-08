package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	path := flag.String("file", "2021/day8/input.txt", "the commands")
	flag.Parse()

	entries, err := parse(*path)
	if err != nil {
		log.Fatalf("cannot parse file: %v\n", err)
	}

	fmt.Println(count(entries))
	fmt.Println(resolve(entries)) // TODO result is fine for testinput.txt but wrong for input.txt
}

func resolve(entries []Entry) int {
	sum := 0
	for _, entry := range entries {
		fmt.Println("->", entry.Decode())
		sum += entry.Decode()
	}
	return sum
}

func count(entries []Entry) int {
	count := 0
	what := []int{2, 3, 4, 7}
	for _, entry := range entries {
		for _, i := range what {
			for _, digit := range entry.Digits {
				if len(digit) == i {
					count++
				}
			}
		}
	}

	return count
}

func parse(fname string) (entries []Entry, err error) {
	file, err := os.Open(fname)
	if err != nil {
		return nil, fmt.Errorf("cannot open file: %w", err)
	}

	defer file.Close()

	for i := 0; ; i++ {
		var e Entry
		_, err = fmt.Fscanf(file, "%s %s %s %s %s %s %s %s %s %s | %s %s %s %s\n",
			&e.Patterns[0], &e.Patterns[1], &e.Patterns[2], &e.Patterns[3], &e.Patterns[4], &e.Patterns[5], &e.Patterns[6], &e.Patterns[7], &e.Patterns[8], &e.Patterns[9],
			&e.Digits[0], &e.Digits[1], &e.Digits[2], &e.Digits[3])

		e.Sort()

		if err == io.EOF {
			return entries, nil
		}

		if err != nil {
			return nil, fmt.Errorf("cannot parse line %d: %w", i, err)
		}

		entries = append(entries, e)
	}
}

type Entry struct {
	Patterns [10]Pattern
	Digits   [4]Pattern
}

func (e *Entry) Sort() {
	for i, pattern := range e.Patterns {
		e.Patterns[i] = pattern.Sort()
	}

	for i, digit := range e.Digits {
		e.Digits[i] = digit.Sort()
	}
}

func (e Entry) Deduce() [10]Pattern {
	_1 := e.pickByLen(2)[0]
	_4 := e.pickByLen(4)[0]
	_7 := e.pickByLen(3)[0]
	_8 := e.pickByLen(7)[0]

	_6_or_9_or_0 := e.pickByLen(6)
	var _0_or_9 []Pattern
	var _6 Pattern
	for _, p := range _6_or_9_or_0 {
		if len(_1.Minus(p)) != 0 {
			_6 = p
		} else {
			_0_or_9 = append(_0_or_9, p)
		}
	}

	var _9, _0 Pattern
	if len(_0_or_9[0].Minus(_4)) == 2 {
		_9 = _0_or_9[0]
		_0 = _0_or_9[1]
	} else {
		_9 = _0_or_9[1]
		_0 = _0_or_9[0]
	}

	_9 = _9.Sort()
	_0 = _0.Sort()

	_a := _7.Minus(_1).Sort()
	_c := _9.Minus(_6).Sort()
	_f := _1.Minus(_c).Sort()
	_g := _9.Minus(_4).Minus(_a).Sort()
	_d := _4.Minus(_0).Sort()

	_3 := _a.Plus(_c).Plus(_f).Plus(_d).Plus(_g).Sort()

	var _5_or_2 []Pattern
	for _, pattern := range e.Patterns {
		switch pattern {
		case _0:
		case _1:
		case _3:
		case _4:
		case _6:
		case _7:
		case _8:
		case _9:
		default:
			_5_or_2 = append(_5_or_2, pattern)
		}
	}

	var _5, _2 Pattern
	if _5_or_2[0].Minus(_0).Has(_c.First()) {
		_2 = _5_or_2[0]
		_5 = _5_or_2[1]
	} else {
		_2 = _5_or_2[1]
		_5 = _5_or_2[0]
	}

	return [10]Pattern{
		_0.Sort(), _1.Sort(), _2.Sort(), _3.Sort(), _4.Sort(), _5.Sort(), _6.Sort(), _7.Sort(), _8.Sort(), _9.Sort(),
	}
}

func (e Entry) Decode() int {
	decoded := e.Deduce()
	tmp := ""
	for _, digit := range e.Digits {
		digit = digit.Sort()
		for num, pattern := range decoded {
			if pattern == digit {
				tmp += strconv.Itoa(num)
			}
		}
	}

	i, _ := strconv.Atoi(tmp)
	return i
}

func (e Entry) pickByLen(l int) []Pattern {
	var r []Pattern
	for _, pattern := range e.Patterns {
		if len(pattern) == l {
			r = append(r, pattern)
		}
	}

	return r
}

type Pattern string

func (p Pattern) Minus(other Pattern) Pattern {
	var n string
	for _, b := range p {
		if !other.Has(byte(b)) {
			n += string(b)
		}
	}

	return Pattern(n)
}

func (p Pattern) Plus(other Pattern) Pattern {
	n := string(p)
	for _, b := range other {
		if !p.Has(byte(b)) {
			n += string(b)
		}
	}

	return Pattern(n)
}

func (p Pattern) Has(c byte) bool {
	for _, b := range p {
		if byte(b) == c {
			return true
		}
	}

	return false
}

func (p Pattern) First() byte {
	return p[0]
}

func (p Pattern) Sort() Pattern {
	tmp := make([]int, 0, len(p))
	for _, r := range p {
		tmp = append(tmp, int(r))
	}
	sort.Ints(tmp)

	sb := strings.Builder{}
	for _, i := range tmp {
		sb.WriteRune(rune(i))
	}

	return Pattern(sb.String())
}
