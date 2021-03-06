package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
)

func Day15Part2() {
	initStr := aocio.ReadFileSplit("../../assets/input-15", ",")
	numbers := parseNumbers(initStr)
	cache := make(map[int]int)

	for i, n := range numbers[:len(numbers)-1] {
		cache[n] = i
	}

	for i := len(numbers); i < 30000000; i++ {
		current := 0
		prev := numbers[i-1]
		if j, ok := cache[prev]; ok {
			current = i - (j + 1)
		}
		cache[prev] = i - 1
		numbers = append(numbers, current)
	}

	fmt.Printf("Solution day 15 part 2: %d\n", numbers[len(numbers)-1])
}
