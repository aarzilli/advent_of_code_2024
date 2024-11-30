package util

import (
	"container/heap"
)

// see ../../aoc2022/12.go

/*

Usage:

djk := NewDijkstra[state](start)

var cur state
for djk.PopTo(&cur) {
	if isTarget(cur) {
		// target found
	}

	djk.Add(neighbor1)
	djk.Add(Neighbor2)
	...
}
*/

type Dijkstra[T comparable] struct {
	h      heapInternal[T]
	Dist   map[T]int
	Parent map[T]T
}

func NewDijkstra[T comparable](start T) *Dijkstra[T] {
	djk := &Dijkstra[T]{
		h: heapInternal[T]{
			m: make(map[T]int),
		},
		Dist:   make(map[T]int),
		Parent: make(map[T]T),
	}
	heap.Init(&djk.h)
	djk.Add(start, start, 0)
	return djk
}

// Len returns the size of the fringe.
func (djk *Dijkstra[T]) Len() int {
	return len(djk.h.v)
}

// PopTo extracts the nearest state in the fringe.
func (djk *Dijkstra[T]) PopTo(n *T) bool {
	if len(djk.h.v) <= 0 {
		return false
	}
	node := heap.Pop(&djk.h).(heapNode[T])
	djk.Dist[node.s] = node.dist
	*n = node.s
	return true
}

func (djk *Dijkstra[T]) Pop() T {
	var r T
	if !djk.PopTo(&r) {
		panic("empty")
	}
	return r
}

func (djk *Dijkstra[T]) Seen(n T) bool {
	_, ok := djk.Dist[n]
	return ok
}

// Add adds link from cur to nb, with distance dist (dist is the distance
// between cur and nb).
// Returns true if this is the shortest distance between start and nb that
// we have so far.
func (djk *Dijkstra[T]) Add(cur, nb T, dist int) (bool, int) {
	if djk.Seen(nb) {
		return false, djk.Dist[nb]
	}
	curdist := djk.Dist[cur]
	cd, ok := djk.h.m[nb]
	if !ok || curdist+dist < cd {
		// TODO: remove old version of nb from the heap
		heap.Push(&djk.h, heapNode[T]{nb, curdist + dist})
		djk.Parent[nb] = cur
		return true, curdist + dist
	}
	return false, cd
}

func (djk *Dijkstra[T]) PathTo(n T) []T {
	r := []T{n}
	for {
		parent, ok := djk.Parent[n]
		if !ok || parent == n {
			Reverse(r)
			return r
		}
		r = append(r, parent)
		n = parent
	}
}

type heapNode[T comparable] struct {
	s    T
	dist int
}

type heapInternal[T comparable] struct {
	v []heapNode[T]
	m map[T]int
}

func (h *heapInternal[T]) Push(x any) {
	el := x.(heapNode[T])
	h.v = append(h.v, el)
	h.m[el.s] = el.dist
}

func (h *heapInternal[T]) Pop() any {
	r := h.v[len(h.v)-1]
	h.v = h.v[:len(h.v)-1]
	delete(h.m, r.s)
	return r
}

func (h *heapInternal[T]) Len() int {
	return len(h.v)
}

func (h *heapInternal[T]) Less(i, j int) bool {
	return h.v[i].dist < h.v[j].dist
}

func (h *heapInternal[T]) Swap(i, j int) {
	h.v[i], h.v[j] = h.v[j], h.v[i]
}
