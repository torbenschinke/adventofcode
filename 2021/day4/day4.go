package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	path := flag.String("file", "2021/day4/input.txt", "the commands")
	flag.Parse()

	buf, err := ioutil.ReadFile(*path)
	if err != nil {
		log.Fatalf("cannot read file: %v\n", err)
	}

	lines := strings.Split(string(buf), "\n")
	drawn, boards, err := parse(lines)
	if err != nil {
		log.Fatalf("cannot parse: %v", err)
	}

	score := winners(drawn, boards)
	fmt.Printf("part 1 score: %d\n", score[0].score(score[0].DrawnNums))
	fmt.Printf("part 2 score: %d\n", score[len(score)-1].score(score[len(score)-1].DrawnNums))
}

func parse(lines []string) (DrawnNums, []Board, error) {
	var drawn DrawnNums
	for _, s := range strings.Split(lines[0], ",") {
		num, err := strconv.Atoi(s)
		if err != nil {
			return nil, nil, fmt.Errorf("cannot parse drawn num: %w", err)
		}

		drawn = append(drawn, num)
	}

	var boards []Board
	for startOfBoard := 2; startOfBoard < len(lines); startOfBoard += 6 {
		var board Board
		for row := 0; row < 5; row++ {
			for col := 0; col < 5; col++ {
				num, err := strconv.Atoi(strings.TrimSpace(lines[startOfBoard+row][col*3 : col*3+2]))
				if err != nil {
					return nil, nil, fmt.Errorf("cannot parse board num in line %d: %w", startOfBoard+row, err)
				}

				board[row][col] = num
			}
		}
		boards = append(boards, board)
	}

	return drawn, boards, nil
}
func winners(nums DrawnNums, boards []Board) []Solution {
	var solutions []Solution
	alreadySolved := make(map[int]struct{})
	for i := range nums {
		drawn := nums[0:i]
		for bid, board := range boards {
			if _, ok := alreadySolved[bid]; !ok && board.solved(drawn) {
				solutions = append(solutions, Solution{
					DrawnNums: drawn,
					Board:     board,
				})
				alreadySolved[bid] = struct{}{}
			}
		}
	}

	return solutions
}

type DrawnNums []int

func (n DrawnNums) has(v int) bool {
	for _, i := range n {
		if i == v {
			return true
		}
	}

	return false
}

type Solution struct {
	DrawnNums
	Board
}

type Board [5][5]int

func (b Board) score(nums DrawnNums) int {
	score := 0
	for _, row := range b {
		for _, val := range row {
			if !nums.has(val) {
				score += val
			}
		}
	}

	return score * nums[len(nums)-1]
}

func (b Board) solved(nums DrawnNums) bool {
	return b.solvedRow(nums) >= 0 || b.solvedCol(nums) >= 0
}

func (b Board) solvedRow(nums DrawnNums) (row int) {
	for rI, row := range b {
		contained := 0
		for _, col := range row {
			if nums.has(col) {
				contained++
			}
		}
		if contained == len(row) {
			return rI
		}
	}

	return -1
}

func (b Board) solvedCol(nums DrawnNums) (col int) {
	for colI := 0; colI < 5; colI++ {
		contained := 0
		for _, row := range b {
			if nums.has(row[colI]) {
				contained++
			}
		}
		if contained == len(b) {
			return colI
		}
	}

	return -1
}
