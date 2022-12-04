package day4

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func Part1() {
	var subsets int
	groups := explode(input)
	for _, v := range groups {
		results := lo.Intersect(v.groupA, v.groupB)
		if len(v.groupA) == len(results) || len(v.groupB) == len(results) {
			subsets++
		}
	}
	fmt.Println(subsets)
}

func Part2() {
	var subsets int
	groups := explode(input)
	for _, v := range groups {
		results := lo.Intersect(v.groupA, v.groupB)
		if len(results) > 0 {
			subsets++
		}
	}
	fmt.Println(subsets)
}

type group struct {
	groupA []int
	groupB []int
}

func explode(in string) []group {
	out := []group{}
	groups := strings.Split(in, "\n")
	for _, v := range groups {
		inputGroups := strings.Split(v, ",")
		assignmentA := strings.Split(inputGroups[0], "-")
		start, err := strconv.Atoi(assignmentA[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(assignmentA[1])
		if err != nil {
			log.Fatal(err)
		}
		groupASlice := []int{}
		for i := start; i <= end; i++ {
			groupASlice = append(groupASlice, i)
		}
		assignmentB := strings.Split(inputGroups[1], "-")
		start, err = strconv.Atoi(assignmentB[0])
		if err != nil {
			log.Fatal(err)
		}
		end, err = strconv.Atoi(assignmentB[1])
		if err != nil {
			log.Fatal(err)
		}
		groupBSlice := []int{}
		for i := start; i <= end; i++ {
			groupBSlice = append(groupBSlice, i)
		}
		out = append(out, group{groupA: groupASlice, groupB: groupBSlice})
	}
	return out
}
