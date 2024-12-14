package main

import (
	. "aoc/util"
	"os"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	groups := Input(os.Args[1], "\n\n", true)
	part1(groups, 0)
	part1(groups, 10000000000000)
}

func part1(groups []string, delta int) {
	fname := os.Args[1]+".preproc"
	fh, err := os.Create(fname)
	Must(err)
	
	for i := range groups {
		lines := Spac(groups[i], "\n", -1)
		if len(lines) != 3 {
			Pf("%#v\n", lines)
			panic("blah")
		}
		
		btna := Getints(lines[0], false)
		btnb := Getints(lines[1], false)
		priz := Getints(lines[2], false)
		
		priz[0] += delta
		priz[1] += delta
		
		fmt.Fprintf(fh, "Minimize[{ 3a+b, %da+%db == %d && %da+%db == %d }, {a, b}, Integers]\n",
			btna[0], btnb[0], priz[0], btna[1], btnb[1], priz[1])
	}
	
	Must(fh.Close())
	
	out, err := exec.Command("wolframscript", "13.wls", fname).CombinedOutput()
	Must(err)
	Sol(strings.TrimSpace(string(out)))
	
	Must(os.Remove(fname))
}
