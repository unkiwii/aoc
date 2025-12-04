package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// --- Day 4: Printing Department ---
//
// You ride the escalator down to the printing department. They're clearly
// getting ready for Christmas; they have lots of large rolls of paper
// everywhere, and there's even a massive printer in the corner (to handle the
// really big print jobs).
//
// Decorating here will be easy: they can make their own decorations. What you
// really need is a way to get further into the North Pole base while the
// elevators are offline.
//
// "Actually, maybe we can help with that," one of the Elves replies when you
// ask for help. "We're pretty sure there's a cafeteria on the other side of
// the back wall. If we could break through the wall, you'd be able to keep
// moving. It's too bad all of our forklifts are so busy moving those big rolls
// of paper around."
//
// If you can optimize the work the forklifts are doing, maybe they would have
// time to spare to break through the wall.
//
// The rolls of paper (@) are arranged on a large grid; the Elves even have a
// helpful diagram (your puzzle input) indicating where everything is located.
//
// For example:
//
// ..@@.@@@@.
// @@@.@.@.@@
// @@@@@.@.@@
// @.@@@@..@.
// @@.@@@@.@@
// .@@@@@@@.@
// .@.@.@.@@@
// @.@@@.@@@@
// .@@@@@@@@.
// @.@.@@@.@.
//
// The forklifts can only access a roll of paper if there are fewer than four
// rolls of paper in the eight adjacent positions. If you can figure out which
// rolls of paper the forklifts can access, they'll spend less time looking and
// more time breaking down the wall to the cafeteria.
//
// In this example, there are 13 rolls of paper that can be accessed by a
// forklift (marked with x):
//
// ..xx.xx@x.
// x@@.@.@.@@
// @@@@@.x.@@
// @.@@@@..@.
// x@.@@@@.@x
// .@@@@@@@.@
// .@.@.@.@@@
// x.@@@.@@@@
// .@@@@@@@@.
// x.x.@@@.x.
//
// Consider your complete diagram of the paper roll locations. How many rolls
// of paper can be accessed by a forklift?
func Day4Part1(filename string) int {
	grid := NewGridFromFile(filename)
	return grid.Mark()
}

// --- Part Two ---
//
// Now, the Elves just need help accessing as much of the paper as they can.
//
// Once a roll of paper can be accessed by a forklift, it can be removed. Once
// a roll of paper is removed, the forklifts might be able to access more rolls
// of paper, which they might also be able to remove. How many total rolls of
// paper could the Elves remove if they keep repeating this process?
//
// Starting with the same example as above, here is one way you could remove as
// many rolls of paper as possible, using highlighted @ to indicate that a roll
// of paper is about to be removed, and using x to indicate that a roll of
// paper was just removed:
//
// Initial state:
// ..@@.@@@@.
// @@@.@.@.@@
// @@@@@.@.@@
// @.@@@@..@.
// @@.@@@@.@@
// .@@@@@@@.@
// .@.@.@.@@@
// @.@@@.@@@@
// .@@@@@@@@.
// @.@.@@@.@.
//
// Remove 13 rolls of paper:
// ..xx.xx@x.
// x@@.@.@.@@
// @@@@@.x.@@
// @.@@@@..@.
// x@.@@@@.@x
// .@@@@@@@.@
// .@.@.@.@@@
// x.@@@.@@@@
// .@@@@@@@@.
// x.x.@@@.x.
//
// Remove 12 rolls of paper:
// .......x..
// .@@.x.x.@x
// x@@@@...@@
// x.@@@@..x.
// .@.@@@@.x.
// .x@@@@@@.x
// .x.@.@.@@@
// ..@@@.@@@@
// .x@@@@@@@.
// ....@@@...
//
// Remove 7 rolls of paper:
// ..........
// .x@.....x.
// .@@@@...xx
// ..@@@@....
// .x.@@@@...
// ..@@@@@@..
// ...@.@.@@x
// ..@@@.@@@@
// ..x@@@@@@.
// ....@@@...
//
// Remove 5 rolls of paper:
// ..........
// ..x.......
// .x@@@.....
// ..@@@@....
// ...@@@@...
// ..x@@@@@..
// ...@.@.@@.
// ..x@@.@@@x
// ...@@@@@@.
// ....@@@...
//
// Remove 2 rolls of paper:
// ..........
// ..........
// ..x@@.....
// ..@@@@....
// ...@@@@...
// ...@@@@@..
// ...@.@.@@.
// ...@@.@@@.
// ...@@@@@x.
// ....@@@...
//
// Remove 1 roll of paper:
// ..........
// ..........
// ...@@.....
// ..x@@@....
// ...@@@@...
// ...@@@@@..
// ...@.@.@@.
// ...@@.@@@.
// ...@@@@@..
// ....@@@...
//
// Remove 1 roll of paper:
// ..........
// ..........
// ...x@.....
// ...@@@....
// ...@@@@...
// ...@@@@@..
// ...@.@.@@.
// ...@@.@@@.
// ...@@@@@..
// ....@@@...
//
// Remove 1 roll of paper:
// ..........
// ..........
// ....x.....
// ...@@@....
// ...@@@@...
// ...@@@@@..
// ...@.@.@@.
// ...@@.@@@.
// ...@@@@@..
// ....@@@...
//
// Remove 1 roll of paper:
// ..........
// ..........
// ..........
// ...x@@....
// ...@@@@...
// ...@@@@@..
// ...@.@.@@.
// ...@@.@@@.
// ...@@@@@..
// ....@@@...
//
// Stop once no more rolls of paper are accessible by a forklift. In this
// example, a total of 43 rolls of paper can be removed.
//
// Start with your original diagram. How many rolls of paper in total can be
// removed by the Elves and their forklifts?
func Day4Part2(filename string) int {
	grid := NewGridFromFile(filename)

	step := 0
	result := 0

	for {
		step++
		m := grid.Mark()
		if m == 0 {
			return result
		}
		s := grid.Sweep()
		if m != s {
			log.Fatalf("marked %d rolls, but sweeped %d\n", m, s)
		}

		result += s
	}
}

