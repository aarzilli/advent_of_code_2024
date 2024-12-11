package main

import (
	. "aoc/util"
	"fmt"
	"os"
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
	v := Getints(lines[0], false)
	m := make(map[int]int)
	for _, x := range v {
		m[x]++
	}
	pln(m)
	for i := 0; i < 75; i++ {
		m = step(m)
		pln(i, tot(m), len(m))
		//pln(m)
	}
	Sol(tot(m))
}

func step(m map[int]int) map[int]int {
	r := make(map[int]int)
	for x, cnt := range m {
		if x == 0 {
			r[1] += cnt
		} else {
			s := fmt.Sprintf("%d", x)
			if len(s)%2 == 0 {
				r[Atoi(s[:len(s)/2])] += cnt
				r[Atoi(s[len(s)/2:])] += cnt
			} else {
				r[x*2024] += cnt
			}
		}
	}
	return r
}

func tot(m map[int]int) int {
	r := 0
	for _, cnt := range m {
		r += cnt
	}
	return r
}
