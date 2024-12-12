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

var M [][]byte
var seen = make(Set[pos])

type pos struct {
	i, j int
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	M = make([][]byte, len(lines))
	for i := range lines {
		M[i] = []byte(lines[i])
	}
	part1 := 0
	for i := range M {
		for j := range M[i] {
			if seen[pos{i, j}] {
				continue
			}

			m := mark(pos{i, j})
			p, _ := perimeter(pos{i, j}, m)
			part1 += len(m) * p
		}
	}
	Sol(part1)

	/*
		fat := make([][]byte, len(M)*3)
		for i := range M {
			fat[3*i] = make([]byte, len(M[i])*3)
			fat[3*i+1] = make([]byte, len(M[i])*3)
			fat[3*i+2] = make([]byte, len(M[i])*3)

			for j := range M[i] {
				fat[3*i][3*j] = M[i][j]
				fat[3*i][3*j+1] = M[i][j]
				fat[3*i][3*j+2] = M[i][j]
				fat[3*i+1][3*j] = M[i][j]
				fat[3*i+1][3*j+1] = M[i][j]
				fat[3*i+1][3*j+2] = M[i][j]
				fat[3*i+2][3*j] = M[i][j]
				fat[3*i+2][3*j+1] = M[i][j]
				fat[3*i+2][3*j+2] = M[i][j]
			}
		}

		M = fat
		clear(seen)

		for i := range M {
			pln(string(M[i]))
		}

		pln()
	*/

	clear(seen)

	pln("BLAH")

	for i := 7; i <= 9; i++ {
		for j := 78; j <= 80; j++ {
			pf("%c", M[i][j])
		}
		pln()
	}

	part2 := 0
	for i := range M {
		for j := range M[i] {
			if seen[pos{i, j}] {
				continue
			}

			pln("marking", string(get(pos{i, j})))
			m := mark(pos{i, j})
			_, pset := perimeter(pos{i, j}, m)
			s := sides(pos{i, j}, get(pos{i, j}), pset, m)
			pln("marked", string(get(pos{i, j})), len(m), s)
			part2 += len(m) * s
		}
	}
	Sol(part2)
}

// 879304 wrong

func mark(start pos) Set[pos] {
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

func get(p pos) byte {
	if p.i < 0 || p.i >= len(M) || p.j < 0 || p.j >= len(M[p.i]) {
		return 0
	}
	return M[p.i][p.j]
}

func perimeter(start pos, m Set[pos]) (int, Set[pos]) {
	ch := M[start.i][start.j]
	r := 0
	pset := make(Set[pos])
	for p := range m {
		//isp := false
		for _, n := range []pos{pos{p.i - 1, p.j}, pos{p.i + 1, p.j}, pos{p.i, p.j - 1}, pos{p.i, p.j + 1}} {
			if get(n) != ch {
				r++
				//isp = true
			}
		}

		/*for _, n := range []pos{ pos{p.i-1, p.j-1}, pos{p.i-1, p.j+1}, pos{p.i+1, p.j-1}, pos{p.i+1, p.j+1}} {
			if get(n) != ch {
				isp = true
			}
		}
		if isp {
			pset[p] = true
		}*/
	}
	return r, pset
}

func isborder(p pos) bool {
	ch := get(p)
	for _, n := range []pos{pos{p.i - 1, p.j}, pos{p.i + 1, p.j}, pos{p.i, p.j - 1}, pos{p.i, p.j + 1}} {
		if get(n) != ch {
			return true
		}
	}
	return false
}

func isborderv(p pos) bool {
	ch := get(p)
	for _, n := range []pos{pos{p.i, p.j - 1}, pos{p.i, p.j + 1}} {
		if get(n) != ch {
			return true
		}
	}
	return false
}

var dirs = []pos{pos{+1, 0}, pos{0, +1}, pos{-1, 0}, pos{0, -1}}

func sides2(start pos, ch byte, pset Set[pos], m Set[pos]) (int, Set[pos]) {
	seen := make(Set[pos])

	dir := 0
	cur := start

	outcur := pos{start.i, start.j - 1}
	if m[outcur] {
		outcur = pos{start.i, start.j + 1}
		if m[outcur] {
			panic("can't find outside")
		}
	}

	pln(pset)

	next := func(dir int) (pos, pos) {
		return pos{cur.i + dirs[dir].i, cur.j + dirs[dir].j}, pos{outcur.i + dirs[dir].i, outcur.j + dirs[dir].j}
	}

	r := 0

	for {
		pln(cur)
		n, outn := next(dir)
		if m[n] && !m[outn] {
			pln("continuing", n, outn)
			cur = n
			outcur = outn
		} else if !m[n] {
			// convex angle
			switch dir {
			case 0:
				if outcur.j < cur.j {
					dir = 1
				} else if outcur.j > cur.j {
					dir = 3
				} else {
					panic("blah")
				}
				outcur = cur
				outcur.i++
			case 1:
				if outcur.i < cur.i {
					dir = 0
				} else if outcur.i > cur.i {
					dir = 2
				} else {
					panic("blah")
				}
				outcur = cur
				outcur.j++
			case 2:
				if outcur.j < cur.j {
					dir = 1
				} else if outcur.j > cur.j {
					dir = 3
				} else {
					panic("blah")
				}
				outcur = cur
				outcur.i--
			case 3:
				if outcur.i < cur.i {
					dir = 0
				} else if outcur.i > cur.i {
					dir = 2
				} else {
					panic("blah")
				}
				outcur = cur
				outcur.j--
			}
			r++
			if m[outcur] {
				panic("blah")
			}

		} else if m[n] {
			// concave angle
			switch dir {
			case 0:
				if outcur.j < cur.j {
					dir = 3
				} else if outcur.j > cur.j {
					dir = 1
				} else {
					panic("blah")
				}
				cur = outcur
				cur.i++
			case 1:
				// (unverified)
				if outcur.i < cur.i {
					dir = 2
				} else if outcur.i > cur.i {
					dir = 0
				} else {
					panic("blah")
				}
				cur = outcur
				cur.j++
			case 2:
				// (unverified)
				if outcur.j < cur.j {
					dir = 3
				} else if outcur.j > cur.j {
					dir = 1
				} else {
					panic("blah")
				}
				cur = outcur
				cur.i--
			case 3:
				// (unverified)
				if outcur.i < cur.i {
					dir = 2
				} else if outcur.i > cur.i {
					dir = 0
				} else {
					panic("blah")
				}
				cur = outcur
				cur.j--
			}
			r++
			if !m[cur] {
				panic("blah")
			}
		} else {
			panic("blah")
		}
		seen[cur] = true
		if cur == start && dir == 0 {
			break
		}
	}
	return r, seen
}

func sides(start pos, ch byte, pset Set[pos], m Set[pos]) int {
	r, seen := sides2(start, ch, pset, m)
	for i := range M {
		for j := range M[i] {
			p := pos{i, j}
			if !m[p] || seen[p] {
				continue
			}
			if !isborderv(p) {
				continue
			}
			pln("restarting from", p)
			r2, seen2 := sides2(p, ch, pset, m)
			for p2 := range seen2 {
				seen[p2] = true
			}
			pln("restart from", p, "got", r2, "sides")
			r += r2
		}
	}
	return r
}
