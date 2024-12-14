package main

import (
	. "aoc/util"
	"os"
)

type pos struct {
	i, j int
}

func main() {
	for _, t := range []struct {
		fname        string
		part1, part2 int
	}{
		{"12.example", 140, 80},
		{"12.example2", 1930, 1206},
		{"12.example3", -1, 368},
		{"12.example4", -1, 236},
	} {
		part1, part2 := solve(t.fname)
		if t.part1 >= 0 && t.part1 != part1 {
			panic("failed")
		}
		if t.part2 != part2 {
			panic("failed")
		}

	}
	part1, part2 := solve(os.Args[1])
	Sol(part1, 1486324)
	Sol(part2, 898684)
}

func solve(fname string) (int, int) {
	lines := Input(fname, "\n", true)
	M := make([][]byte, len(lines))
	for i := range lines {
		M[i] = []byte(lines[i])
	}
	var seen = make(Set[pos])
	part1, part2 := 0, 0
	for i := range M {
		for j := range M[i] {
			if seen[pos{i, j}] {
				continue
			}

			start := pos{i, j}

			m := mark(start, M, seen)
			p := perimeter(start, M, m)
			s := sides(start, M, m)
			part1 += len(m) * p
			part2 += len(m) * s
		}
	}
	return part1, part2
}

func mark(start pos, M [][]byte, seen Set[pos]) Set[pos] {
	fringe := []pos{start}
	m := make(Set[pos])
	for len(fringe) > 0 {
		cur := fringe[0]
		fringe = fringe[1:]
		if m[cur] || seen[cur] {
			continue
		}
		m[cur] = true
		seen[cur] = true

		add := func(i, j int) {
			if i < 0 || i >= len(M) || j < 0 || j >= len(M[i]) || seen[pos{i, j}] {
				return
			}
			if M[i][j] != M[cur.i][cur.j] {
				return
			}
			fringe = append(fringe, pos{i, j})
		}

		add(cur.i+1, cur.j)
		add(cur.i, cur.j+1)
		add(cur.i-1, cur.j)
		add(cur.i, cur.j-1)
	}
	return m
}

func get(M [][]byte, p pos) byte {
	if p.i < 0 || p.i >= len(M) || p.j < 0 || p.j >= len(M[p.i]) {
		return 0
	}
	return M[p.i][p.j]
}

func perimeter(start pos, M [][]byte, m Set[pos]) int {
	ch := M[start.i][start.j]
	r := 0
	for p := range m {
		for _, n := range []pos{pos{p.i - 1, p.j}, pos{p.i + 1, p.j}, pos{p.i, p.j - 1}, pos{p.i, p.j + 1}} {
			if get(M, n) != ch {
				r++
			}
		}
	}
	return r
}

func isborderv(M [][]byte, p pos) bool {
	ch := get(M, p)
	for _, n := range []pos{pos{p.i, p.j - 1}, pos{p.i, p.j + 1}} {
		if get(M, n) != ch {
			return true
		}
	}
	return false
}

type Dir int

const (
	Down Dir = iota
	Right
	Up
	Left
)

var dirs = []pos{pos{+1, 0}, pos{0, +1}, pos{-1, 0}, pos{0, -1}}

func sides2(start pos, m Set[pos]) (int, Set[pos]) {
	seen := make(Set[pos])

	dir := Down
	cur := start

	outcur := pos{start.i, start.j - 1}
	if m[outcur] {
		outcur = pos{start.i, start.j + 1}
		if m[outcur] {
			panic("can't find outside")
		}
	}

	r := 0

	for {
		n := pos{cur.i + dirs[dir].i, cur.j + dirs[dir].j}
		outn := pos{outcur.i + dirs[dir].i, outcur.j + dirs[dir].j}
		if m[n] && !m[outn] {
			cur = n
			outcur = outn
		} else if !m[n] {
			// convex angle
			switch dir {
			case Down, Up:
				if outcur.j < cur.j {
					dir = Right
				} else if outcur.j > cur.j {
					dir = Left
				} else {
					panic("blah")
				}
			case Right, Left:
				if outcur.i < cur.i {
					dir = Down
				} else if outcur.i > cur.i {
					dir = Up
				} else {
					panic("blah")
				}
			}
			outcur = n
			r++
		} else if m[n] {
			// concave angle
			switch dir {
			case Down, Up:
				if outcur.j < cur.j {
					dir = Left
				} else if outcur.j > cur.j {
					dir = Right
				} else {
					panic("blah")
				}
			case Right, Left:
				if outcur.i < cur.i {
					dir = Up
				} else if outcur.i > cur.i {
					dir = Down
				} else {
					panic("blah")
				}
			}
			cur = outn
			r++
		} else {
			panic("blah")
		}
		seen[cur] = true
		if cur == start && dir == Down {
			break
		}
	}
	return r, seen
}

func sides(start pos, M [][]byte, m Set[pos]) int {
	r, seen := sides2(start, m)
	for i := range M {
		for j := range M[i] {
			p := pos{i, j}
			if !m[p] || seen[p] {
				continue
			}
			if !isborderv(M, p) {
				continue
			}
			r2, seen2 := sides2(p, m)
			for p2 := range seen2 {
				seen[p2] = true
			}
			r += r2
		}
	}
	return r
}
