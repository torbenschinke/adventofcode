package day5_test

import (
	"fmt"
	"testing"

	"github.com/torbenschinke/adventofcode/2020/day5"
)

const (
	//nolint:lll
	set1 = "BBBFBFFRLR\nBFFFFFFRLL\nBFBBBBFLLL\nBBFBFFFRRR\nBBBFBFBRLR\nFBFFBBBLRL\nFBFFBFFRRL\nFFBBFFFRRL\nBBBFBFBLLL\nBBBBFFFLLR\nFBBFFBFRRR\nBFFFBBBLRR\nFFBBFFBRLR\nFFBFFBFRRL\nBBBBFBFRLL\nBBBBBFFLRL\nFFBBBFBRRR\nFFFBBBFLLR\nFFBFFBBLRL\nFBBFBBFLRL\nFFBBBBFRRR\nBFBFBFFRLL\nBBFBFFFLLL\nFFFBBBFRRL\nBBFFBFBRLL\nFBFBFBFLLL\nBBFFFFBLRR\nFFBBBFFRRL\nFFFBFBBLRL\nFBBFBBBLLL\nBBFBFFBRLR\nFBFBFBBRLL\nFFFBBFBLRL\nFFBFFBFLRR\nBBFBBBFRRL\nBBFBBFBLLR\nBBFBBBFRLL\nBFFBBBBRLR\nFBFBFBBRLR\nFBBBFFBRRR\nFBBBFBFLLR\nFFBFBBFLRL\nBFBFBFBRLL\nBFFFBFBLRR\nFFBBBFBRLR\nBBFBFFBLRR\nBFFFFBBLLL\nBFFFFBFLLR\nFBFBBBBRLR\nBBFFBBFLRL\nBFBFFBBRLR\nFFFBFBFRLL\nBFFFBBFRRL\nBBFBFBFRLL\nBFBFFFFRRL\nFFBFBBBRLR\nFBBFFFFLRR\nFBFBBFBRLR\nBBFFBFFLRR\nBBBFFBFLRR\nBFFFFFBRLR\nFBBFFFBRRL\nBFFFBBFRRR\nBBBFBFBRLL\nBFBBFBBRLR\nBFBFBFBLLL\nBBBFFFFRLL\nBFFFFBFRLL\nFFBBFFBLLL\nFBFFBBFRRL\nFFBBFBFRLR\nFFBFFFFRLR\nBFBBBFBLLL\nFBBFFBFRRL\nFBFBFFFLRR\nBFFBBBFLRR\nFBBFBBBRLL\nFFBFBBFLLR\nFFFBFFBRLR\nBBFBBFFRRR\nBFBBBFFLLL\nBFBBFBBRRR\nBBFFFBBLLR\nBBFBFFBLLR\nBFFBBFFRRL\nFFFBBFBRRR\nBBFBBFFLLR\nBFFBBBBRLL\nBFFBBFFRLR\nFFFBFBFLRL\nFBFFBBFLLR\nBFBBFBFRRL\nBFBFFFBRLL\nFBBFBFFLRR\nFFBFBBFLLL\nBFBBFBBRRL\nFBBBFFFLRR\nBBBFBBFRRL\nBBFBBBFRRR\nBFFBFFFLRL\nBFBFBBFRRL\nFFBBFFFLLR\nFBFFFFBLLL\nBFBBBBFRLR\nFFBFBFBRLL\nFBFFFBFLRL\nFFBFBFFLLL\nFBBBFFBRLL\nFBBBBBFLLR\nFFBFBBBLLL\nBFFBBBBLLL\nBFFBBBFRLL\nFBFFFFBLRL\nBFBFFBFRRR\nFFBFFFFRLL\nBBFBBFFRRL\nBBFBFBBLLL\nFFBBBBBRRR\nBBBFFFBLRL\nFBFBBBFRLL\nBFBFFFBRRL\nFFBFBBBRLL\nFBBBFBFLLL\nFBFBFFBLRL\nFBBBFFFLRL\nFBBBFBBRRL\nFBBFFBBRRR\nFFBFFBBRRR\nFBFFFFFLLL\nFFBBBFFRLL\nFBBFFFBRRR\nBFBBBFBRRR\nFFBFFBBRRL\nFFFBFBFLLR\nFBBFBFBRLL\nFBFFBFFRLR\nBFBBFFBLRR\nFFBFBBBLRR\nFFBFBBFLRR\nBFFFBFFLRL\nBBFFFBBRLL\nFFFBBFFLRL\nBBFBBFFLLL\nBFFBFFBRLR\nBBBFFBFRRR\nFBFFBFBRLL\nFBBBFFBLLR\nFFFBBFFLRR\nFFFBBFBLLL\nFBFBFBBLLR\nFBBBFBFRLR\nBFBBBFBLRL\nFBFBFFBLLL\nFBFBBFBRRR\nBFFBBBBRRR\nBFFFBFBLLR\nFFFBFBBRRL\nBBBBFBBLLL\nFBFFBFBLRR\nBFFFFFFLLL\nFFBFBFBRLR\nFBBBFFFLLR\nBBBFFFBLRR\nFBFFFFBRRL\nFBFFFFBRLR\nFBFFBBBLRR\nBBFBFBFRRL\nBFFFBFBRLL\nFBFBBFFRRR\nBFFFFFFRLR\nBFBBBBBLLR\nFBBBBFBLRR\nFBBFBBFRLR\nBBFBBBFLRR\nFBFFFBBLRR\nBBBBFFFRLL\nBFBFFBBRRR\nFBFFFBBLLR\nFFBBBFFRLR\nFBBFFFBRLL\nFBBFFBBRRL\nBFFFFFFLLR\nBBFBFBBRRL\nFBFFFBFRRL\nBBBBFFFLRL\nFFBFFBBLRR\nFBFFFFFRLL\nBFFFFBFRRR\nFFBFFFBRRR\nFBFBBFBRRL\nFFBFBFBLRR\nBBFBFFFRLR\nBBFFBBBRRR\nBBFFFBFLRL\nFBBBBBFLRL\nFBBBFFBRLR\nFBBFBBBLLR\nFFBFBBFRLR\nBFFBFBBLRL\nBBBFBFBLRL\nBFFBBFBRLR\nBFFFFFBRRL\nFBBFFBBLRR\nBFBBBBFRRL\nBBBBFFFRLR\nBFFFFFBLLR\nBBBBFBFRRR\nBFFBFBBRLR\nBFBFBFFLRL\nBBFFBBBRLR\nBBFFFBFLLR\nFFBFFBFLLL\nBBFFFFFRRR\nBBFBFFFLLR\nBFBFFFFLRR\nFBFBFBBRRL\nFBFFFFFRLR\nBBBBFFBLLL\nBBFFFBBLRR\nFFBBBBBRLL\nFBFFFFFLRR\nBFFFBBFLRL\nFBBBBFBLLL\nFBFBBFFLLL\nBBFFFBBRLR\nFFFBFFBRRR\nBBFFBBBRLL\nBFBFFBBLLR\nFFFBBBBRRL\nBFFFFBFLLL\nFBBFBBBRRR\nBFBBFBBLLL\nBFBFBBFLRL\nFBFFBFFRRR\nFBBFFFFRLL\nFBBBFBFRRR\nBBBFBFFLRR\nBBFFBBBLRR\nFFBBFBBRRR\nBFBBBFBLLR\nBBFBFFBRLL\nFFBBFBBRLL\nFFBBBFFLRL\nBBFBBBBRRR\nFFFBFFFLRR\nBFBBBFBRLL\nFFBBFBBLLR\nBFBFBFFLRR\nBBBBFBBRRL\nFBBFFBFLLR\nBFBBBFFRLL\nBBBFFBBLLL\nBFBFFBBLRR\nFBFFBBBRLR\nFBBFBFBRRL\nFBFBBBFRRR\nFBFFFBBRLR\nBBFFFFFLRR\nBBBBFFBRLL\nBFFBFFFRRR\nFFBFBFFLRR\nFFFBFBFRLR\nBFFFBBFRLL\nFFBFBBBRRR\nBFFFFBBRLR\nBBBFFFFRRL\nFBFBFBFRRR\nFBBBBFBRLR\nFBFBBBFRLR\nBBBBFFFLRR\nBFFBFFFRLR\nBFFFBBBLLL\nBFBBFFBRLR\nBFBFFBFLLL\nBFBFBBFRLL\nBFFBBFBLLL\nBFBBFFFRLL\nFFBBFFBLRL\nBBBBFFBRRL\nFBBBFFBLRR\nFFBBBBBLLR\nFBFFFBFLLR\nBBFFBFFRRL\nBFBFBBBLLL\nFBBBBFBRRL\nFBFBBBBRRR\nBBFBBFBRRR\nFBFBFFFRLL\nBFBFFBFLRL\nBFBFFBBLRL\nFFBFBFFRLR\nBFBBFBBRLL\nFBBBBBBRLR\nBBFBBBBLRL\nFFBBFBFLLL\nFFBBFBBLLL\nFFFBBFBRLR\nBBBBFBFLLL\nFBFBFFBLRR\nFFBBFFFLRL\nBBFBFFFRLL\nFBFFFBFRLL\nBBBFBBFLLL\nFFBBFFBLRR\nBBBBFBFLRR\nFBFFFFFRRL\nBFBFFBFRLR\nFBFBFFBRLR\nFBBFBBFRLL\nBFFFFFFLRR\nBBBFFFFLLL\nFFFBBFFRLR\nFFBBFFFRLR\nBBBFFBBRRL\nFBBFBFBLRL\nBFFFBBBLRL\nFBBFBBBRLR\nBFBFFFBLLL\nFBFFFBBLRL\nFBFFBFFRLL\nBFBFBFBLLR\nBBFFBFBRLR\nBBFFFFBRLR\nFFFFBBBRRR\nFFBBBBBLLL\nBBFBBBBRRL\nBBBBBFFLLR\nFFFBFFFRLL\nBFFBBFFLRR\nBFFBBFFLRL\nBBBFFBFRLL\nFFBFBFBLLL\nBFBBBBBLLL\nFFBBBBFLRL\nFBFFBBFLRL\nBBFFFFBLRL\nBBBFBBBRLL\nBFBFFFBLLR\nBFBFBFBRLR\nBFBBFFFRRR\nBFFBFBFLLL\nFBBFBFBLLR\nFBFBFBFLRR\nFFBFFBFLLR\nFFBFFFBLLL\nFBFFBFFLRL\nFFBBFBFLLR\nBFFFBFFLRR\nBBFFFBFRLL\nFBFBBBFLRR\nFFBBFFFLRR\nFBBBFBBLLL\nBFFBFBFRLR\nFFBFBBBLLR\nBBBFFBFLLR\nBFBFFFBLRL\nBFFBBFFLLL\nFBFBFFFRRL\nBFFBBBBLLR\nFFBBBFFLLR\nBBFFBBFRRL\nFFBFFBFLRL\nBBFFBFBRRL\nBFBBBFFLRR\nBFFBFBBRRR\nFBBFBBBLRL\nFFBFBFFRRR\nBBFBFFFLRL\nBBBFFBBLRR\nFBBFBFBLRR\nBFBBBBFLRL\nFBFFFBBRRL\nBFFFFBFLRR\nBBFFBBFRLL\nBFFBFBBLRR\nFBBBBFFRLL\nBFBFFFFLRL\nFBBBBBBLLL\nBFBBBBFRLL\nBFBFFFFLLL\nFFBBBFBLRL\nBBFFFFBLLL\nBFBBFBFLRR\nFFFFBBBRRL\nBFFBFBBLLL\nFFBBFFFRLL\nBFFFBFBRRR\nFBBBFBBRRR\nBBFBBFBLRR\nBFFBFFBLRL\nBBFFFBBLRL\nBFFBBBFRRR\nFFBFFBBRLR\nBFFBFBFLLR\nBFFBBFFRLL\nFBBFBFBRLR\nBFFFFFBLRR\nBFBFFBFRRL\nBFBBFFBLLL\nFBFFFFBRRR\nBBFFBFFLLL\nFBBBBBBLLR\nBFFFFFBRRR\nBFBFBFFRRR\nBBFFBFFRLR\nFBBBFFBLLL\nBFBFBBFLLL\nBFBFBBFRLR\nBBFFBFFLLR\nFBBFFFBLRR\nFBFFBBFLRR\nBFFFFBBLLR\nFFBFBFBLRL\nFFFBBBFRLL\nBFFBBFBRRR\nBFFBBBFRLR\nBBBBFFFRRR\nFBBBBFBLRL\nFBFFFFBLRR\nBFFFBBBRLL\nFBFBBFFRLR\nFFBBBBFLLR\nFFBFFBFRLR\nFFBBBBBRRL\nFBFFFFFLLR\nBFBBBFFLLR\nFBFFBBFRLL\nBBFFFBFRRL\nBBFFBFFLRL\nFBFBFFFLLL\nBBFFBFFRLL\nFFFBFFBLLL\nBBFBBFBLRL\nBBFFBFBLRL\nBFBFBFBRRL\nFBBBFBBLRR\nFBBBFBFLRL\nFBBBBFFRLR\nFFBFBFFLLR\nBFFFBFBRLR\nFBBFFBBRLR\nBBFFBBFRRR\nBBBFFBBRRR\nFFBFFFBRLR\nBBBBFFBLRR\nBBFBFBFLLL\nBBFFFBBLLL\nBFBBFFBLRL\nFBFBFBFRRL\nBBFBBBFLLR\nFBBBBBFRRL\nBFBFBFBLRR\nFBFBFFFRRR\nBFBFFBFLLR\nBBBFFBFRRL\nBBBFBFFLLL\nBFFFBFFRRR\nFFFBBBBLRR\nBFFBBFBLRR\nBFFBFBFLRL\nBBBBFFBLRL\nFBBFBBFRRR\nFFFBBBBLLL\nFBFFFBBRRR\nFBBBFFFLLL\nFFBBBBBRLR\nBFBFFFFRRR\nFBBFFFFRRR\nBBBFBFFLRL\nBFFFBFFRLR\nBFBFFFBLRR\nBFBBFBFRLR\nBBFFFFFRLR\nBFFFBBFLLL\nBBBFFFFLRL\nBBFFBBBLLL\nFBFFBFBRRL\nBFBBBBBRLR\nFBBFFFBLLL\nFBBBFBFLRR\nBBFBBBFLLL\nFFBFFBBLLR\nBFFFBBFLRR\nBBBBFFBRRR\nBBFFFFBRRL\nFBFFFBFLLL\nBBFFBBFLLL\nBFFFFFFRRL\nFBBFFBBRLL\nFFBBFBFRRL\nBBBFFBBRLL\nBFBFBBFRRR\nFFBFFFBLRL\nBFFFFBBRRR\nFFFBBFBRRL\nBBBFFBFLRL\nBFFBBFBLLR\nFFFBFFBLRL\nFBBBFBBRLL\nBFBBFBBLLR\nFBBFFFFLLR\nFFBFBFBRRL\nFBBFFBBLLL\nBBBBFFFLLL\nFBBBBBBLRR\nFFFBBBBRLR\nFBBBBFFRRL\nBBFFBBBRRL\nFBBFFFFRLR\nFFBFBFFLRL\nFBFFFBFRRR\nFBBFBBFLLR\nFBBBBBFLLL\nFFBFBFFRLL\nFFBFFFFLRR\nFFBBBBFLRR\nFFFBFBBLLR\nBBBFBFBLRR\nFFBBFBBRLR\nFBFFFFBLLR\nBBFBFBBRLR\nFBBFFFFLLL\nBBBBFBBLLR\nBFBBBBFRRR\nBBBFBFFRRR\nBFFBBBBLRL\nBBBFBBBRRR\nFBFFBFFLLL\nFFBFBFFRRL\nBBBBFFBRLR\nBFBBFFFRRL\nFFBFFFBRLL\nBBBBFBFLRL\nFBBBFBBLRL\nFFBBBFFLLL\nFFFBFFBLLR\nBFBBFBBLRR\nFFBFBBBLRL\nBFBBFFFLRR\nFFFBFFFLLR\nBFBBBBFLLR\nBFFBFBFRLL\nBFFBFBBRLL\nBBFBFFFLRR\nBBBFBBFLLR\nFBFBFFFLRL\nFBFFBFBLLL\nFFFBFFBLRR\nBFFFBFFRLL\nBBFFBBFLLR\nBBFFBBBLRL\nFBFBBBBLLL\nFBBFBBBLRR\nFBBBBBFLRR\nFBBBBBBLRL\nFBFBBFBLLL\nFFFBFFBRLL\nFBFBBFBLRR\nFBBFFBFRLR\nBBFFFFFRLL\nFFBBFBBRRL\nBFBFBBBRLL\nBBBFFFFLLR\nFBFFFBFRLR\nBFFFBFFLLL\nBBBFFFBRLL\nFBBBBFBLLR\nFBBFBFBLLL\nBBFBFBFLRR\nBFBFBFBLRL\nFBBBBBBRRL\nBFBFFFFRLL\nBBBFBBBLLL\nBFBFFFBRLR\nFFFBBBBRRR\nFFBFFFBRRL\nFBFBBBBLRL\nBBFBBFBRRL\nBBBFFBBLRL\nFBFBFBBLLL\nFBBBBBBRRR\nFFFBFBBRLR\nFBFFBFBRRR\nFFBBFBBLRL\nBBFBBBBLRR\nBFFFFFFRRR\nBBFFBFBLRR\nFFFBBBFLRR\nBBFFBBFLRR\nBFBFBFFRLR\nFBFBFFBRLL\nFFBFFFFLLL\nFBFBFFBLLR\nBFFFFBBLRR\nFBBBFFFRRR\nFFFBFFFRRL\nFBFBBBFRRL\nFBBFFFBLRL\nFBBBFFBRRL\nBBBFBBFLRR\nBFBBFBFRRR\nFFBFFBFRLL\nBBBFBBBRRL\nFBBFBFFLRL\nBBBFBFFRLL\nFFBBBFBLLL\nBFFFFFFLRL\nFFFBFFFLRL\nFBFBBFFLRL\nBFFBFBFLRR\nBBBFFFFRRR\nFBFBBBFLLL\nBFFBFBFRRL\nBFFFFBFRLR\nBFBBBFFLRL\nFBBBFBFRRL\nBFBBFBFLLR\nBFBFBBBLRR\nFFFBBBBLLR\nFBFBBBBLRR\nBFFBFFFLLR\nBBFBBFFLRR\nBBBBFBBRLR\nFFFBBFBLLR\nFFFBFBFLLL\nBBFBBFFLRL\nBFBBFFBRLL\nBFBFFFBRRR\nFFFBFBFLRR\nFFFBBFFRLL\nFFBBBBFRLL\nBFBBBFFRRR\nFBBFBBFRRL\nBBFBBBFRLR\nFBFBFBBLRL\nFBFBFFFLLR\nBFFFFBBRLL\nBFFFBBBRRR\nFBBFBFBRRR\nFBFBFBBLRR\nBFFBFFBLLL\nFFBFFFBLLR\nFBBBBBFRLL\nFFFBBBFRRR\nFFFBFFBRRL\nBBFBBBBLLL\nBBBBFBBLRL\nBBFFFFBRLL\nBFFFBBBRRL\nFBBBFBFRLL\nBFFBBFFRRR\nBFFFFFBLRL\nFBFBBFFLLR\nBFBBBFBRRL\nBBFBBFBRLR\nBBBBFBFLLR\nFBBBFBBRLR\nFBFBFBFLRL\nBBBFFBBRLR\nBBBBBFFLLL\nBFFBFBBRRL\nBBFFBBFRLR\nBFBBBBBLRR\nFFBFFBFRRR\nBFFFBBBLLR\nBFFFBFBRRL\nBFFFBFFLLR\nFBBFBBFLRR\nBBBFFFBRLR\nBBBBFBBRRR\nBFBFBFBRRR\nBBBFBBBLLR\nFBBFBFFRLL\nBBBFBBFRRR\nFFBFBBFRLL\nBFFBFFBRLL\nFBFBBFBLRL\nBBBFBFBRRR\nBFFFBBBRLR\nBFFBFBFRRR\nFFBFBFBRRR\nBFBFBBBRLR\nBBFFBFBLLR\nFFBFFFFRRR\nBFBBFFFLLL\nBFFFFBBLRL\nFBFFBFFLRR\nBFFBFFFRLL\nBBFBFFBRRL\nFBBFFBFLRL\nFBBFBFFLLL\nFBFFFBFLRR\nFBFBFFBRRR\nFBBBBFFLLL\nBFBFBBBRRR\nBBFBFBFLLR\nFFBBBFBRRL\nBBFBFFBLLL\nFBBBBFFRRR\nFFFBFBBLRR\nFFFBBBFRLR\nBBBFFFBRRL\nFFBBFBBLRR\nFFBBFBFRLL\nBBBFFFBRRR\nBBFBBFFRLL\nFBBFFBFLLL\nFBFBBFFRLL\nFFBBFBFRRR\nBBBBFFBLLR\nFBFFBBFLLL\nFBFBBFBLLR\nFFBBFFFRRR\nBBBFBBFRLL\nBFFBBFBRLL\nFBBBBBBRLL\nFBBFFFBRLR\nFFBBBBBLRL\nBFBBFBBLRL\nBBFBFBFRLR\nBFBFBBFLRR\nFFBFBBFRRR\nFFBFFFFLRL\nBBFBFFBRRR\nBBFBFBFLRL\nFFBFFBBRLL\nBBFFBFBLLL\nBBFFBBBLLR\nFBBFBFFLLR\nFBBBBFBRLL\nBBFFFFFLLL\nBBFBFFFRRL\nBBFFFBFRRR\nBBBFBFFLLR\nFFBBBFBLLR\nFBFBBBBLLR\nBBFFFFFRRL\nFBFBFBFRLR\nFFBFBBBRRL\nFFFBBFFRRL\nFFFBBBFLLL\nBFBFFFFRLR\nBFFBBBFRRL\nBFFFBBFLLR\nBFFBFFBRRL\nBBFBBBBRLL\nBBFBFBBRLL\nBBFBBBBRLR\nBFFFBBFRLR\nFBBFFBBLRL\nBBBBFBFRRL\nBFFBFFBRRR\nFFBBFFBLLR\nFFBBBBFRLR\nFBFFBBBLLR\nFFBBFFBRRL\nFBFBFFFRLR\nBFFBFFBLLR\nFBFFFBBLLL\nBFBBFFFLRL\nFFFBBFFRRR\nBFBBFFFRLR\nBFBFFBBRLL\nBBFFFFBRRR\nFBFFFFFLRL\nBFFBBBFLRL\nBFBBBFBLRR\nFFBFFBBLLL\nBFBFBBBRRL\nFBBBFBBLLR\nBFFFBFFRRL\nFBBFFBBLLR\nBBFFFFBLLR\nBFBBFBFRLL\nFBFBBBFLRL\nBFFBFFFLLL\nBBBBFBBRLL\nBBBFBBFRLR\nFBFFBBBRLL\nFBFBBBFLLR\nBFBBFFBRRL\nFBBFFFFLRL\nBFFBBBBRRL\nFBFBFBFRLL\nBFFBBBFLLR\nFBFBFFBRRL\nBBFBFBBRRR\nFFFFBBBRLR\nBBFFFBFRLR\nFBBBFFFRLR\nFBBBBFFLRR\nBBFFFBFLRR\nFBFBBFBRLL\nBBFFFBBRRL\nFFFBBBBLRL\nBBFBBFBRLL\nFBBBFFFRRL\nBFBFBBBLRL\nFFFBBBBRLL\nBBFFFFFLRL\nBBBBFBBLRR\nBFBBBBBRRR\nBBFFBFBRRR\nBFFFBFBLRL\nBFBFFBBLLL\nBBFFFFFLLR\nBFBBBFBRLR\nFBFFBFBLRL\nFFBFBFBLLR\nFBBFBBBRRL\nFBFBBBBRRL\nBFBFBFFLLL\nFBFFFFFRRR\nBBFBBBFLRL\nFFBFFFFRRL\nBBFBFFBLRL\nFFBFFFFLLR\nFBFFBBBLLL\nBBBBFBFRLR\nBBBFFBBLLR\nFFBBBBFLLL\nFBBBBFBRRR\nFBBBBBFRLR\nFFFBBBFLRL\nBBFFFBFLLL\nFBFFBFBRLR\nFFBBFBFLRL\nFFFBBFFLLR\nFFFBFFFRLR\nBFBFFBBRRL\nBBFFBFFRRR\nFFFBFBBRRR\nBFBFBFFLLR\nFBBBFFBLRL\nFFBBBBFRRL\nBBFBFBBLRR\nFFBBFFBRRR\nBFBBBBFLRR\nBFFBBBBLRR\nFBBBBFFLRL\nFBBFFFFRRL\nFFBBFBFLRR\nFFBBBFBRLL\nFFFBFBFRRL\nFFBBFFFLLL\nFBBFBFFRRL\nBBFFFBBRRR\nBBBFFBFRLR\nBFBFFBFRLL\nFBBFFFBLLR\nBFBBBBBRRL\nBFBBFBFLLL\nBBBFBFBRRL\nFBFFFBBRLL\nFBBFBFFRLR\nBFFBBBFLLL\nFFFBFBBRLL\nBBBBFFFRRL\nFBFBBFFLRR\nBFFFFBBRRL\nBBFBBFFRLR\nFBFFBFFLLR\nBBBFFBFLLL\nBFBBBBBLRL\nBFFFBFBLLL\nBFBFBBBLLR\nBBBFFFFLRR\nFBBBBBFRRR\nFBBFBFFRRR\nBBFBFBFRRR\nFFFBFFFLLL\nFBFFBBBRRL\nBBFBBBBLLR\nFBFBFBBRRR\nFFFBFBBLLL\nFFBBBFFRRR\nBBBFFFFRLR\nBFBFFFFLLR\nFBFFFFBRLL\nFBBFFBFLRR\nBFFBFFFRRL\nBBFBFBBLRL\nFFFBBFFLLL\nBFBBFFBRRR\nFFBBBBBLRR\nBBBFFFBLLL\nFFBFFFBLRR\nFFFBFBFRRR\nFFFBBFBLRR\nFBFFBBFRLR\nBFBBFFFLLR\nBFFBFFBLRR\nFBFFBBFRRR\nFFBBFFBRLL\nFBFBFBFLLR\nFBBFBBFLLL\nBFBBBBBRLL\nFBFFBFBLLR\nBFBBFBFLRL\nBFBBBFFRRL\nBBBFBFFRRL\nFFBBBFFLRR\nFBBBFFFRLL\nFFFBBFBRLL\nBFFBBFBRRL\nBFBFBBFLLR\nFFBFBBFRRL\nFBBBBFFLLR\nFFFBFFFRRR\nBBBFFFBLLR\nBBBFBBBLRR\nBBFBBFBLLL\nFFBBBFBLRR\nBFBFFBFLRR\nBFFBBFBLRL\nBFBBFFBLLR\nFBFFBBBRRR\nBBBFBBFLRL\nBFBFBFFRRL\nBBBFBBBLRL\nFBFBBBBRLL\nBFFFFBFRRL\nBFFFFFBRLL\nBFFFFBFLRL\nFBBFFBFRLL\nBBBFBFBLLR\nBBBFBBBRLR\nBFFBFFFLRR\nBFFBFBBLLR\nFBFBBFFRRL\nBBFBFBBLLR\nBFFFFFBLLL\nBFFBBFFLLR"
)

