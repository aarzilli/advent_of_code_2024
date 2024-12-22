package main

import (
	. "aoc/util"
	"os"
)

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))
	input := Vatoi(lines)
	//Pln(input)

	vv := make([][]byte, len(input))

	part1 := 0
	for i, x := range input {
		vv[i] = make([]byte, 0, 2001)
		for range 2000 {
			vv[i] = append(vv[i], byte(x%10))
			x = next(x)
		}
		part1 += x
	}
	Sol(part1, 19241711734)

	seqs := make(map[[4]byte]int)
	for i := range vv {
		allseqs(vv[i], seqs)
	}

	max := 0
	for seq := range seqs {
		if seqs[seq] > max {
			max = seqs[seq]
		}
	}
	Sol(max, 2058)
}

func next(n int) int {
	n = prune(mix(n, n*64))
	n = prune(mix(n, n/32))
	return prune(mix(n, n*2048))
}

func mix(a, b int) int {
	return a ^ b
}

func prune(n int) int {
	return n % 16777216
}

func allseqs(v []byte, m map[[4]byte]int) {
	seen := make(Set[[4]byte])
	seq := [4]byte{0, v[1] - v[0], v[2] - v[1], v[3] - v[2]}
	for i := 4; i < len(v); i++ {
		seq[0] = seq[1]
		seq[1] = seq[2]
		seq[2] = seq[3]
		seq[3] = v[i] - v[i-1]
		if !seen[seq] {
			m[seq] += int(v[i])
		}
		seen[seq] = true
	}
}
