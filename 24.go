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

	// 45bit + 45bit -> 46bit

	// 122, 10, 136, 170, 105, 48, 193

	//swap(10, 105) // should be 10, 48 probably (but it could also be others)

	//swap(105, 193)

	//fix()

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

func fix() bool {
	Pln("FIX ENTERED")
	defer Pln("FIX EXITING")
	correct := make(Set[int])

	for curbit := range 45 {
		zname := fmt.Sprintf("z%02d", curbit)
		if !check1z(curbit, zname) {
			Pln("failure at bit", curbit)
			//TODO: print subgraph of z<i> that could be incorrect
			candidates := []int{}
			for i := range subgraph(zname) {
				if !correct[i] {
					candidates = append(candidates, i)
				}
			}

			if recheckok(curbit) {
				panic("FUCK")
			}

			Pln("candidates", candidates)

			for i := range candidates {
				for j := i + 1; j < len(candidates); j++ {
					Pln("swapping", candidates[i], candidates[j])
					swap(candidates[i], candidates[j])
					recheckout := recheckok(curbit)
					hasloopout := hasloop()
					if recheckout && !hasloopout {
						Pln("trying with swapped", candidates[i], candidates[j])
						if fix() {
							return true
						}
					} else {
						Pln("    failed because recheck=", recheckout, "hasloop=", hasloopout)
					}
					swap(candidates[i], candidates[j])
					if recheckok(curbit) {
						panic("FUCK")
					}
				}
			}

			return false
		}

		for i := range subgraph(zname) {
			correct[i] = true
		}
	}

	return true
}

func check1z(i int, zname string) bool {
	V1, _ := runint(1<<i, 0)
	ok1 := V1[zname] == 1
	V2, _ := runint(0, 1<<i)
	ok2 := V2[zname] == 1
	return ok1 && ok2
}

func recheckok(m int) bool {
	for i := 0; i <= m; i++ {
		zname := fmt.Sprintf("z%02d", i)
		if !check1z(i, zname) {
			return false
		}
	}
	return true
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
