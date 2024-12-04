package main

import (
	. "aoc/util"
	"os"
)

var lines []string

func get(i, j int) byte {
	if i < 0 || i >= len(lines) || j < 0 || j >= len(lines[i]) {
		return 0
	}
	return lines[i][j]
}

func checkstr(si, sj, di, dj int, str string) bool {
	for i := range str {
		if get(si+i*di, sj+i*dj) != str[i] {
			return false
		}
	}
	return true
}

func main() {
	lines = Input(os.Args[1], "\n", true)
	cnt := 0
	for i := range lines {
		for j := range lines[i] {
			if lines[i][j] == 'X' {
				for di := -1; di <= 1; di++ {
					for dj := -1; dj <= 1; dj++ {
						if di == 0 && dj == 0 {
							continue
						}
						if checkstr(i, j, di, dj, "XMAS") {
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
				if (checkstr(i-1, j-1, +1, +1, "MAS") || checkstr(i+1, j+1, -1, -1, "MAS")) && (checkstr(i-1, j+1, +1, -1, "MAS") || checkstr(i+1, j-1, -1, +1, "MAS")) {
					cnt2++
				}
			}
		}
	}
	Sol(cnt2)
}
