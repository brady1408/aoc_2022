package day3

import (
	"log"
	"os"
)

var input string

func init() {
	data, err := os.ReadFile("day_3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = string(data)
}
