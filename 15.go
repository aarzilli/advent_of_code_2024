package main

import (
	. "aoc/util"
	"os"
	"strings"
)

type pos struct {
	i, j int
}

var M [][]byte

func main() {
	groups := Input(os.Args[1], "\n\n", true)
	Pf("len %d\n", len(groups))

	lines := Noempty(Spac(groups[0], "\n", -1))

	var cur pos

	reinit1(&cur, lines)

	instrs := []byte(strings.ReplaceAll(groups[1], "\n", ""))
	cur = run(instrs, cur, false)

	//showmap(cur)

	Sol(score())

	reinit2(&cur, lines)

	showmap(cur)
	cur = run(instrs, cur, true)
	showmap(cur)
	Sol(score())
}

func reinit1(cur *pos, lines []string) {
	M = make([][]byte, len(lines))
	for i := range lines {
		M[i] = []byte(lines[i])
		for j := range M[i] {
			if M[i][j] == '@' {
				*cur = pos{i, j}
				M[i][j] = '.'
			}
		}
	}
}

func reinit2(cur *pos, lines []string) {
	M = make([][]byte, len(lines))
	for i := range lines {
		M[i] = make([]byte, 2*len(lines[i]))
		for j := range lines[i] {
			ch := lines[i][j]
			if ch == '@' {
				*cur = pos{i, 2 * j}
				ch = '.'
			}
			if ch == 'O' {
				M[i][2*j] = '['
				M[i][2*j+1] = ']'
			} else {
				M[i][2*j] = ch
				M[i][2*j+1] = ch
			}
		}
	}
}

func run(instrs []byte, cur pos, part2 bool) pos {
	for _, instr := range instrs {
		var dir pos
		switch instr {
		case '<':
			dir = pos{0, -1}
		case '>':
			dir = pos{0, +1}
		case '^':
			dir = pos{-1, 0}
		case 'v':
			dir = pos{+1, 0}
		default:
			panic("blah")
		}

		next := addpos(cur, dir)

		if !part2 {
			if pushto(next, dir, '.') {
				cur = next
			}
		} else {
			if pushto1(next, dir) {
				cur = next
			}
			//showmap(cur)
		}
	}
	return cur
}

func addpos(a, b pos) pos {
	return pos{a.i + b.i, a.j + b.j}
}

func pushto(p, dir pos, item byte) bool {
	switch M[p.i][p.j] {
	case '#':
		return false
	case '.':
		M[p.i][p.j] = item
		return true
	case 'O':
		r := pushto(addpos(p, dir), dir, 'O')
		if r {
			M[p.i][p.j] = item
		}
		return r
	default:
		panic("blah")
	}
}

func pushto1(p, dir pos) bool {
	switch M[p.i][p.j] {
	case '#':
		return false
	case '.':
		return true
	default:
		if box := isbox(p); box != nil {
			needtomove := make(Set[[2]pos])
			needtomove[*box] = true
			r := pushto2(addbox(*box, dir), dir, needtomove, *box)
			if r {
				//Pln("moving shit")
				for len(needtomove) > 0 {
					for p := range needtomove {
						n1, n2 := addpos(p[0], dir), addpos(p[1], dir)
						if dir == (pos{+1, 0}) || dir == (pos{-1, 0}) {
							if isempty(n1) && isempty(n2) {
								M[p[0].i][p[0].j] = '.'
								M[p[1].i][p[1].j] = '.'
								M[n1.i][n1.j] = '['
								M[n2.i][n2.j] = ']'
								delete(needtomove, p)
							}
						} else if dir == (pos{0, -1}) {
							if isempty(n1) {
								M[p[0].i][p[0].j] = '.'
								M[p[1].i][p[1].j] = '.'
								M[n1.i][n1.j] = '['
								M[n2.i][n2.j] = ']'
								delete(needtomove, p)
							}
						} else if dir == (pos{0, +1}) {
							if isempty(n2) {
								M[p[0].i][p[0].j] = '.'
								M[p[1].i][p[1].j] = '.'
								M[n1.i][n1.j] = '['
								M[n2.i][n2.j] = ']'
								delete(needtomove, p)
							}
						} else {
							panic("blah")
						}
					}
				}
				//Pln("done")
			}
			return r
		} else {
			panic("blah")
		}
	}
}

func isempty(p pos) bool {
	return M[p.i][p.j] == '.'
}

func pushto2(dst [2]pos, dir pos, needtomove Set[[2]pos], curbox [2]pos) bool {
	//Pln(dst, dir)
	p1 := dst[0]
	p2 := dst[1]
	if M[p1.i][p1.j] == '#' || M[p2.i][p2.j] == '#' {
		return false
	}

	if M[p1.i][p1.j] == '.' && M[p2.i][p2.j] == '.' {
		return true
	}

	var r1, r2 bool

	if box := isbox(p1); box != nil && *box != curbox {
		//Pln("moving first", box)
		needtomove[*box] = true
		r1 = pushto2(addbox(*box, dir), dir, needtomove, *box)
	} else {
		r1 = true
	}

	if box := isbox(p2); box != nil && *box != curbox {
		//Pln("moving second", box)
		needtomove[*box] = true
		r2 = pushto2(addbox(*box, dir), dir, needtomove, *box)
	} else {
		r2 = true
	}

	return r1 && r2
}

func addbox(box [2]pos, dir pos) [2]pos {
	return [2]pos{addpos(box[0], dir), addpos(box[1], dir)}
}

func showmap(cur pos) {
	for i := range M {
		for j := range M[i] {
			if (pos{i, j}) == cur {
				Pf("@")
			} else {
				Pf("%c", M[i][j])
			}
		}
		Pln()
	}
	Pln()
}

func score() int {
	r := 0
	for i := range M {
		for j := range M[i] {
			if M[i][j] == 'O' || M[i][j] == '[' {
				r += 100*i + j
			}
		}
	}
	return r
}

func isbox(p pos) *[2]pos {
	switch M[p.i][p.j] {
	case '.':
		return nil
	case '#':
		return nil
	case '[':
		return &[2]pos{p, pos{p.i, p.j + 1}}
	case ']':
		return &[2]pos{pos{p.i, p.j - 1}, p}
	default:
		panic("blah")
	}
}
