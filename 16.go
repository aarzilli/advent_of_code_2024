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
	fringe := make(Set[state])
	for _, p := range bestpath {
		inbest[p.p] = true
		fringe[p] = true
	}

	for len(fringe) > 0 {
		p := OneKey(fringe)
		delete(fringe, p)

		check := func(i, j int, dir Dir, cost int) {
			n := state{pos{i, j}, dir}
			if djk.Dist[n]+cost == djk.Dist[p] {
				fringe[n] = true
				inbest[n.p] = true
			}
		}

		switch p.dir {
		case Up:
			check(p.p.i+1, p.p.j, Up, 1)
			check(p.p.i, p.p.j, Right, 1000)
			check(p.p.i, p.p.j, Left, 1000)
			check(p.p.i, p.p.j, Down, 2000)
		case Down:
			check(p.p.i-1, p.p.j, Down, 1)
			check(p.p.i, p.p.j, Right, 1000)
			check(p.p.i, p.p.j, Left, 1000)
			check(p.p.i, p.p.j, Up, 2000)
		case Right:
			check(p.p.i, p.p.j-1, Right, 1)
			check(p.p.i, p.p.j, Up, 1000)
			check(p.p.i, p.p.j, Down, 1000)
			check(p.p.i, p.p.j, Left, 2000)
		case Left:
			check(p.p.i, p.p.j+1, Left, 1)
			check(p.p.i, p.p.j, Up, 1000)
			check(p.p.i, p.p.j, Down, 1000)
			check(p.p.i, p.p.j, Right, 2000)
		}
	}

	Sol(len(inbest))

	/*
		seen := make(Set[pos])
		seen[start] = true
		enum(start, Right, 0, mindist, seen, []pos{start})
		Sol(len(inbest))*/
}

/*
var inbest = make(Set[pos])

func enum(cur pos, dir Dir, dist int, mindist int, seen Set[pos], p []pos) bool {
	if dist > mindist {
		return false
	}
	//Pln("enum", p, dist)
	if cur == end {
		for p := range seen {
			inbest[p] = true
		}
		Pln("minpath", dist, p)
	}
	enumto := func(i, j int) bool {
		if M[i][j] == '#' || seen[pos{i,j}] {
			return false
		}
		seen[pos{i,j}] = true
		r := enum(pos{i,j}, dir, dist+1, mindist, seen, append(p, pos{i,j}))
		delete(seen, pos{i,j})
		return r
	}
	enumto2 := func(newdir Dir, dd int) bool {
		r := enum(cur, newdir, dist+dd, mindist, seen, p)
		return r
	}
	switch dir {
	case Up:
		enumto(cur.i-1, cur.j)
	case Down:
		enumto(cur.i+1, cur.j)
	case Left:
		enumto(cur.i, cur.j-1)
	case Right:
		enumto(cur.i, cur.j+1)
	default:
		panic("blah")
	}
	enumto2((dir+1)%4, 1000)
	enumto2((dir+3)%4, 1000)
	enumto2((dir+2)%4, 2000)
	return false
}
*/
