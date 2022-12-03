package day3

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	// lop "github.com/samber/lo/parallel"
)

const key string = `_abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ`

type rucksack struct {
	group                 int
	uniqueItems           string
	compartmentA          string
	compartmentB          string
	intersectingCharacter string
}

func Part1() {
	rucksacks := getRucksacks(input)

	// add them up
	var total int
	for _, v := range rucksacks {
		total += strings.Index(key, v.intersectingCharacter)
	}
	fmt.Println(total)
}

func Part2() {
	rucksacks := getRucksacks(input)

	groupItems := make(map[int]map[string]int)

	// find groups
	for _, value := range rucksacks {
		for _, v := range value.uniqueItems {
			if _, ok := groupItems[value.group]; !ok {
				groupItems[value.group] = make(map[string]int)
			}
			groupItems[value.group][string(v)]++
		}
	}

	groupBadges := make(map[int]string)
	for key, value := range groupItems {
		for k, v := range value {
			if v == 3 {
				groupBadges[key] = k
			}
		}
	}

	var total int
	for _, v := range groupBadges {
		total += strings.Index(key, v)
	}

	fmt.Println(total)
}

func getRucksacks(in string) []rucksack {
	rucksacks := []rucksack{}
	// inSlice := []string{}
	// for _, v := range in {
	// 	inSlice = append(inSlice, string(v))
	// }

	for k, v := range strings.Split(in, "\n") {
		r := rucksack{
			group:        k / 3,
			uniqueItems:  string(lo.Uniq([]byte(v))),
			compartmentA: v[:(len(v) / 2)],
			compartmentB: v[(len(v) / 2):],
		}
		fmt.Println(r.uniqueItems)
		hash := make(map[string]bool)
		for _, v := range r.compartmentA {
			hash[string(v)] = true
		}
		for _, v := range r.compartmentB {
			if hash[string(v)] {
				r.intersectingCharacter = string(v)
			}
		}
		rucksacks = append(rucksacks, r)
	}
	return rucksacks
}
