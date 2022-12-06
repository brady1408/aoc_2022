package day5

import (
	"log"
	"os"
)

var input string

func init() {
	data, err := os.ReadFile("day_5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = string(data)
}
