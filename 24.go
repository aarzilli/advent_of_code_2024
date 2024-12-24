package main

import (
	. "aoc/util"
	"fmt"
	"os"
	"sort"
	"strings"
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

	// calculated by hand
	swapnames("z08", "vvr")
	swapnames("bkr", "rnq")
	swapnames("z28", "tfb")
	swapnames("z39", "mqh")

	inand := make([]string, 45)
	inxor := make([]string, 45)

	for i := range 45 {
		inand[i] = findnodeshape(wirename("x", i), "AND", wirename("y", i)).r
		inxor[i] = findnodeshape(wirename("x", i), "XOR", wirename("y", i)).r
	}

	precarryand := make([]string, 45)
	carry := make([]string, 45)
	carry[0] = inand[0]

	for i := 1; i < 45; i++ {
		// carry[i] e` OR(AND(x[i], y[i]), AND(XOR(x[i], y[i]), carry[i-1]))
		// carry[i] e` OR(inand[i], AND(xor(x[i], y[i]), carry[i-1]))
		// carry[i] e` OR(inand[i], precarryand[i])

		// precarryand[i] e` AND(xor(x[i], y[i]), carry[i-1])
		// precarryand[i] e` AND(inxor[i], carry[i-1])

		precarryand[i] = findnodeshape(inxor[i], "AND", carry[i-1]).r
		carry[i] = findnodeshape(inand[i], "OR", precarryand[i]).r
		Pf("carry[%d] is %s from %s:AND(x%02d, y%02d) and %s:AND(...)\n", i, carry[i], inand[i], i, i, precarryand[i])
	}

	sort.Strings(swapped)
	Sol(strings.Join(swapped, ","))
}

var swapped = []string{}

func swapnames(a, b string) {
	swapped = append(swapped, a, b)
	for _, n := range G {
		if n.r == a {
			n.r = b
		} else if n.r == b {
			n.r = a
		}
	}
}

func wirename(pfx string, i int) string {
	return fmt.Sprintf("%s%02d", pfx, i)
}

func findnodeshape(a1, op, a2 string) *node {
	var r *node
	cand := func(n *node) {
		if r != nil {
			panic(fmt.Errorf("duplicate noe '%s %s %s", a1, op, a2))
		}
		r = n
	}
	for _, n := range G {
		if n.a1 == a1 && n.op == op && n.a2 == a2 {
			cand(n)
		}
		if n.a2 == a1 && n.op == op && n.a1 == a2 {
			cand(n)
		}
	}
	if r != nil {
		return r
	}
	panic(fmt.Errorf("could not find node '%s %s %s'", a1, op, a2))
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
		w, ok := V[fmt.Sprintf("z%02d", i)]
		if !ok {
			return -1
		}
		z += w
	}

	return z
}

func findnode(r string) int {
	for i := range G {
		if G[i].r == r {
			return i
		}
	}
	panic("not found")
}

func hasloop() bool {
	for i := range 46 {
		zname := fmt.Sprintf("z%02d", i)
		if hasloopm(zname, make(Set[string])) {
			return true
		}
	}
	return false
}

func hasloopm(r string, path Set[string]) bool {
	if r[0] == 'x' || r[0] == 'y' {
		return false
	}
	if path[r] {
		return true
	}

	path[r] = true

	n := G[findnode(r)]
	if hasloopm(n.a1, path) {
		return true
	}
	if hasloopm(n.a2, path) {
		return true
	}
	delete(path, r)
	return false
}
