package main

import (
	. "aoc/util"
	"maps"
	"os"
	"slices"
)

var M [][]byte
var Dims [2]int

var start, end pos

type pos struct {
	i, j int
}

type cheat struct {
	s1, s2 pos
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))

	M = make([][]byte, len(lines))
	for i := range lines {
		M[i] = []byte(lines[i])
		for j := range M[i] {
			switch M[i][j] {
			case 'S':
				start = pos{i, j}
				M[i][j] = '.'
			case 'E':
				end = pos{i, j}
				M[i][j] = '.'
			}
		}
	}

	Pln(start, end)

	djk := NewDijkstra[pos](end)
	var cur pos
	for djk.PopTo(&cur) {
		add := func(i, j int) {
			if i < 0 || i >= len(M) || j < 0 || j >= len(M[i]) || M[i][j] == '#' {
				return
			}
			djk.Add(cur, pos{i, j}, 1)
		}
		add(cur.i+1, cur.j)
		add(cur.i-1, cur.j)
		add(cur.i, cur.j+1)
		add(cur.i, cur.j-1)
	}

	Dims = [2]int{len(M), len(M[0])}

	minsave := 100
	if os.Args[1] == "20.example" {
		minsave = 1
	}

	part1 := getcheats(djk, 2, minsave)

	if os.Args[1] == "20.example" {
		for _, k := range slices.Sorted(maps.Keys(part1)) {
			Pln("there is", len(part1[k]), "cheat that saves", k, "seconds")
		}
	}

	cnt := 0
	for k := range part1 {
		if k >= 100 {
			cnt += len(part1[k])
		}
	}
	Sol(cnt)

	part2 := getcheats(djk, 20, minsave)

	if os.Args[1] == "20.example" {
		for _, k := range slices.Sorted(maps.Keys(part2)) {
			if k < 50 {
				continue
			}
			Pln("there is", len(part2[k]), "cheat that saves", k, "seconds")
		}
	}

	cnt2 := 0
	for k := range part2 {
		if k >= 100 {
			cnt2 += len(part2[k])
		}
	}
	Sol(cnt2)
}

func dist(p1, p2 pos) int {
	return Abs(p1.i-p2.i) + Abs(p1.j-p2.j)
}

func getcheats(djk *Dijkstra[pos], D, minsave int) map[int]Set[cheat] {
	part2 := make(map[int]Set[cheat])
	for i := range M {
		for j := range M[i] {
			p := pos{i, j}
			toend, ok := djk.Dist[p]
			if !ok {
				continue
			}
			if toend < minsave {
				continue
			}

			for i := max(p.i-D, 0); i < min(p.i+D+1, len(M)); i++ {
				remd := D - Abs(p.i-i)
				for j := max(p.j-remd, 0); j < min(p.j+remd+1, len(M[0])); j++ {
					p2 := pos{i, j}

					if p == p2 {
						continue
					}

					if M[p2.i][p2.j] != '.' {
						continue
					}

					steps := dist(p, p2)
					if steps < 2 {
						continue
					}

					newtoend, ok := djk.Dist[p2]
					if !ok {
						continue
					}

					saved := toend - newtoend - steps

					if saved < minsave {
						continue
					}

					if part2[saved] == nil {
						part2[saved] = make(Set[cheat])
					}

					part2[saved][cheat{p, p2}] = true
				}
			}
		}
	}
	return part2
}
