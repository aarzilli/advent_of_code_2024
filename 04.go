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

var lines []string

func get(i, j int) byte {
	if i < 0 || i >= len(lines) || j < 0 || j >= len(lines[i]) {
		return 0
	}
	return lines[i][j]
}

func check(si, sj, di, dj int) bool {
	return get(si, sj) == 'X' && get(si+di, sj+dj) == 'M' && get(si+2*di, sj+2*dj) == 'A' && get(si+3*di, sj+3*dj) == 'S'
}

func check2(si, sj, di, dj int) bool {
	return get(si, sj) == 'M' && get(si+di, sj+dj) == 'A' && get(si+2*di, sj+2*dj) == 'S'
}

func p2(i, j int) bool {
	return (check2(i-1, j-1, +1, +1) || check2(i+1, j+1, -1, -1)) &&
		(check2(i-1, j+1, +1, -1) || check2(i+1, j-1, -1, +1))
}

func main() {
	lines = Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	cnt := 0
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == 'X' {
				for di := -1; di <= 1; di++ {
					for dj := -1; dj <= 1; dj++ {
						if di == 0 && dj == 0 {
							continue
						}
						if check(i, j, di, dj) {
							cnt++
						}
					}
				}
			}
		}
	}
	Sol(cnt)

	cnt2 := 0
	for i := range lines {
		for j := range lines {
			if lines[i][j] == 'A' {
				if p2(i, j) {
					cnt2++
				}
			}
		}
	}
	Sol(cnt2)
}
