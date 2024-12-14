package main

import (
	. "aoc/util"
	"fmt"
	"os"
	_ "time"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

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

func part1(t int) (int, int, int, int) {
	q1, q2, q3, q4 := 0, 0, 0, 0

	for i := range ptcls {
		px, py := ptcls[i].pos(t)

		switch {
		case px < X/2 && py < Y/2:
			q1++
		case px > X/2 && py < Y/2:
			q2++
		case px < X/2 && py > Y/2:
			q3++
		case px > X/2 && py > Y/2:
			q4++
		}
	}

	return q1, q2, q3, q4
}

type pos struct {
	x, y int
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	if len(lines) < 20 {
		X = 11
		Y = 7
	}

	T := 100

	for i := range lines {
		fields := Getints(lines[i], true)
		var p ptcl
		p.startx, p.starty, p.vx, p.vy = fields[0], fields[1], fields[2], fields[3]
		pln(p)
		ptcls = append(ptcls, p)
	}

	q1, q2, q3, q4 := part1(T)
	pln(q1, q2, q3, q4)
	Sol(q1 * q2 * q3 * q4)

	T2 := 7344

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
				pf(".")
			} else {
				pf("%d", c)
			}
		}
		pln()
	}

	cands := []int{}
	for t := 0; t < 10000000; t++ {
		if t%1000 == 0 {
			pln(t)
		}
		m := make(Set[pos])
		for i := range ptcls {
			px, py := ptcls[i].pos(t)
			m[pos{px, py}] = true
		}

		if hastree(m) {
			pln("Candidate", t)
			cands = append(cands, t)
		}
		/*q1, q2, q3, q4 := part1(t)
		if q1 != q2 || q3 != q4 {
			continue
		}
		_ = q3
		_ = q4
		pln("at time", t)
		time.Sleep(100 * time.Millisecond)
		pf("\033[H\033[2J")*/
	}
	pln("Candidates:", cands)
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
