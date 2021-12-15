package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func main() {
	path := flag.String("file", "2021/day14/input.txt", "the commands")
	flag.Parse()

	seed, rules, err := parse(*path)
	if err != nil {
		log.Fatalf("cannot parse file: %v\n", err)
	}

	fmt.Println(part1(seed, 10, rules))
	fmt.Println("--")
	fmt.Println(part2(seed, 40, rules)) //test is correct but actual sample not

}

func part1(seed string, iterations int, rules []Rule) int {
	str := SliceBuf(seed)

	for step := 0; step < iterations; step++ {
		fmt.Println(step, len(str))
	nextPair:
		for i := 0; i < len(str)-1; i += 2 {
			for _, rule := range rules {
				if str[i] == rule.pattern[0] && str[i+1] == rule.pattern[1] {
					//str.join(str[:i], rule.char, str[i:])
					str.insert(i+1, rule.char)
					//i++
					//	fmt.Println("inserting", string(rule.char))
					continue nextPair
				}
			}

		}
		//fmt.Println(string(str))
	}

	min, max := str.count()
	return max - min
}

func part2(seed string, iterations int, rules []Rule) int {
	counts := map[[2]byte]int{}
	for i := 0; i < len(seed)-1; i++ {
		counts[[2]byte{seed[i], seed[i+1]}] = 1
	}

	for step := 0; step < iterations; step++ {
		tmp := map[[2]byte]int{}
		for bytes, c := range counts {
			for _, rule := range rules {
				if rule.pattern == bytes {
					left, right := [2]byte{rule.pattern[0], rule.char}, [2]byte{rule.char, rule.pattern[1]}
					cleft, cright := tmp[left], tmp[right]
					tmp[left], tmp[right] = c+cleft, c+cright
					fmt.Printf("%s => %s(%d) - %s(%d)\n", string(rule.pattern[:]), string(left[:]), c+cleft, string(right[:]), c+cright)
				}
			}
		}

		counts = tmp
	}

	for bytes, i := range counts {
		fmt.Println(string(bytes[:]), i)
	}

	flat := deflate(counts)
	flat[seed[0]] = flat[seed[0]] + 1
	flat[seed[len(seed)-1]] = flat[seed[len(seed)-1]] + 1
	min, max := minMax(flat)
	return (max - min) / 2
}

func parse(fname string) (seed string, rules []Rule, err error) {
	file, err := os.Open(fname)
	if err != nil {
		return "", nil, fmt.Errorf("cannot open file: %w", err)
	}

	defer file.Close()

	if _, err := fmt.Fscanf(file, "%s\n\n", &seed); err != nil {
		return "", nil, fmt.Errorf("cannot read seed: %w", err)
	}

	for {
		var src, dst string
		if _, err := fmt.Fscanf(file, "%s -> %s\n", &src, &dst); err != nil {
			if err == io.EOF {
				break
			}

			return "", nil, fmt.Errorf("cannot read rule: %w", err)
		}

		rules = append(rules, Rule{
			pattern: [2]byte{src[0], src[1]},
			char:    dst[0],
		})

	}

	return
}

type Rule struct {
	pattern [2]byte
	char    byte
}

func (r Rule) String() string {
	return string(r.pattern[:]) + " -> " + string(r.char)
}

type SliceBuf []byte

func (s *SliceBuf) join(a []byte, c byte, b []byte) {
	tmp := (*s)[:0]
	tmp = append(tmp, a...)
	tmp = append(tmp, c)
	tmp = append(tmp, b...)
	*s = tmp
}

func (s *SliceBuf) insert(pos int, v byte) {
	tmp := *s
	tmp = append(tmp, 0)
	copy(tmp[pos+1:], tmp[pos:])
	tmp[pos] = v
	*s = tmp
}

func (s *SliceBuf) count() (min, max int) {
	count := map[byte]int{}
	for _, b := range *s {
		count[b] = count[b] + 1
	}

	min, max = minMax(count)

	return
}

func minMax(count map[byte]int) (min, max int) {
	min, max = math.MaxInt, 0
	for _, i := range count {
		if i < min {
			min = i
		}
		if i > max {
			max = i
		}
	}

	return
}

func deflate(count map[[2]byte]int) map[byte]int {
	r := map[byte]int{}
	for bytes, i := range count {
		for _, b := range bytes {
			c := r[b]
			r[b] = c + i
		}
	}

	return r
}
