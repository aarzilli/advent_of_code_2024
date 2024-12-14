package main

import (
	. "aoc/util"
	"os"
)

var X = 101
var Y = 103

type ptcl struct {
	startx, starty, vx, vy int
}

var ptcls = []ptcl{}

func (p *ptcl) pos(t int) (int, int) {
	x := (p.startx + p.vx*t) % X
	y := (p.starty + p.vy*t) % Y
	if x < 0 {
		x += X
	}
	if y < 0 {
		y += Y
	}
	return x, y
}

type pos struct {
	x, y int
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	if len(lines) < 20 {
		X = 11
		Y = 7
	}

	T := 100

	for i := range lines {
		fields := Getints(lines[i], true)
		var p ptcl
		p.startx, p.starty, p.vx, p.vy = fields[0], fields[1], fields[2], fields[3]
		ptcls = append(ptcls, p)
	}

	var q [4]int

	for i := range ptcls {
		px, py := ptcls[i].pos(T)

		if px == X/2 || py == Y/2 {
			continue
		}

		q[2*(2*px/X)+2*py/Y]++
	}
	Sol(q[0]*q[1]*q[2]*q[3])

	if len(lines) < 20 {
		return
	}

	var T2 int

	for t := 0; t < 100000; t++ {
		m := make(Set[pos])
		for i := range ptcls {
			px, py := ptcls[i].pos(t)
			m[pos{px, py}] = true
		}

		if hastree(m) {
			Sol(t)
			T2 = t
			break
		}
	}

	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			c := 0
			for i := range ptcls {
				px, py := ptcls[i].pos(T2)
				if px == x && py == y {
					c++
				}
			}
			if c == 0 {
				Pf(".")
			} else {
				Pf("%d", c)
			}
		}
		Pln()
	}
}

func hastree(m Set[pos]) bool {
particleSearch:
	for p := range m {
		for d := 1; d <= 10; d++ {
			if !m[pos{p.x + d, p.y + d}] || !m[pos{p.x - d, p.y + d}] {
				continue particleSearch
			}
		}
		return true
	}
	return false
}
