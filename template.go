package main

import (
	. "aoc/util"
	"os"
)

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))
}
