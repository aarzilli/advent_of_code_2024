package main

import (
	. "aoc/util"
	"fmt"
	"os"
	"sort"
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
	var left, right []int
	rightm := make(map[int]int)
	for i := range lines {
		v := Getints(lines[i], false)
		left = append(left, v[0])
		right = append(right, v[1])
		rightm[v[1]]++
	}
	sort.Ints(left)
	sort.Ints(right)
	tot := 0
	for i := range left {
		tot += Abs(left[i] - right[i])
	}
	Sol(tot)
	
	tot2 := 0
	for _, x := range left {
		tot2 += x * rightm[x]
	}
	Sol(tot2)
}
