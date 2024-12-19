package main

import (
	. "aoc/util"
	"os"
	"regexp"
	"strings"
)

func main() {
	groups := Input(os.Args[1], "\n\n", true)
	Pf("len %d\n", len(groups))
	pats := Noempty(Spac(groups[0], ",", -1))
	Pf("%q\n", pats)
	rx := regexp.MustCompile("^(" + strings.ReplaceAll(groups[0], ", ", "|") + ")+$")
	Pln(rx)
	lines := Noempty(Spac(groups[1], "\n", -1))
	part1 := 0
	part2 := 0
	for _, line := range lines {
		if rx.MatchString(line) {
			Pln(line, "is possible", rx.FindString(line))
			part1++
			n := enum(line, pats, make(map[string]int))
			Pln("matches", n)
			part2 += n
		}
	}
	Sol(part1)
	Sol(part2)
}

func enum(line string, pats []string, memo map[string]int) int {
	if line == "" {
		memo[line] = 1
		return 1
	}
	if r, ok := memo[line]; ok {
		return r
	}
	r := 0
	for _, pat := range pats {
		if strings.HasPrefix(line, pat) {
			r += enum(line[len(pat):], pats, memo)
		}
	}
	memo[line] = r
	return r
}
