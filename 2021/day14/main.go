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

	fmt.Println(apply(seed, 10, rules))
	fmt.Println(apply(seed, 40, rules))
}

func apply(seed string, iterations int, rules []Rule) int {
	counts := map[[2]byte]int{}
	for i := 0; i < len(seed)-1; i++ {
		k := [2]byte{seed[i], seed[i+1]}
		counts[k] = counts[k] + 1
	}

	for step := 0; step < iterations; step++ {
		tmp := map[[2]byte]int{}
		for bytes, c := range counts {
			for _, rule := range rules {
				if rule.pattern == bytes {
					left, right := [2]byte{rule.pattern[0], rule.char}, [2]byte{rule.char, rule.pattern[1]}
					cleft, cright := tmp[left], tmp[right]
					tmp[left], tmp[right] = c+cleft, c+cright
				}
			}
		}

		counts = tmp
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
