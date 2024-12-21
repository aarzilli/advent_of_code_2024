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
		//l2 := multilevellenold(line)
		Pln(num, l)
		part1 += num * l
	}
	Sol(part1)
}

const DIFFICULTY = 26
//const DIFFICULTY = 3

// part1 example: 126384
// part1 real: 270084

// part2 wrong: 154115708116294 (too low)
// part2 wrong: 61952932092390 (too low)

func multilevellenold(s string) int {
	combos := multilevel(s)
	min := -1
	mincombo := ""
	for k := range combos {
		if min < 0 || len(k) < min {
			min = len(k)
			mincombo = k
		}
	}
	Pln("multilevellenold", s, mincombo)
	return min
}

func multilevellen(s string) int {
	cur := byte('A')
	r := 0
	for i := range s {
		ch := s[i]
		n := multilevellen1('N', DIFFICULTY, cur, ch)
		//Pln(string(cur), string(ch), n)
		r += n
		cur = ch
	}
	return r
}

var multilevellen1memomap = make(map[string]int)

func multilevellen1(padkind byte, n int, start, end byte) int {
	k := fmt.Sprintf("%d:%c%c%c", n, padkind, start, end)
	r, ok := multilevellen1memomap[k]
	if ok {
		return r
	}
	r = multilevellen1real(padkind, n, start, end)
	//Pln("multilevellen1real", string(padkind), n, "from:", string(start), "to:", string(end), "minsteps:", r)
	multilevellen1memomap[k] = r
	return r
}

func multilevellen1real(padkind byte, n int, start, end byte) int {
	if n == 0 {
		return 1
	}
	
	moves := allmoves(padkind, start, end)
	
	min := -1
	for _, move := range moves {
		cur := byte('A')
		l := 0
		for i := range move {
			ch := move[i]
			l += multilevellen1('A', n-1, cur, ch)
			cur = ch
		}
		if min < 0 || l < min {
			min = l
		}
	}
	return min
}

/*func multilevellen1(start, end byte) int {
	k := fmt.Sprintf("%c%c", start, end)
	r, ok := multilevellen1memomap[k]
	if ok {
		return r
	}
	r = multilevellen1real(start, end)
	multilevellen1memomap[k] = r
	return r
}

func multilevellen1real(start, end byte) int {
	combos := make(Set[string])
	combos[string(end)] = true
	for i := 0; i < 3; i++ {
		ncombos := make(Set[string])
		padkind := byte('N')
		s := start
		if i > 0 {
			padkind = 'A'
			s = 'A'
		}
		for combo := range combos {
			r := codesfor(padkind, s, combo)
			for _, p := range r {
				ncombos[p] = true
			}
		}
		combos = ncombos
	}
	min := -1
	mincombo := ""
	for k := range combos {
		if min < 0 || len(k) < min {
			min = len(k)
			mincombo = k
		}
	}
	Pln("multilevellen1real", string(start), string(end), mincombo, min)
	return min
}*/

func multilevel(s string) Set[string] {
	curcombos := make(Set[string])
	curcombos[s] = true

	for i := 0; i < DIFFICULTY; i++ {
		nextcombos := make(Set[string])
		padkind := byte('N')
		if i > 0 {
			padkind = 'A'
		}
		for combo := range curcombos {
			r := codesfor(padkind, 'A', combo)
			for _, p := range r {
				nextcombos[p] = true
			}
		}
		curcombos = nextcombos
	}
	return curcombos
}

var codesformemomap = make(map[string][]string)

func codesfor(padkind, start byte, s string) []string {
	k := fmt.Sprintf("%c%c%s", padkind, start, s)
	r, ok := codesformemomap[k]
	if ok {
		return r
	}
	r = codesforreal(padkind, start, s)
	codesformemomap[k] = r
	return r
}

func codesforreal(padkind, start byte, s string) []string {
	cur := start
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
