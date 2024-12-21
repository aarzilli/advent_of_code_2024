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

	part1 := 0
	part2 := 0
	for _, line := range lines {
		num := Atoi(line[:len(line)-1])
		l := multilevellen(3, line)
		l2 := multilevellen(26, line)
		part1 += num * l
		part2 += num * l2
	}
	Sol(part1)
	Sol(part2)
}

func multilevellen(difficulty int, s string) int {
	cur := byte('A')
	r := 0
	for i := range s {
		ch := s[i]
		n := multilevellen1('N', difficulty, cur, ch)
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

	startp, ok1 := pad[start]
	endp, ok2 := pad[end]

	if !ok1 || !ok2 {
		panic("blah")
	}

	r = enummoves(pad, startp, endp, []byte{})
	allmovesmemomap[k] = r
	return r
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