type Grid [][]Cell

type Cell struct {
	value      byte
	neighbours byte
	free       bool
}

func NewGridFromFile(filename string) Grid {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file %q: %v", filename, err)
	}

	var grid Grid

	r := bufio.NewReader(file)
	for line, _, err := r.ReadLine(); err != io.EOF; line, _, err = r.ReadLine() {
		if err != nil {
			log.Fatalf("can't read bank: %v", err)
		}

		row := make([]Cell, len(line))
		for x, c := range line {
			row[x].value = c
		}
		grid = append(grid, row)
	}

	return grid
}

func (g Grid) Show() {
	for y := range len(g) {
		for x := range len(g[y]) {
			if g[y][x].free {
				fmt.Printf("x")
			} else {
				fmt.Printf("%c", g[y][x].value)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g Grid) Mark() int {
	count := 0
	for y := range len(g) {
		for x := range len(g[y]) {
			if g[y][x].value == '.' {
				// do not count empty spaces
				continue
			}
			n := g.CountNeighboursOf(x, y)
			g[y][x].neighbours = n
			g[y][x].free = n < 4
			if n < 4 {
				count++
			}
		}
	}
	return count
}

func (g Grid) Sweep() int {
	count := 0
	for y := range len(g) {
		for x := range len(g[y]) {
			if g[y][x].free {
				g[y][x].value = '.'
				g[y][x].free = false
				count++
			}
		}
	}
	return count
}

func (g Grid) CountNeighboursOf(x, y int) byte {
	//       x-1   x   x+1
	//
	// y-1   -,-  -,0  -,+
	//
	//  y    0,-  N/A  0,+
	//
	// y+1   +,-  +,0  +,+
	var count byte
	count += g.CountRollAt(x-1, y-1)
	count += g.CountRollAt(x-1, y)
	count += g.CountRollAt(x-1, y+1)
	count += g.CountRollAt(x, y-1)
	count += g.CountRollAt(x, y+1)
	count += g.CountRollAt(x+1, y-1)
	count += g.CountRollAt(x+1, y)
	count += g.CountRollAt(x+1, y+1)
	return count
}

func (g Grid) CountRollAt(x, y int) byte {
	if g.ValueAt(x, y) == '@' {
		return 1
	}
	return 0
}

func (g Grid) ValueAt(x, y int) byte {
	if y < 0 || y >= len(g) || x < 0 || x >= len(g[y]) {
		return 0
	}
	return g[y][x].value
}
