package main

import "fmt"

var (
	seedTest = []int{3, 4, 3, 1, 2}
	seed     = []int{3, 5, 4, 1, 2, 1, 5, 5, 1, 1, 1, 1, 4, 1, 4, 5, 4, 5, 1, 3, 1, 1, 1, 4, 1, 1, 3, 1, 1, 5, 3, 1, 1, 3, 1, 3, 1, 1, 1, 4, 1, 2, 5, 3, 1, 4, 2, 3, 1, 1, 2, 1, 1, 1, 4, 1, 1, 1, 1, 2, 1, 1, 1, 3, 1, 1, 4, 1, 4, 1, 5, 1, 4, 2, 1, 1, 5, 4, 4, 4, 1, 4, 1, 1, 1, 1, 3, 1, 5, 1, 4, 5, 3, 1, 4, 1, 5, 2, 2, 5, 1, 3, 2, 2, 5, 4, 2, 3, 4, 1, 2, 1, 1, 2, 1, 1, 5, 4, 1, 1, 1, 1, 3, 1, 5, 4, 1, 5, 1, 1, 4, 3, 4, 3, 1, 5, 1, 1, 2, 1, 1, 5, 3, 1, 1, 1, 1, 1, 5, 1, 1, 1, 1, 1, 1, 1, 2, 2, 5, 5, 1, 2, 1, 2, 1, 1, 5, 1, 3, 1, 5, 2, 1, 4, 1, 5, 3, 1, 1, 1, 2, 1, 3, 1, 4, 4, 1, 1, 5, 1, 1, 4, 1, 4, 2, 3, 5, 2, 5, 1, 3, 1, 2, 1, 4, 1, 1, 1, 1, 2, 1, 4, 1, 3, 4, 1, 1, 1, 1, 1, 1, 1, 2, 1, 5, 1, 1, 1, 1, 2, 3, 1, 1, 2, 3, 1, 1, 3, 1, 1, 3, 1, 3, 1, 3, 3, 1, 1, 2, 1, 3, 2, 3, 1, 1, 3, 5, 1, 1, 5, 5, 1, 2, 1, 2, 2, 1, 1, 1, 5, 3, 1, 1, 3, 5, 1, 3, 1, 5, 3, 4, 2, 3, 2, 1, 3, 1, 1, 3, 4, 2, 1, 1, 3, 1, 1, 1, 1, 1, 1}
)

func main() {
	fmt.Println(simulatePopulation(80, seed))
	fmt.Println(simulatePopulation(256, seed))
}

type Fishes [9]int

func (f *Fishes) Evolve() {
	oldGen, newGen := f[0], f[0]
	copy(f[:], f[1:])
	f[8] = newGen
	f[6] += oldGen
}

func simulatePopulation(days int, seed []int) int {
	var fishes Fishes
	for _, i := range seed {
		fishes[i]++
	}

	for i := 0; i < days; i++ {
		fishes.Evolve()
	}

	c := 0
	for _, i := range fishes {
		c += i
	}

	return c
}
