package main

import (
	"fmt"
	"github.com/j-sattler/adventofcode-2020/internal/aocio"
	"sort"
)

func Day05Part2() {
	passes := aocio.ReadFileSplit("../../assets/input-05", "\n")
	var myId int

	ids := make([]int, len(passes))
	for i, pass := range passes {
		row := 0
		col := 0
		for _, v := range pass {
			switch v {
			case 'B': // shift row 1 and set to 1
				row <<= 1
				row |= 1
			case 'F': // shift row 1 keep 0
				row <<= 1
			case 'R': // shift col 1 set to 1
				col <<= 1
				col |= 1
			case 'L': // shift col 1 keep 0
				col <<= 1
			}
		}

		// insert id; unsorted
		ids[i] = row*8 + col
	}

	// sort array
	sort.Ints(ids)

	for i, id := range ids {
		// check if the next id is + 2 away to find missing seat
		if id+2 == ids[i+1] {
			myId = id + 1
			break
		}
	}

	fmt.Printf("solution day 5 part 2: %d\n", myId)

}
