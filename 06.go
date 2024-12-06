package main

import (
	. "aoc/util"
	"os"
)

type pos struct {
	i, j int
}

type posdir struct {
	pos
	dir int
}

var M [][]byte

func step(cur *posdir) bool {
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

func isloop(start posdir) bool {
	cur := start
	seen := make(Set[posdir])
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
	M = make([][]byte, len(lines))
	var start posdir
	for i := range lines {
		M[i] = []byte(lines[i])
		for j := range M {
			if M[i][j] == '^' {
				start.i = i
				start.j = j
				start.dir = 0
				M[i][j] = '.'
			}
		}
	}
	cur := start

	seen := make(Set[pos])
	for {
		seen[cur.pos] = true
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
