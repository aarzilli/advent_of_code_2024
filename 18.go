package main

import (
	. "aoc/util"
	"os"
)

var M [][]byte

type pos struct {
	i, j int
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))

	sz := 71
	if len(lines) < 30 {
		sz = 7
	}

	M = make([][]byte, sz)
	for i := range M {
		M[i] = make([]byte, sz)
	}

	ncorr := 1024
	if len(lines) < 30 {
		ncorr = 12
	}

	for i := range ncorr {
		v := Getints(lines[i], false)
		M[v[1]][v[0]] = 1
	}

	showmap()

	d := findpath()
	Sol(d)

	for i := ncorr; i <= len(lines); i++ {
		v := Getints(lines[i], false)
		M[v[1]][v[0]] = 1
		if findpath() < 0 {
			Sol(lines[i])
			break
		}
	}
}

func showmap() {
	for i := range M {
		for j := range M[i] {
			if M[i][j] == 0 {
				Pf(".")
			} else {
				Pf("#")
			}
		}
		Pln()
	}
}

func findpath() int {
	djk := NewDijkstra[pos](pos{0, 0})
	var cur pos
	for djk.PopTo(&cur) {
		if cur.i == len(M)-1 && cur.j == len(M[cur.i])-1 {
			return djk.Dist[cur]
		}

		add := func(i, j int) {
			if i < 0 || i >= len(M) || j < 0 || j >= len(M[i]) || M[i][j] != 0 {
				return
			}
			djk.Add(cur, pos{i, j}, 1)
		}

		add(cur.i+1, cur.j)
		add(cur.i-1, cur.j)
		add(cur.i, cur.j+1)
		add(cur.i, cur.j-1)
	}

	return -1
}
