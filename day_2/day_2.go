package day2

import (
	"fmt"
	"strings"
)

var scoreCard map[string]int
var correctedScoreCard map[string]int

// A rock
// B Paper
// C Scissors

// // X rock
// // Y Paper
// // Z Scissors

// X Lose
// Y Draw
// Z Win

// win 6
// draw 3
// lose 0

// rock 1
// paper 2
// scissors 3

func init() {
	scoreCard = map[string]int{
		"AX": 4,
		"AY": 8,
		"AZ": 3,
		"BX": 1,
		"BY": 5,
		"BZ": 9,
		"CX": 7,
		"CY": 2,
		"CZ": 6,
	}

	correctedScoreCard = map[string]int{
		"AX": 3,
		"AY": 4,
		"AZ": 8,
		"BX": 1,
		"BY": 5,
		"BZ": 9,
		"CX": 2,
		"CY": 6,
		"CZ": 7,
	}
}

func Part1() {
	inputSlice := strings.Split(fullInput, "\n")
	var score int
	for _, v := range inputSlice {
		score += scoreCard[fmt.Sprintf("%s%s", string(v[0]), string(v[2]))]
	}
	fmt.Println(score)
}

func Part2() {
	inputSlice := strings.Split(fullInput, "\n")
	var score int
	for _, v := range inputSlice {
		score += correctedScoreCard[fmt.Sprintf("%s%s", string(v[0]), string(v[2]))]
	}
	fmt.Println(score)
}
