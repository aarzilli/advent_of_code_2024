package main

import (
	. "aoc/util"
	"fmt"
	"os"
)

func pf(fmtstr string, any ...interface{}) {
	fmt.Printf(fmtstr, any...)
}

func pln(any ...interface{}) {
	fmt.Println(any...)
}

type block struct {
	id         int
	start      int
	size       int
	prev, next *block
}

func main() {
	lines := Input(os.Args[1], "\n", true)
	pf("len %d\n", len(lines))
	v := Vatoi(Spac(lines[0], "", -1))
	//pln(v)
	tot := 0
	for i := range v {
		tot += v[i]
	}
	//pln(tot)
	{
		disk := make([]int, 0, tot)
		for i := range v {
			if i%2 == 0 {
				for j := 0; j < v[i]; j++ {
					disk = append(disk, i/2)
				}
			} else {
				for j := 0; j < v[i]; j++ {
					disk = append(disk, -1)
				}
			}
		}
		//pln(disk)

		free := 0
		for i := range disk {
			if disk[i] == -1 {
				free = i
				break
			}
		}

		for i := len(disk) - 1; i > free; i-- {
			if disk[i] == -1 {
				continue
			}
			disk[free] = disk[i]
			disk[i] = -1
			for j := free; j < len(disk); j++ {
				if disk[j] == -1 {
					free = j
					break
				}
			}
		}

		//pln(disk)

		part1 := 0
		for i := range disk {
			if disk[i] >= 0 {
				part1 += i * disk[i]
			}
		}
		Sol(part1)
	}

	var blocks *block
	var last *block
	id2blocks := []*block{}

	start := 0
	for i := range v {
		b := &block{
			start: start,
			size:  v[i],
		}
		start += v[i]
		if i%2 == 0 {
			b.id = i / 2
			id2blocks = append(id2blocks, b)
			if len(id2blocks)-1 != b.id {
				panic("blah")
			}
		} else {
			b.id = -1
		}
		if last == nil {
			blocks = b
			last = b
		} else {
			b.prev = last
			last.next = b
			last = b
		}
	}

	printdisk(blocks)

	for j := len(id2blocks) - 1; j >= 0; j-- {
		b1 := id2blocks[j]
		pln("trying to move", b1.id)
		b2 := findfree(blocks, b1.start, b1.size)
		if b2 == nil {
			pln("no free space")
			continue
		}

		// replace b1 with free space
		b3 := &block{id: -1, start: b1.start, size: b1.size}
		b3.prev = b1.prev
		b3.next = b1.next
		b1.prev.next = b3
		if b1.next != nil {
			b1.next.prev = b3
		}

		b1.prev = nil
		b1.next = nil

		// insert b1
		b1.next = b2
		b1.prev = b2.prev
		b2.prev.next = b1
		b2.prev = b1

		b1.start = b2.start

		b2.size -= b1.size
		b2.start += b1.size

		for b := blocks; b != nil; b = b.next {
			if b.next != nil && b.id == -1 && b.next.id == -1 {
				b.size += b.next.size
				b.next = b.next.next
				if b.next != nil {
					b.next.prev = b
				}
			}
		}

		//printdisk(blocks)
	}

	printdisk(blocks)

	pos := 0
	part2 := 0
	for b := blocks; b != nil; b = b.next {
		for range b.size {
			if b.id >= 0 {
				part2 += pos * b.id
			}
			pos++
		}
	}
	Sol(part2)
}

func findfree(blocks *block, start int, size int) *block {
	for b := blocks; b != nil; b = b.next {
		if b.id != -1 {
			continue
		}
		if b.size < size {
			continue
		}
		if b.start >= start {
			return nil
		}
		return b
	}
	return nil
}

func printdisk(blocks *block) {
	pos := 0
	for b := blocks; b != nil; b = b.next {
		if b.start != pos {
			panic("blah")
		}
		for range b.size {
			if b.id == -1 {
				pf(".")
			} else {
				pf("%d", b.id)
			}
			pos++
		}
	}
	pln()
}
