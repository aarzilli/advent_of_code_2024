package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

// returns x without the last character
func Nolast(x string) string {
	return x[:len(x)-1]
}

// splits a string, trims spaces on every element
func Spac(in, sep string, n int) []string {
	v := strings.SplitN(in, sep, n)
	for i := range v {
		v[i] = strings.TrimSpace(v[i])
	}
	return v
}

// convert string to integer
func Atoi(in string) int {
	n, err := strconv.Atoi(in)
	Must(err)
	return n
}

// convert vector of strings to integer
func Vatoi(in []string) []int {
	r := make([]int, len(in))
	for i := range in {
		var err error
		r[i], err = strconv.Atoi(in[i])
		Must(err)
	}
	return r
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Exit(n int) {
	os.Exit(n)
}

func Pf(fmtstr string, args ...interface{}) {
	fmt.Printf(fmtstr, args...)
}

func Getints(in string, hasneg bool) []int {
	v := Getnums(in, hasneg, false)
	return Vatoi(v)
}

func Getnums(in string, hasneg, hasdot bool) []string {
	r := []string{}
	start := -1

	flush := func(end int) {
		if start < 0 {
			return
		}
		hasdigit := false
		for i := start; i < end; i++ {
			if in[i] >= '0' && in[i] <= '9' {
				hasdigit = true
				break
			}
		}
		if hasdigit {
			r = append(r, in[start:end])
		}
		start = -1
	}

	for i, ch := range in {
		isnumch := false

		switch {
		case hasneg && (ch == '-'):
			isnumch = true
		case hasdot && (ch == '.'):
			isnumch = true
		case ch >= '0' && ch <= '9':
			isnumch = true
		}

		if start >= 0 {
			if !isnumch {
				flush(i)
			}
		} else {
			if isnumch {
				start = i
			}
		}
	}
	flush(len(in))
	return r
}

// removes empty string elements, modifies v
func Noempty(v []string) []string {
	r := v[:0]
	for _, s := range v {
		if s != "" {
			r = append(r, s)
		}
	}
	return r
}

func Max(in []int) int {
	max := in[0]
	for i := range in {
		if in[i] > max {
			max = in[i]
		}
	}
	return max
}

func Min(in []int) int {
	min := in[0]
	for i := range in {
		if in[i] < min {
			min = in[i]
		}
	}
	return min
}

func Input(path string, sep string, noempty bool) []string {
	buf, err := ioutil.ReadFile(path)
	Must(err)
	lines := Spac(string(buf), sep, -1)
	if noempty {
		lines = Noempty(lines)
	}
	return lines
}

var part int = 1
var expected []interface{}

func Expect(v ...interface{}) {
	// for refactoring
	expected = v
}

func Sol(v ...interface{}) {
	fmt.Printf("PART %d: ", part)
	fmt.Println(v...)
	if expected != nil {
		if expected[len(expected)-1] != v[len(v)-1] {
			panic("mismatch!")
		}
		expected = nil
	}
	fmt.Printf("copied to clipboard\n")
	cmd := exec.Command("xclip", "-in", "-selection", "-primary")
	cmd.Stdin = bytes.NewReader([]byte(fmt.Sprintf("%v", v[len(v)-1])))
	cmd.Run()
	part++
}

func Sort[T any](v []T, less func(T, T) bool) {
	sort.Slice(v, func(i, j int) bool { return less(v[i], v[j]) })
}

func Filter[T any](p func(T) bool, v []T) []T {
	r := []T{}
	for i := range v {
		if p(v[i]) {
			r = append(r, v[i])
		}
	}
	return r
}

func Neighbors4(p, max [2]int) [][2]int {
	r := [][2]int{}
	for _, delta := range [][2]int{{-1, 0}, {0, -1}, {0, +1}, {+1, 0}} {
		p2 := [2]int{p[0] + delta[0], p[1] + delta[1]}
		if p2[0] < 0 || p2[1] < 0 || p2[0] >= max[0] || p2[1] >= max[1] {
			continue
		}
		r = append(r, p2)
	}
	return r
}

func OneKey[K comparable, V any](m map[K]V) K {
	for k := range m {
		return k
	}
	panic("blah")
}

func Keys[K comparable, V any](m map[K]V) []K {
	r := []K{}
	for k := range m {
		r = append(r, k)
	}
	return r
}

func Sum(v []int) int {
	tot := 0
	for i := range v {
		tot += v[i]
	}
	return tot
}

func Histo[T comparable](v []T) map[T]int {
	m := make(map[T]int)
	for i := range v {
		m[v[i]]++
	}
	return m
}

func Intersect[K comparable, V1, V2 any](m1 map[K]V1, m2 map[K]V2) map[K]bool {
	r := make(map[K]bool)
	for k := range m1 {
		if _, ok := m2[k]; ok {
			r[k] = true
		}
	}
	return r
}

func Union[K comparable, V1, V2 any](m1 map[K]V1, m2 map[K]V2) map[K]bool {
	r := make(map[K]bool)
	for k := range m1 {
		r[k] = true
	}
	for k := range m2 {
		r[k] = true
	}
	return r
}

func Reverse[T any](v []T) {
	for i := 0; i < len(v)/2; i++ {
		v[i], v[len(v)-i-1] = v[len(v)-i-1], v[i]
	}
}

type Set[T comparable] map[T]bool

func (s Set[T]) Add(x T) {
	s[x] = true
}

func (s Set[T]) AddSet(s2 Set[T]) {
	for x := range s2 {
		s[x] = true
	}
}
