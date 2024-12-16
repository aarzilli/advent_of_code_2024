package main

import (
	. "aoc/util"
	"os"
)

type pos struct {
	i, j int
}

var M [][]byte

var start, end pos

type Dir int

const (
	Right Dir = iota
	Up
	Left
	Down
)

type state struct {
	p   pos
	dir Dir
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))
	M = make([][]byte, len(lines))
	for i := range lines {
		M[i] = []byte(lines[i])
		for j := range M[i] {
			if M[i][j] == 'S' {
				start = pos{i, j}
				M[i][j] = '.'
			} else if M[i][j] == 'E' {
				end = pos{i, j}
				M[i][j] = '.'
			}
		}
	}

	var bestpath []state
	first := true
	djk := NewDijkstra[state](state{start, Right})
	var cur state
	for djk.PopTo(&cur) {
		if cur.p == end {
			if first {
				bestpath = djk.PathTo(cur)
				first = false
				Sol(djk.Dist[cur])
			}
		}
		add := func(i, j int, d int) {
			if M[i][j] == '#' {
				return
			}
			djk.Add(cur, state{pos{i, j}, cur.dir}, d)
		}
		switch cur.dir {
		case Up:
			add(cur.p.i-1, cur.p.j, 1)
		case Down:
			add(cur.p.i+1, cur.p.j, 1)
		case Left:
			add(cur.p.i, cur.p.j-1, 1)
		case Right:
			add(cur.p.i, cur.p.j+1, 1)
		default:
			panic("blah")
		}
		djk.Add(cur, state{cur.p, (cur.dir + 1) % 4}, 1000)
		djk.Add(cur, state{cur.p, (cur.dir + 3) % 4}, 1000)
		djk.Add(cur, state{cur.p, (cur.dir + 2) % 4}, 2000)
	}

	inbest := make(Set[pos])
	seen := make(Set[state])
	fringe := make(Set[state])
	for _, p := range bestpath {
		inbest[p.p] = true
		fringe[p] = true
		seen[p] = true
	}

	for len(fringe) > 0 {
		p := OneKey(fringe)
		delete(fringe, p)

		check := func(i, j int, dir Dir, cost int) {
			n := state{pos{i, j}, dir}
			if !seen[n] && djk.Dist[n]+cost == djk.Dist[p] {
				for _, p2 := range djk.PathTo(n) {
					if seen[p2] {
						continue
					}
					fringe[p2] = true
					inbest[p2.p] = true
					seen[p2] = true
				}
			}
		}

		switch p.dir {
		case Up:
			check(p.p.i+1, p.p.j, Up, 1)
		case Down:
			check(p.p.i-1, p.p.j, Down, 1)
		case Right:
			check(p.p.i, p.p.j-1, Right, 1)
		case Left:
			check(p.p.i, p.p.j+1, Left, 1)
		}
	}

	Sol(len(inbest))
}

