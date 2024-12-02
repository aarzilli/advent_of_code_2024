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

func issafe(v []int) bool {
	alldecreasing, allincreasing, distok := true, true, true
	for i := 1; i < len(v); i++ {
		if v[i] >= v[i-1] {
			alldecreasing = false
		}
		if v[i] <= v[i-1] {
			allincreasing = false
		}
		d := Abs(v[i] - v[i-1])
		if d < 1 || d > 3 {
			distok = false
		}
	}
	return (allincreasing || alldecreasing) && distok
}

func remove1(v []int, i int) []int {
	r := append([]int{}, v[:i]...)
	r = append(r, v[i+1:]...)
	return r
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	cnt, cnt2 := 0, 0
	for i := range lines {
		v := Vatoi(Spac(lines[i], " ", -1))
		ok := issafe(v)
		if ok {
			cnt++
		}

		ok2 := false
		for i := range v {
			v2 := remove1(v, i)
			if issafe(v2) {
				ok2 = true
				break
			}
		}
		if ok2 || ok {
			cnt2++
		}
	}
	Sol(cnt)
	Sol(cnt2)
}
