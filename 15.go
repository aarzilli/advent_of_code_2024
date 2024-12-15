package main

import (
	. "aoc/util"
	"os"
	"strings"
)

type pos struct {
	i, j int
}

var debug = false

var M [][]byte

func main() {
	if len(os.Args) < 2 {
		for _, t := range []struct {
			fname  string
			p1, p2 int
		}{
			{"15.example", 10092, 9021},
			{"15.example2", 2028, -1},
			{"15.example3", 908, 618},
			{"15.txt", 1485257, 1475512},
		} {
			p1, p2 := solve(t.fname)
			Pln(t.fname, p1, p2)
			if t.p1 != -1 && p1 != t.p1 {
				panic("wrong part 1")
			}
			if t.p2 != -1 && p2 != t.p2 {
				panic("wrong part 2")
			}
		}
		return
	}
	debug = true
	part1, part2 := solve(os.Args[1])
	Sol(part1)
	Sol(part2)
}

func solve(fname string) (int, int) {
	groups := Input(fname, "\n\n", true)

	lines := Noempty(Spac(groups[0], "\n", -1))

	var cur pos

	reinit1(&cur, lines)

	instrs := []byte(strings.ReplaceAll(groups[1], "\n", ""))
	cur = run(instrs, cur, false)

	part1 := score()

	reinit2(&cur, lines)

	cur = run(instrs, cur, true)
	part2 := score()

	return part1, part2
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
		if part2 && debug {
			showmap(cur)
			Pln("moving", string(instr))
		}
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

		if !part2 || instr == '<' || instr == '>' {
			if pushto(next, dir, '.') {
				cur = next
			}
		} else {
			if pushto1(next, dir.i) {
				cur = next
			}
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
	case 'O', '[', ']':
		r := pushto(addpos(p, dir), dir, M[p.i][p.j])
		if r {
			M[p.i][p.j] = item
		}
		return r
	default:
		panic("blah")
	}
}

func pushto1(p pos, dir int) bool {
	switch M[p.i][p.j] {
	case '#':
		return false
	case '.':
		return true
	default:
		if box := isbox(p); box != nil {
			r := pushto2(addbox(*box, dir), dir, false)
			if r {
				fill(*box, false)
				pushto2(addbox(*box, dir), dir, true)
			}
			return r
		} else {
			panic("blah")
		}
	}
}

func fill(box [2]pos, v bool) {
	if v {
		M[box[0].i][box[0].j] = '['
		M[box[1].i][box[1].j] = ']'
	} else {
		M[box[0].i][box[0].j] = '.'
		M[box[1].i][box[1].j] = '.'
	}
}

func pushto2(dst [2]pos, dir int, exec bool) bool {
	if M[dst[0].i][dst[0].j] == '#' || M[dst[1].i][dst[1].j] == '#' {
		return false
	}

	m := make(Set[[2]pos])
	for _, p := range dst {
		if box := isbox(p); box != nil {
			m[*box] = true
		}
	}

	ok := true
	for box := range m {
		if exec {
			fill(box, false)
		}
		r1 := pushto2(addbox(box, dir), dir, exec)
		if !r1 {
			ok = false
			break
		}
	}

	if exec {
		fill(dst, true)
	}

	return ok
}

func addbox(box [2]pos, dir int) [2]pos {
	return [2]pos{addpos(box[0], pos{ dir, 0 }), addpos(box[1], pos{ dir, 0 })}
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
