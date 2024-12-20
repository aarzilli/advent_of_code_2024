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

	part1 := make(map[int]Set[cheat])

	for i := range M {
		for j := range M[i] {
			p := pos{i, j}
			toend, ok := djk.Dist[p]
			if !ok {
				continue
			}

			for _, s1v := range Neighbors4(p.tov(), Dims) {
				s1 := posfromv(s1v)
				if M[s1.i][s1.j] != '#' {
					continue
				}

				for _, s2v := range Neighbors4(s1.tov(), Dims) {
					s2 := posfromv(s2v)
					if M[s2.i][s2.j] != '.' {
						continue
					}

					newtoend, ok := djk.Dist[s2]
					if !ok {
						continue
					}

					if newtoend+2 > toend {
						continue
					}

					save := toend - newtoend - 2

					if save == 0 {
						continue
					}

					if part1[save] == nil {
						part1[save] = make(Set[cheat])
					}

					part1[save][cheat{s1, s2}] = true
				}
			}
		}
	}

	cnt := 0
	for k := range part1 {
		if k >= 100 {
			cnt += len(part1[k])
		}
	}
	Sol(cnt)

	/*
		D := 20
		part2 := make(map[int]Set[cheat])
		for i := range M {
			for j := range M[i] {
				p := pos{i, j}
				toend, ok := djk.Dist[p]
				if !ok {
					continue
				}

				_ = toend

				for s2, steps := range enum(p, D) {
					newtoend, ok := djk.Dist[s2]
					if !ok {
						continue
					}

					if steps > 20 {
						panic("too long")
					}

					if M[s2.i][s2.j] == '#' {
						panic("into a wall")
					}

					///Pln("from", p, "to", s2, "in", steps, "steps")

					saved := toend - newtoend - steps

					if saved <= 0 {
						continue
					}

					if part2[saved] == nil {
						part2[saved] = make(Set[cheat])
					}

					if saved == 74 {
						Pln("shortcut start", p, "shortcut end", s2, "old dist", toend, "new dist", newtoend, "number of steps in the walls", steps, "saves", saved)
					}

					part2[saved][cheat{p, s2}] = true
				}
			}
		}
	*/

	part2 := make(map[int]Set[cheat])
	for i := range M {
		for j := range M[i] {
			p := pos{i, j}
			toend, ok := djk.Dist[p]
			if !ok {
				continue
			}

			for i := range M {
				for j := range M[i] {
					p2 := pos{i, j}

					if p == p2 {
						continue
					}

					steps := dist(p, p2)
					if steps > 20 {
						continue
					}

					newtoend, ok := djk.Dist[p2]
					if !ok {
						continue
					}

					saved := toend - newtoend - steps

					if saved <= 0 {
						continue
					}

					if part2[saved] == nil {
						part2[saved] = make(Set[cheat])
					}

					if saved == 74 {
						Pln("shortcut start", p, "shortcut end", p2, "old dist", toend, "new dist", newtoend, "number of steps in the walls", steps, "saves", saved)
					}

					part2[saved][cheat{p, p2}] = true
				}
			}
		}
	}
	for _, k := range slices.Sorted(maps.Keys(part2)) {
		Pln("there is", len(part2[k]), "cheat that saves", k, "seconds")
		if len(part2[k]) < 10 {
			Pln(part2[k])
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

// 261453

type cheat struct {
	s1, s2 pos
}

func (p pos) tov() [2]int {
	return [2]int{p.i, p.j}
}

func posfromv(v [2]int) pos {
	return pos{v[0], v[1]}
}

func dist(p1, p2 pos) int {
	return Abs(p1.i-p2.i) + Abs(p1.j-p2.j)
}

/*
func enum(p pos, n int) func(yield func(pos, int) bool) {
	return func(yield func(pos, int) bool) {
		djk := NewDijkstra[pos](p)
		var cur pos
		for djk.PopTo(&cur) {
			add := func(i, j int) {
				if i < 0 || i >= len(M) || j < 0 || j >= len(M[i]) {
					return
				}

				djk.Add(cur, pos{i, j}, 1)
			}
			if M[cur.i][cur.j] == '.' {
				if djk.Dist[cur] <= n {
					yield(cur, djk.Dist[cur])
				}
				if djk.Dist[cur] == 0 {
					add(cur.i+1, cur.j)
					add(cur.i-1, cur.j)
					add(cur.i, cur.j+1)
					add(cur.i, cur.j-1)
				}
			} else if djk.Dist[cur] <= n {
				add(cur.i+1, cur.j)
				add(cur.i-1, cur.j)
				add(cur.i, cur.j+1)
				add(cur.i, cur.j-1)
			}
		}
	}
}*/

func enum(p pos, D int) func(yield func(pos, int) bool) {
	return func(yield func(pos, int) bool) {
		enumtrue(p, D, 0, make(Set[pos]), yield)
	}
}

func enumtrue(p pos, D, n int, inpath Set[pos], yield func(pos, int) bool) {
	if n >= D {
		return
	}

	for _, p2v := range Neighbors4(p.tov(), Dims) {
		p2 := posfromv(p2v)
		if inpath[p2] {
			continue
		}

		if M[p2.i][p2.j] == '.' {
			yield(p2, n+1)
		} else {
			inpath[p2] = true
			enumtrue(p2, D, n+1, inpath, yield)
			delete(inpath, p2)
		}
	}
}