//nolint:scopelint
func TestID(t *testing.T) {
	tests := []struct {
		name string
		args day5.Pass
		want int
	}{
		{
			name: "example 1",
			args: "FBFBBFFRLR",
			want: 357,
		},

		{
			name: "example 2",
			args: "BFFFBBFRRR",
			want: 567,
		},

		{
			name: "example 3",
			args: "FFFBBBFRRR",
			want: 119,
		},

		{
			name: "example 4",
			args: "BBFFBBFRLL",
			want: 820,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.ID(); got != tt.want {
				t.Errorf("ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

//nolint:scopelint
func TestHighestSeatID(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "actual puzzle",
			args: set1,
			want: 994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day5.HighestSeatID(tt.args)
			if got != tt.want {
				t.Errorf("HighestSeatID() = %v, want %v", got, tt.want)
			}

			fmt.Printf("AdventOfCode/2020/Day/5/1: highest seat ID is: %d\n", got)
		})
	}
}

//nolint:scopelint
func TestFindMissingPassID(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "actual puzzle",
			args: set1,
			want: 741,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day5.FindFirstMissingSeatID(tt.args)
			if got != tt.want {
				t.Errorf("FindFirstMissingSeatID() = %v, want %v", got, tt.want)
			}

			fmt.Printf("AdventOfCode/2020/Day/5/2: missing seat ID is: %d\n", got)
		})
	}
}
