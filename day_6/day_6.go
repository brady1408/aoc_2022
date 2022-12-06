package day6

import (
	"fmt"
	"log"

	"github.com/samber/lo"
)

type buff struct {
	size  int
	value []string
}

func newBuff(i int) buff {
	return buff{
		size:  i,
		value: []string{},
	}
}

func (b *buff) push(s string) {
	if len(s) > 1 {
		log.Fatal("push only supports one character")
	}
	b.value = append(b.value, s)
	if len(b.value) > b.size {
		b.value = b.value[len(b.value)-b.size:]
	}
}

func Part1() {
	b := newBuff(4)
	for k, v := range input {
		b.push(string(v))

		if a := lo.Uniq(b.value); len(a) == b.size {
			fmt.Println(k+1, b, len(b.value))
			return
		}
	}
}

func Part2() {
	b := newBuff(14)
	for k, v := range input {
		b.push(string(v))
		if a := lo.Uniq(b.value); len(a) == b.size {
			fmt.Println(k+1, b, len(b.value))
			return
		}
	}
}
