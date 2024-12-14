package main

import (
	. "aoc/util"
	"fmt"
	"os"
)

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))
}
