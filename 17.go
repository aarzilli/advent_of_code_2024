package main

import (
	. "aoc/util"
	"fmt"
	"math"
	"os"
	_ "slices"
	"strings"
)

func main() {
	groups := Input(os.Args[1], "\n\n", true)
	Pf("len %d\n", len(groups))
	regs := getregs(groups[0])

	text := Getints(groups[1], false)
	Pln(text)
	run(text, regs, true)

	if os.Args[1] != "17.txt" && os.Args[1] != "17.example3" {
		Pln(os.Args[1])
		os.Exit(1)
	}

	Pln("input len:", len(text))
	{
		/*a := 37293246
		b := 0
		c := 0
		_ = c
		_ = b
		for a != 0 {
			b := (((a%8)^6)^(a / pow(2, ((a%8)^6)))) ^ (a / pow(2, ((a%8)^6)))
			a = a/8
			Pln(b)
		}*/
	}

	test2 := func(a int) {
		regs["A"] = a
		regs["B"] = 0
		regs["C"] = 0
		out := run(text, regs, false)
		Pln(out)
	}
	_ = test2

	enum(text, 0, len(text)-1)

	/*test1 := func(a int) {
		text := text[:len(text)-2]
		//Pln(text)
		regs["A"] = a
		regs["B"] = 0
		regs["C"] = 0
		out := run(text, regs, false)
		Pf("%03b %v\n", a, out)
	}


	for a := 0; a <= 7; a++ {
		test1(a)
	}

	Pln(text)
	test2(0b111_110_110_110)*/

	/*
		for a := range 1000000 {
			regs := make(map[string]int)
			regs["A"] = a
			regs["B"] = 0
			regs["C"] = 0
			out := run(text, regs, false)
			if slices.Equal(out, text) {
				Sol(a)
				break
			}
			Pln(a, out)
		}*/
}

func enum(text []int, asofar int, i int) bool {
	if i < 0 {
		Sol(asofar)
		return true
	}

	tgt := text[i]
	for a := 0; a <= 7; a++ {
		regs := make(map[string]int)
		regs["A"] = asofar<<3 + a
		regs["B"] = 0
		regs["C"] = 0
		out := run(text, regs, false)
		Pln("digit", i, "wanted value", tgt, "asofar", asofar, "testing", a, "output", out)
		if out[0] == tgt {
			if enum(text, asofar<<3+a, i-1) {
				return true
			}
		}
	}
	return false
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func getregs(s string) map[string]int {
	initregs := Noempty(Spac(s, "\n", -1))
	regs := make(map[string]int)
	regs["A"] = Getints(initregs[0], false)[0]
	regs["B"] = Getints(initregs[1], false)[0]
	regs["C"] = Getints(initregs[2], false)[0]
	return regs
}

func run(text []int, regs map[string]int, part1 bool) []int {
	pc := 0
	out := []int{}
	for pc < len(text) {
		switch text[pc] {
		case 0: // adv
			//Pf("adv %b\n", regs["A"])
			regs["A"] = regs["A"] / int(math.Pow(2, float64(combo(text[pc+1], regs))))
		case 1: // bxl
			regs["B"] = regs["B"] ^ text[pc+1]
		case 2: // bst
			regs["B"] = combo(text[pc+1], regs) % 8
		case 3: // jnz
			//Pln("loop", regs["A"])
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
	sout := []string{}
	for i := range out {
		sout = append(sout, fmt.Sprintf("%d", out[i]))
	}
	if part1 {
		Sol(strings.Join(sout, ","))
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

// 1,5,0,1,7,4,1,0,3
