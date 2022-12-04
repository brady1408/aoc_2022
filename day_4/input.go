package day4

import (
	"log"
	"os"
)

var input string

func init() {
	data, err := os.ReadFile("day_4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input = string(data)
}
