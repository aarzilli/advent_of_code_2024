package main

import (
	. "aoc/util"
	"os"
)

type pos struct {
	i, j int
}

var M [][]int

func main() {
	lines := Input(os.Args[1], "\n", true)
	M = make([][]int, len(lines))
	for i := range lines {
		M[i] = make([]int, len(lines[i]))
		for j := range lines[i] {
			M[i][j] = int(lines[i][j] - '0')
		}
	}

	part1 := 0
	part2 := 0
	for i := range M {
		for j := range M[i] {
			if M[i][j] == 0 {
				d := make(Set[pos])
				k := visit2(pos{i, j}, d)
				part1 += len(d)
				part2 += k
			}
		}
	}
	Sol(part1)
	Sol(part2)

}

func visit2(cur pos, d Set[pos]) int {
	if M[cur.i][cur.j] == 9 {
		d[cur] = true
		return 1
	}
	r := 0
	for _, next := range []pos{pos{cur.i + 1, cur.j}, pos{cur.i - 1, cur.j}, pos{cur.i, cur.j + 1}, pos{cur.i, cur.j - 1}} {
		if next.i < 0 || next.i >= len(M) || next.j < 0 || next.j >= len(M[next.i]) || M[cur.i][cur.j]+1 != M[next.i][next.j] {
			continue
		}
		r += visit2(next, d)
	}
	return r
}
