package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"log"
	"strconv"
	"strings"
)

func Day07Part2() {

	rules := aocio.ReadFileSplit("../../assets/input-07", "\n")

	required := requiredBags(shinyGold, rules)

	fmt.Printf("Solution day 7 part 2: %d\n", required)

}

func requiredBags(valid string, rules []string) int {

	for _, rule := range rules {

		if strings.HasPrefix(rule, valid) {
			split := strings.Split(rule, " bags contain ")

			if strings.HasPrefix(split[1], "no") {
				continue
			}

			// <num> <color> e.g. 1 dull magenta
			childBags := strings.Split(strings.ReplaceAll(split[1], ".", ""), ", ")
			sum := 0

			for _, childBag := range childBags {
				countBag := strings.SplitN(childBag, " ", 2)
				count, err := strconv.Atoi(countBag[0])
				if err != nil {
					log.Fatal("Failed to convert count")
				}
				bag := strings.ReplaceAll(countBag[1], " bags", "")
				sum += requiredBags(bag, rules) * count
				sum += count
			}
			return sum
		}
	}
	return 0
}
