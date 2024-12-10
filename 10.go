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

type pos struct {
	i, j int
}

var M [][]int

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	M = make([][]int, len(lines))
	for i := range lines {
		M[i] = make([]int, len(lines[i]))
		for j := range lines[i] {
			M[i][j] = int(lines[i][j] - '0')
		}
	}

	part1 := 0
	for i := range M {
		for j := range M[i] {
			if M[i][j] == 0 {
				k := visit(pos{i, j})
				pln(k)
				part1 += k
			}
		}
	}
	Sol(part1)

	part2 := 0
	for i := range M {
		for j := range M[i] {
			if M[i][j] == 0 {
				k := visit2(pos{i, j})
				pln(k)
				part2 += k
			}
		}
	}
	Sol(part2)

}

func visit(start pos) int {
	fringe := []pos{start}
	seen := make(Set[pos])
	dest := make(Set[pos])

	for len(fringe) > 0 {
		cur := fringe[0]
		fringe = fringe[1:]
		if seen[cur] {
			continue
		}
		seen[cur] = true

		add := func(i, j int) {
			next := pos{i, j}
			if next.i < 0 || next.i >= len(M) || next.j < 0 || next.j >= len(M[next.i]) {
				return
			}
			if M[cur.i][cur.j]+1 != M[next.i][next.j] {
				return
			}
			fringe = append(fringe, next)
		}

		if M[cur.i][cur.j] == 9 {
			dest[cur] = true
		}

		add(cur.i+1, cur.j)
		add(cur.i-1, cur.j)
		add(cur.i, cur.j+1)
		add(cur.i, cur.j-1)
	}
	return len(dest)
}

func visit2(cur pos) int {
	if M[cur.i][cur.j] == 9 {
		return 1
	}
	r := 0
	for _, next := range []pos{pos{cur.i + 1, cur.j}, pos{cur.i - 1, cur.j}, pos{cur.i, cur.j + 1}, pos{cur.i, cur.j - 1}} {
		if next.i < 0 || next.i >= len(M) || next.j < 0 || next.j >= len(M[next.i]) {
			continue
		}
		if M[cur.i][cur.j]+1 != M[next.i][next.j] {
			continue
		}
		r += visit2(next)
	}
	return r
}
