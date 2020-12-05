// Package day2 solves https://adventofcode.com/2020/day/2.
package day2

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	errMissDash      = fmt.Errorf("row misses dash")
	errMissLetterSep = fmt.Errorf("row missing letter separator")
	errTruncated     = fmt.Errorf("rule is truncated")
	errMissPassword  = fmt.Errorf("not a valid 'rule: password' format")
)

// Row represents a password rule including the password itself.
type Row struct {
	DirectiveA, DirectiveB int    // DirectiveA and DirectiveB are either inclusive Min/Max values or exclusive Index
	Letter                 rune   // Letter which must occur at least Min and at most Max time.
	Password               string // Password represents the actual password
}

// ParseRow tries to parse the given text as a rule. Letter must be an ASCII char.
func ParseRow(str string) (Row, error) {
	var r Row

	numSep := strings.IndexByte(str, '-')
	if numSep == -1 {
		return r, errMissDash
	}

	letSep := strings.IndexByte(str, ' ')
	if letSep == -1 {
		return r, errMissLetterSep
	}

	a, err := strconv.Atoi(str[0:numSep])
	if err != nil {
		return r, fmt.Errorf("unable to parse min value: %w", err)
	}

	b, err := strconv.Atoi(str[numSep+1 : letSep])
	if err != nil {
		return r, fmt.Errorf("unable to parse max value: %w", err)
	}

	if letSep+1 >= len(str) {
		return r, errTruncated
	}

	pwdSep := strings.LastIndexByte(str, ' ')
	if pwdSep == -1 {
		return r, errMissPassword
	}

	r.DirectiveA = a
	r.DirectiveB = b
	r.Letter = rune(str[letSep+1]) // this is not generally correct, because it truncates codepoints
	r.Password = str[pwdSep+1:]

	return r, nil
}

// CanValidate1 returns  true if this rule validates the according Password where DirectiveA and DirectiveB represents
// inclusive Min/Max occurrences. This only works correctly is Row.Letter is ASCII.
func CanValidate1(row Row) bool {
	count := 0

	for _, rn := range row.Password {
		if rn == row.Letter {
			count++
		}
	}

	return count >= row.DirectiveA && count <= row.DirectiveB
}

// CanValidate2 returns if this rule validates the according Password where DirectiveA and DirectiveB
// the indices where either of both must be the letter. This only works correctly if Row.Letter is ASCII.
func CanValidate2(row Row) bool {
	a := row.DirectiveA - 1
	b := row.DirectiveB - 1

	if a < 0 || b < 0 || a >= len(row.Password) || b >= len(row.Password) {
		return false
	}

	aIsLetter := row.Password[a] == byte(row.Letter)
	bIsLetter := row.Password[b] == byte(row.Letter)

	return aIsLetter != bIsLetter
}

// CountValid reads 'rule: password' strings like '1-3 a: abcde' separated by newlines and applies
// the closure on it.
func CountValid(lines string, validates func(row Row) bool) (int, error) {
	count := 0

	for _, line := range strings.Split(lines, "\n") {
		row, err := ParseRow(line)
		if err != nil {
			return -1, fmt.Errorf("unable to parse row: %w", err)
		}

		if validates(row) {
			count++
		}
	}

	return count, nil
}
