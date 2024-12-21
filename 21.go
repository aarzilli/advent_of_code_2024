package main

import (
	. "aoc/util"
	"fmt"
	"os"
)

type pos struct {
	i, j int
}

var numpad = map[byte]pos{
	'7': pos{0, 0},
	'8': pos{0, 1},
	'9': pos{0, 2},

	'4': pos{1, 0},
	'5': pos{1, 1},
	'6': pos{1, 2},

	'1': pos{2, 0},
	'2': pos{2, 1},
	'3': pos{2, 2},

	'0': pos{3, 1},
	'A': pos{3, 2},
}

var arrowpad = map[byte]pos{
	'^': pos{0, 1},
	'A': pos{0, 2},
	'<': pos{1, 0},
	'v': pos{1, 1},
	'>': pos{1, 2},
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))
	//Pf("%q\n", allmoves('N', '7', 'A'))
	//Pf("%q\n", codesfor('N', "029A"))
	part1 := 0
	for _, line := range lines {
		num := Atoi(line[:len(line)-1])
		l := multilevellen(line)
		Pln(num, l)
		part1 += num * l
	}
	Sol(part1)
}

// part1 example: 126384
// part1 real: 270084

func multilevellen(s string) int {
	combos := multilevel(s)
	min := -1
	for k := range combos {
		if min < 0 || len(k) < min {
			min = len(k)
		}
	}
	return min
}

func multilevel(s string) Set[string] {
	curcombos := make(Set[string])
	curcombos[s] = true

	for i := 0; i < 3; i++ {
		nextcombos := make(Set[string])
		padkind := byte('N')
		if i > 0 {
			padkind = 'A'
		}
		for combo := range curcombos {
			r := codesfor(padkind, combo)
			for _, p := range r {
				nextcombos[p] = true
			}
		}
		curcombos = nextcombos
	}
	return curcombos
}

var codesformemomap = make(map[string][]string)

func codesfor(padkind byte, s string) []string {
	k := fmt.Sprintf("%c%s", padkind, s)
	r, ok := codesformemomap[k]
	if ok {
		return r
	}
	r = codesforreal(padkind, s)
	codesformemomap[k] = r
	return r
}

func codesforreal(padkind byte, s string) []string {
	cur := byte('A')
	r := [][]string{}
	for i := range s {
		ch := s[i]
		r2 := allmoves(padkind, cur, ch)
		r = append(r, r2)
		cur = ch
	}
	return enumstr(r)
}

func enumstr(r [][]string) []string {
	if len(r) == 0 {
		return []string{""}
	}

	rr := []string{}
	for _, step := range r[0] {
		r2 := enumstr(r[1:])
		for _, p := range r2 {
			rr = append(rr, step+p)
		}
	}
	return rr
}

var allmovesmemomap = make(map[string][]string)

func allmoves(padkind, start, end byte) []string {
	k := fmt.Sprintf("%c%c%c", padkind, start, end)
	r, ok := allmovesmemomap[k]
	if ok {
		return r
	}
	var pad map[byte]pos
	switch padkind {
	case 'N':
		pad = numpad
	case 'A':
		pad = arrowpad
	default:
		panic("blah")
	}
	r = allmovesreal(pad, start, end)
	allmovesmemomap[k] = r
	return r
}

func allmovesreal(pad map[byte]pos, start, end byte) []string {
	startp, ok1 := pad[start]
	endp, ok2 := pad[end]

	if !ok1 || !ok2 {
		panic("blah")
	}

	return enummoves(pad, startp, endp, []byte{})
}

func enummoves(pad map[byte]pos, startp, endp pos, path []byte) []string {
	if startp == endp {
		return []string{string(path) + "A"}
	}

	r := []string{}

	isvalid := func(p pos) bool {
		for _, p2 := range pad {
			if p2 == p {
				return true
			}
		}
		return false
	}

	maybe := func(nextp pos, ch byte) {
		if isvalid(nextp) {
			r2 := enummoves(pad, nextp, endp, append(path, ch))
			r = append(r, r2...)
		}
	}

	if startp.i < endp.i {
		nextp := startp
		nextp.i++
		maybe(nextp, 'v')
	} else if startp.i > endp.i {
		nextp := startp
		nextp.i--
		maybe(nextp, '^')
	}

	if startp.j < endp.j {
		nextp := startp
		nextp.j++
		maybe(nextp, '>')
	} else if startp.j > endp.j {
		nextp := startp
		nextp.j--
		maybe(nextp, '<')
	}

	return r
}
