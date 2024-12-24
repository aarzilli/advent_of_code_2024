package main

import (
	. "aoc/util"
	"fmt"
	"os"
)

type node struct {
	op     string
	a1, a2 string
	r      string
}

var G = make([]*node, 0)

const writedot = false

func main() {
	groups := Input(os.Args[1], "\n\n", true)
	Pf("len %d\n", len(groups))

	var V = make(map[string]int)

	for _, line := range Noempty(Spac(groups[0], "\n", -1)) {
		v := Noempty(Spac(line, ":", -1))
		V[v[0]] = Atoi(v[1])
	}

	for _, line := range Noempty(Spac(groups[1], "\n", -1)) {
		v := Noempty(Spac(line, " ", -1))
		G = append(G, &node{
			op: v[1],
			a1: v[0],
			a2: v[2],
			r:  v[4],
		})
	}

	n := run(V)
	Sol(n)

	if writedot {
		fh, err := os.Create("24.dot")
		Must(err)
		fmt.Fprintf(fh, "digraph problem {\n")
		for _, n := range G {
			fmt.Fprintf(fh, "%s [label=%q]\n", n.r, n.r+" "+n.op)
			fmt.Fprintf(fh, "%s -> %s\n", n.a1, n.r)
			fmt.Fprintf(fh, "%s -> %s\n", n.a2, n.r)
		}
		fmt.Fprintf(fh, "}\n")
		Must(fh.Close())
	}

	// 45bit + 45bit -> 46bit

	/*

		Pln("subgraph of z00", subgraph2string(subgraph("z00")))

		correct := make(Set[int])

		// 122, 10, 136, 170, 105, 48, 193

		swap(10, 105) // should be 10, 48 probably (but it could also be others)

		for i := range 45 {
			zname := fmt.Sprintf("z%02d", i)
			V1, _ := runint(1<<i, 0)
			ok1 := V1[zname] == 1
			V2, _ := runint(0, 1<<i)
			ok2 := V2[zname] == 1
			if !ok1 || !ok2 {
				Pln("failure at bit", i)
				//TODO: print subgraph of z<i> that could be incorrect
				for i := range subgraph(zname) {
					if !correct[i] {
						Pln("swap candidate", i, G[i])
					}
				}
				break
			}

			for i := range subgraph(zname) {
				correct[i] = true
			}
		}*/
}

func tobool(x int) bool {
	return x != 0
}

func toint(x bool) int {
	if x {
		return 1
	}
	return 0
}

func run(V map[string]int) int {
	for {
		did := false
		for _, n := range G {
			a1n, ok1 := V[n.a1]
			a2n, ok2 := V[n.a2]
			_, ok3 := V[n.r]
			a1 := tobool(a1n)
			a2 := tobool(a2n)
			var r int
			if ok1 && ok2 && !ok3 {
				//Pln("doing node", n.r)
				did = true
				switch n.op {
				case "AND":
					r = toint(a1 && a2)
				case "OR":
					r = toint(a1 || a2)
				case "XOR":
					r = toint(a1 != a2)
				default:
					panic("blah")
				}
				V[n.r] = r
			}
		}
		if !did {
			break
		}
	}

	var z int
	for i := 45; i >= 0; i-- {
		z = z * 2
		z += V[fmt.Sprintf("z%02d", i)]
	}

	return z
}

func runint(x, y int64) (map[string]int, int) {
	V := make(map[string]int)
	towires(V, x, "x")
	towires(V, y, "y")
	return V, run(V)
}

func towires(V map[string]int, n int64, pfx string) {
	for i := range 45 {
		V[fmt.Sprintf("%s%02d", pfx, i)] = int(n % 2)
		n /= 2
	}
}

func findnode(r string) int {
	for i := range G {
		if G[i].r == r {
			return i
		}
	}
	panic("not found")
}

func subgraph(r string) Set[int] {
	m := make(Set[int])
	subgraphm(r, m)
	return m
}

func subgraphm(r string, m Set[int]) {
	if r[0] == 'x' || r[0] == 'y' {
		return
	}
	i := findnode(r)
	m[i] = true
	n := G[i]
	subgraphm(n.a1, m)
	subgraphm(n.a2, m)
}

func subgraph2string(v Set[int]) []string {
	r := []string{}
	for i := range v {
		r = append(r, G[i].String())
	}
	return r
}

func (n *node) String() string {
	return fmt.Sprintf("%s %s %s -> %s", n.a1, n.op, n.a2, n.r)
}

func swap(i, j int) {
	n1 := G[i]
	n2 := G[j]
	n1.r, n2.r = n2.r, n1.r
}
