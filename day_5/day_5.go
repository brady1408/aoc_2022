package day5

import (
	"fmt"
	"log"
	"strings"
	"unicode"
)

type stack []string

type instruction struct {
	amount int
	from   int
	to     int
}

func (s *stack) push(in string) {
	*s = append(*s, in)
}
func (s *stack) pushN(v []string) {
	*s = append(*s, v...)
}
func (s *stack) pop() string {
	if len(*s) == 0 {
		log.Fatal("stack is empty")
	}
	out := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return out
}
func (s *stack) popN(n int) []string {
	if len(*s) < n {
		panic("stack is empty")
	}
	v := (*s)[len(*s)-n:]
	*s = (*s)[:len(*s)-n]
	return v
}

func (s *stack) Peek() string {
	if len(*s) == 0 {
		return ""
	}
	return (*s)[len(*s)-1]
}

func Part1() {

	// Split the two input sections
	inputs := strings.Split(input, "\n\n")

	stacks := parseChart(inputs[0])
	instructions := parseInstructions(inputs[1])

	processStacks(stacks, instructions)

	peakAll(stacks)
}

func Part2() {
	// Split the two input sections
	inputs := strings.Split(input, "\n\n")

	stacks := parseChart(inputs[0])
	instructions := parseInstructions(inputs[1])

	processStacks9001(stacks, instructions)

	peakAll(stacks)
}

func parseChart(s string) []stack {
	// split input
	chartLines := strings.Split(s, "\n")

	stacks := []stack{}

	//loop backwards to fill the stacks correctly
	for i := len(chartLines) - 1; i >= 0; i-- {
		for k, v := range chartLines[i] {
			if unicode.IsLetter(v) {
				pos := k / 4
				if len(stacks) < pos+1 {
					stacks = append(stacks, stack{})
				}
				stacks[pos].push(string(v))
			}
		}
	}

	return stacks
}

func parseInstructions(s string) []instruction {
	out := []instruction{}
	lines := strings.Split(s, "\n")
	for _, v := range lines {
		var amount, from, to int
		fmt.Sscanf(v, "move %d from %d to %d", &amount, &from, &to)
		out = append(out, instruction{amount: amount, from: from, to: to})
	}
	return out
}

func processStacks(s []stack, i []instruction) {
	for _, v := range i {
		for i := 0; i < v.amount; i++ {
			val := s[v.from-1].pop()
			s[v.to-1].push(val)
		}
	}
}

func processStacks9001(s []stack, i []instruction) {
	for _, v := range i {
		val := s[v.from-1].popN(v.amount)
		s[v.to-1].pushN(val)
	}
}

func peakAll(stacks []stack) {
	var out string
	for _, v := range stacks {
		out += v.Peek()
	}
	fmt.Println(out)
}
