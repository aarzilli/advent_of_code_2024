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

type pos2 struct {
	i, j int
}

type pos struct {
	pos2
	dir int
}

var M [][]byte

func step(cur *pos) bool {
	next := *cur
	switch cur.dir {
	case 0:
		next.i--
	case 1:
		next.j++
	case 2:
		next.i++
	case 3:
		next.j--
	}
	if next.i < 0 || next.i >= len(M) || next.j < 0 || next.j >= len(M[next.i]) {
		return true
	}
	if M[next.i][next.j] == '#' {
		cur.dir = (cur.dir + 1) % 4
		return false
	}
	*cur = next
	return false
}

func isloop(start pos) bool {
	cur := start
	seen := make(Set[pos])
	for {
		if seen[cur] {
			return true
		}
		seen[cur] = true
		if step(&cur) {
			return false
		}
	}
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	M = make([][]byte, len(lines))
	found := false
	var cur pos
	for i := range lines {
		M[i] = []byte(lines[i])
		for j := range M {
			if M[i][j] == '^' {
				cur.i = i
				cur.j = j
				cur.dir = 0
				M[i][j] = '.'
				found = true
			}
		}
	}
	if !found {
		panic("blah")
	}
	start := cur

	seen := make(Set[pos2])
	for {
		seen[cur.pos2] = true
		if step(&cur) {
			break
		}
	}
	Sol(len(seen))

	part2 := 0
	cnt := 0
	for i := range M {
		for j := range M[i] {
			if M[i][j] == '.' {
				cnt++
				pln(cnt, len(M)*len(M[0]))
				M[i][j] = '#'
				il := isloop(start)
				if il {
					part2++
				}
				M[i][j] = '.'
			}
		}
	}
	Sol(part2)
}
