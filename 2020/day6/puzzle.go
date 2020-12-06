// Package day6 solves the according puzzle from https://adventofcode.com/2020/day/6.
package day6

import (
	"math/bits"
	"strings"
)

// CountAnyone counts the amount of "yes" answers per group (separated by newline) and returns the total sum.
func CountAnyone(str string) (count int) {
	quests := uint32(0) // start with no bits set

	for _, line := range strings.Split(str, "\n") {
		if line == "" {
			count += bits.OnesCount32(quests)
			quests = 0 // reset quest flags to zeros

			continue
		}

		for _, quest := range line {
			// Normalize question identifier to a zero based index for the position.
			// Then shift 1 to the required position and set that bit to 1 in our variable.
			quests |= 1 << (quest - 'a')
		}
	}

	count += bits.OnesCount32(quests) // example does not contain final new line

	return
}

// CountEveryone counts the amount of "yes" answers per group (separated by newline) and returns
// the total sum, but only if it has been acknowledged by each member.
func CountEveryone(str string) (count int) {
	quests := uint32(0) // start with no bits set
	firstMember := true

	for _, line := range strings.Split(str, "\n") {
		if line == "" {
			count += bits.OnesCount32(quests)
			quests = 0 // reset quest flags to zeros
			firstMember = true

			continue
		}

		lineFlags := uint32(0)
		for _, quest := range line {
			// Normalize question identifier to a zero based index for the position, just as normal.
			// Then shift 1 to the required position and set that bit to 1 in our variable.
			lineFlags |= 1 << (quest - 'a')
		}

		// the first group member gives the pattern for all others.
		if firstMember {
			quests |= lineFlags
		} else {
			// afterwards keep only those bits, which are set for subsequent member comparisons.
			quests &= lineFlags
		}

		firstMember = false
	}

	count += bits.OnesCount32(quests) // example does not contain final new line

	return count
}
