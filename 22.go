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

	part1 := 0
	for i, x := range input {
		for range 2000 {
			x = next(x)
		}
		_ = i
		//Pln(input[i], x)
		part1 += x
	}
	Sol(part1)

	Pln()
	example := 123
	vexample := []int{example % 10}
	for i := 0; i < 10; i++ {
		//Pln(example, example%10)
		example = next(example)
		vexample = append(vexample, example%10)
	}

	examplem := make(map[[4]int]int)
	allseqs(vexample, examplem)
	Pln(examplem)

	//allseqs(vexample, make(Set[[4]int]))

	//Pln(score(vexample, []int{ -1, -1, 0, 2 }))

	vv := make([][]int, len(input))

	for i, x := range input {
		vv[i] = append(vv[i], x%10)
		for range 2000 {
			x = next(x)
			vv[i] = append(vv[i], x%10)
		}
	}

	//Pln("part 2 example", score(vv[3], []int{ -2, 1, -1, 3 }))
	//Pln(scoreall(vv, []int{ -2, 1, -1, 3 }))

	//Pln(check(vv[0]))

	//m := make(Set[[4]int])
	//allseqs(vv[0], m)

	/*
		max := 0
		for a := -9; a <= 9; a++ {
			for b := -9; b <= 9; b++ {
				for c := -9; c <= 9; c++ {
					for d := -9; d <= 9; d++ {
						s := scoreall(vv, []int{ a, b, c, d })
						Pln(a, b, c, d, "give", s)
						if s > max {
							max = s
						}
					}
				}
			}
		}
		Sol(max)*/

	seqs := make(map[[4]int]int)
	for i := range vv {
		if i%100 == 0 {
			Pln("calculating all seqs", i)
		}
		allseqs(vv[i], seqs)
	}

	max := 0
	for seq := range seqs {
		if seqs[seq] > max {
			max = seqs[seq]
		}
	}
	Sol(max)

	/*
		max := 0
		cnt := 0
		for seq := range seqs {
			if cnt % 100 == 0 {
				Pln("scoring", cnt)
			}
			cnt++
			s := scoreall(vv, seq[:])
			if s > max {
				max = s
			}
		}
		Sol(max)*/
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

func check(v []int) (int, int) {
	min := v[0]
	max := v[0]

	for i := 1; i < len(v); i++ {
		d := v[i] - v[i-1]
		if d < min {
			min = d
		}
		if d > max {
			max = d
		}
	}
	return min, max
}

func score(v []int, seq []int) int {
	for i := 1; i < len(v); i++ {
		ok := true
		//Pln("at", i, v[i] - v[i-1])
		for j := range seq {
			if i+j >= len(v) {
				ok = false
				break
			}
			if v[i+j]-v[i+j-1] != seq[j] {
				ok = false
				break
			}
		}
		if ok {
			return v[i+len(seq)-1]
		}
	}
	return 0
}

func scoreall(vv [][]int, seq []int) int {
	r := 0
	for i := range vv {
		r += score(vv[i], seq)
	}
	return r
}

func allseqs(v []int, m map[[4]int]int) {
	seen := make(Set[[4]int])
	for i := 4; i < len(v); i++ {
		seq := [4]int{}
		for j := 0; j < len(seq); j++ {
			seq[j] = v[i-3+j] - v[i-4+j]
		}
		if !seen[seq] {
			m[seq] += v[i]
		}
		seen[seq] = true
	}
}
