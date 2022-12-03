package day1

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

func Part1() {
	fmt.Println(findMost(input))
}

func Part2() {
	fmt.Println(findTop3(input))
}

func findMost(in string) int {
	ss := strings.Split(in, "\n\n")
	var totals []int
	for _, v := range ss {
		ts := strings.Split(v, "\n")
		var total int
		for _, v := range ts {
			i, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			total += i
		}
		totals = append(totals, total)
	}
	sort.IntSlice(totals).Sort()
	return totals[len(totals)-1]
}

func findTop3(in string) int {
	ss := strings.Split(in, "\n\n")
	var totals []int
	for _, v := range ss {
		ts := strings.Split(v, "\n")
		var total int
		for _, v := range ts {
			i, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			total += i
		}
		totals = append(totals, total)
	}
	sort.IntSlice(totals).Sort()
	top := totals[len(totals)-3:]
	var top3 int
	for _, v := range top {
		top3 += v
	}
	return top3
}
