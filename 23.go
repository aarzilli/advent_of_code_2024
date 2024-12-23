package main

import (
	. "aoc/util"
	"os"
	"sort"
	"strings"
)

var G = make(map[string]Set[string])

func main() {
	lines := Input(os.Args[1], "\n", true)
	Pf("len %d\n", len(lines))
	for _, line := range lines {
		v := Spac(line, "-", -1)
		if G[v[0]] == nil {
			G[v[0]] = make(Set[string])
		}
		G[v[0]][v[1]] = true
		if G[v[1]] == nil {
			G[v[1]] = make(Set[string])
		}
		G[v[1]][v[0]] = true
	}

	//seen := make(Set[string])

	part1 := 0
	for k1 := range G {
		for k2 := range G[k1] {
			if k2 <= k1 {
				continue
			}
			for k3 := range G[k1] {
				if k3 <= k2 {
					continue
				}
				if G[k2][k3] && (k1[0] == 't' || k2[0] == 't' || k3[0] == 't') {
					part1++
					//Pln(k1, k2, k3)
				}
			}
		}
	}
	Sol(part1)

	seen := make(Set[string])
	clique := make(Set[string])
	maxclique := make(Set[string])

	for {
		clique = make(Set[string])
		for k := range G {
			if !seen[k] {
				clique[k] = true
				break
			}
		}

		if len(clique) == 0 {
			break
		}

		Pln("new clique from", clique)

		for k := range G {
			if seen[k] || clique[k] {
				continue
			}
			ok := true
			for k1 := range clique {
				if !G[k][k1] {
					ok = false
					break
				}
			}
			if ok {
				Pln("\tadding", k, "to clique")
				clique[k] = true
			}
		}

		Pln("\tfinal clique", clique)

		if len(clique) > len(maxclique) {
			maxclique = clique
		}

		for k := range clique {
			seen[k] = true
		}
	}

	keys := Keys(maxclique)
	sort.Strings(keys)
	Sol(strings.Join(keys, ","))

	/*
		for k1 := range G {
			if k1[0] != 't' {
				continue
			}
			seen[k1] = true
			for k2 := range G[k1] {
				if seen[k2] {
					continue
				}
				for k3 := range G[k1] {
					if k3 == k2 {
						continue
					}
					if G[k2][k3] {
						//seen[k2] = true
						//seen[k3] = true
						Pln(k1, k2, k3)
					}
				}
			}
		}*/
}
