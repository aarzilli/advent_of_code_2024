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

type pos struct {
	i, j int
}

var ants []pos

var M [][]byte

func showmap(anodes Set[pos]) {
	for i := range M {
		for j := range M[i] {
			if M[i][j] != '.' {
				pf("%c", M[i][j])
			} else if anodes[pos{i,j}] {
				pf("#")
			} else {
				pf(".")
			}
		}
		pln()
	}
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	M = make([][]byte, len(lines))
	for i := range M {
		M[i] = []byte(lines[i])
		for j := range M {
			if M[i][j] != '.' {
				ants  = append(ants, pos{ i, j })
			}
		}
	}
	
	for i := range M {
		pln(string(M[i]))
	}
	pln()
	
	anodes := make(Set[pos])
	anodes2 := make(Set[pos])
	
	for _, a := range ants {
		for _, b := range ants {
			if M[a.i][a.j] != M[b.i][b.j] || a == b {
				continue
			}
			
			di := b.i - a.i
			dj := b.j - a.j
			c := pos{ b.i + di, b.j + dj }
			if c.i >= 0 && c.i < len(M) && c.j >= 0 && c.j < len(M[c.i]) {
				anodes[c] = true
			}
			
			for k := 0; true; k++ {
				c := pos{ b.i + k*di, b.j + k*dj }
				if c.i < 0 || c.i >= len(M) || c.j < 0 || c.j >= len(M[c.i]) {
					break
				}
				
				anodes2[c] = true
			}
		}
	}
	
	showmap(anodes)
	
	Sol(len(anodes))
	
	showmap(anodes2)
	
	Sol(len(anodes2))
}
