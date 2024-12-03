package main

import (
	. "aoc/util"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	tot := 0
	for _, line := range lines {
		v := re.FindAllStringSubmatch(line, -1)
		for _, m := range v {
			tot += Atoi(m[1]) * Atoi(m[2])
		}
	}
	Sol(tot)

	re2 := regexp.MustCompile(`(?:mul\((\d+),(\d+)\))|don't\(\)|do\(\)`)
	enabled := true
	tot2 := 0
	for _, line := range lines {
		v := re2.FindAllStringSubmatch(line, -1)
		for _, m := range v {
			switch {
			case strings.HasPrefix(m[0], "mul("):
				if enabled {
					tot2 += Atoi(m[1]) * Atoi(m[2])
				}
			case strings.HasPrefix(m[0], "don't("):
				enabled = false
			case strings.HasPrefix(m[0], "do("):
				enabled = true
			default:
				panic("blah")
			}
		}
	}
	Sol(tot2)
}
