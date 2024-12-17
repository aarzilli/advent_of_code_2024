package main

import (
	. "aoc/util"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	groups := Input(os.Args[1], "\n\n", true)
	initregs := Noempty(Spac(groups[0], "\n", -1))
	regs := make(map[string]int)
	regs["A"] = Getints(initregs[0], false)[0]
	regs["B"] = Getints(initregs[1], false)[0]
	regs["C"] = Getints(initregs[2], false)[0]

	text := Getints(groups[1], false)
	out := run(text, regs)
	sout := []string{}
	for i := range out {
		sout = append(sout, fmt.Sprintf("%d", out[i]))
	}
	Sol(strings.Join(sout, ","))

	enum(text, 0, len(text)-1) // depends on properties of the input file
}

func enum(text []int, asofar int, i int) bool {
	if i < 0 {
		Sol(asofar)
		return true
	}

	tgt := text[i]
	for a := 0; a <= 7; a++ {
		newa := asofar<<3 + a
		regs := map[string]int{"A": newa}
		out := run(text, regs)
		Pln("digit", i, "wanted value", tgt, "asofar", asofar, "testing", a, "output", out)
		if out[0] == tgt {
			if enum(text, newa, i-1) {
				return true
			}
		}
	}
	return false
}

func run(text []int, regs map[string]int) []int {
	pc := 0
	out := []int{}
	for pc < len(text) {
		switch text[pc] {
		case 0: // adv
			regs["A"] = regs["A"] / int(math.Pow(2, float64(combo(text[pc+1], regs))))
		case 1: // bxl
			regs["B"] = regs["B"] ^ text[pc+1]
		case 2: // bst
			regs["B"] = combo(text[pc+1], regs) % 8
		case 3: // jnz
			if regs["A"] != 0 {
				pc = text[pc+1] - 2
			}
		case 4: // bxc
			regs["B"] = regs["B"] ^ regs["C"]
		case 5: // out
			out = append(out, combo(text[pc+1], regs)%8)
		case 6: // bdv
			regs["B"] = regs["A"] / int(math.Pow(2, float64(combo(text[pc+1], regs))))
		case 7: // cdv
			regs["C"] = regs["A"] / int(math.Pow(2, float64(combo(text[pc+1], regs))))
		}
		pc += 2
	}
	return out
}

func combo(arg int, regs map[string]int) int {
	if arg < 4 {
		return arg
	}
	switch arg {
	case 4:
		return regs["A"]
	case 5:
		return regs["B"]
	case 6:
		return regs["C"]
	default:
		panic("blah")
	}
}
