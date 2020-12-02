package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Policy struct {
	min, max int
	char string
	pwd string
}

func main(){
	content, err := ioutil.ReadFile("assets/input-2")
	var counter int

	if err != nil { log.Fatal(err) }

	// split input by new line
	strArr := strings.Split(string(content), "\n")

	for _, s := range strArr {
		policy := parsePolicy(s)
		if policy.isValid() { counter++ }
	}
	fmt.Printf("%d passwords are valid\n", counter)
}

func parsePolicy(pStr string) *Policy {
	splitBySpace := strings.Split(pStr, " ")
	minMax := strings.Split(splitBySpace[0], "-")
	char := strings.Split(splitBySpace[1], ":")[0]
	passwd := splitBySpace[2]

	min, err:= strconv.Atoi(minMax[0])
	if err != nil { log.Fatal(err) }

	max, err := strconv.Atoi(minMax[1])
	if err != nil { log.Fatal(err) }

	return &Policy {
		min:  min,
		max:  max,
		char: char,
		pwd:  passwd,
	}
}

func (p *Policy) isValid() bool {
	if p.min - 1 < 0 || p.min > len(p.pwd) || p.max - 1 < 0 || p.max > len(p.pwd) {
		return false
	}
	first := string(p.pwd[p.min - 1])
	second := string(p.pwd[p.max - 1])
	return (first == p.char) != (second == p.char)
}
