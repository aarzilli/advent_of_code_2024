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

const part2 = true

func main() {
	groups := Input(os.Args[1], "\n\n", true)
	
	for i := range groups {
		lines := Spac(groups[i], "\n", -1)
		if len(lines) != 3 {
			pf("%#v\n", lines)
			panic("blah")
		}
		
		btna := Getints(lines[0], false)
		btnb := Getints(lines[1], false)
		priz := Getints(lines[2], false)
		
		if part2 {
			priz[0] += 10000000000000
			priz[1] += 10000000000000
		}
		
		pf("Minimize[{ 3a+b, %da+%db == %d && %da+%db == %d }, {a, b}, Integers]\n",
			btna[0], btnb[0], priz[0], btna[1], btnb[1], priz[1])
	}
}
