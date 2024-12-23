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
				clique[k] = true
			}
		}

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
}
