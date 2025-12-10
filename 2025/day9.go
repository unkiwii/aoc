package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// --- Day 9: Movie Theater ---
//
// You slide down the firepole in the corner of the playground and land in the
// North Pole base movie theater!
//
// The movie theater has a big tile floor with an interesting pattern. Elves
// here are redecorating the theater by switching out some of the square tiles
// in the big grid they form. Some of the tiles are red; the Elves would like
// to find the largest rectangle that uses red tiles for two of its opposite
// corners. They even have a list of where the red tiles are located in the
// grid (your puzzle input).
//
// For example:
//
// 7,1
// 11,1
// 11,7
// 9,7
// 9,5
// 2,5
// 2,3
// 7,3
//
// Showing red tiles as # and other tiles as ., the above arrangement of red
// tiles would look like this:
//
// ..............
// .......#...#..
// ..............
// ..#....#......
// ..............
// ..#......#....
// ..............
// .........#.#..
// ..............
//
// You can choose any two red tiles as the opposite corners of your rectangle;
// your goal is to find the largest rectangle possible.
//
// For example, you could make a rectangle (shown as O) with an area of 24
// between 2,5 and 9,7:
//
// ..............
// .......#...#..
// ..............
// ..#....#......
// ..............
// ..OOOOOOOO....
// ..OOOOOOOO....
// ..OOOOOOOO.#..
// ..............
//
// Or, you could make a rectangle with area 35 between 7,1 and 11,7:
//
// ..............
// .......OOOOO..
// .......OOOOO..
// ..#....OOOOO..
// .......OOOOO..
// ..#....OOOOO..
// .......OOOOO..
// .......OOOOO..
// ..............
//
// You could even make a thin rectangle with an area of only 6 between 7,3 and
// 2,3:
//
// ..............
// .......#...#..
// ..............
// ..OOOOOO......
// ..............
// ..#......#....
// ..............
// .........#.#..
// ..............
//
// Ultimately, the largest rectangle you can make in this example has area 50.
// One way to do this is between 2,5 and 11,1:
//
// ..............
// ..OOOOOOOOOO..
// ..OOOOOOOOOO..
// ..OOOOOOOOOO..
// ..OOOOOOOOOO..
// ..OOOOOOOOOO..
// ..............
// .........#.#..
// ..............
//
// Using two red tiles as opposite corners, what is the largest area of any
// rectangle you can make?
func Day9Part1(filename string) int {
	grid := NewDay9GridFromFile(filename)
	l := len(grid.redTiles)

	var maxArea int
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			a := grid.redTiles[i]
			b := grid.redTiles[j]

			r := NewRect(a, b)
			area := r.Area()
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func Day9Part2(filename string) int {
	return 0
}

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%2d, %2d)", p.X, p.Y)
}

type Rect struct {
	A, B        Point
	Top, Bottom Point
}

func NewRect(a, b Point) Rect {
	return Rect{
		A: a,
		B: b,
		Top: Point{
			X: min(a.X, b.X),
			Y: min(a.Y, b.Y),
		},
		Bottom: Point{
			X: max(a.X, b.X),
			Y: max(a.Y, b.Y),
		},
	}
}

func (r Rect) Area() int {
	return (r.Bottom.X - r.Top.X + 1) * (r.Bottom.Y - r.Top.Y + 1)
}

type Day9Grid struct {
	cells      map[Point]Day9Cell
	redTiles   []Point
	minX, minY int
	maxX, maxY int
}

type Day9Cell byte

const (
	Day9CellEmpty = Day9Cell(0)
	Day9CellRed   = Day9Cell(1)
)

const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

func NewDay9GridFromFile(filename string) Day9Grid {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file %q: %v", filename, err)
	}

	var maxPoint Point
	var redTiles []Point
	points := map[Point]bool{}
	cells := map[Point]Day9Cell{}
	maxX, maxY := 0, 0
	minX, minY := MaxInt, MaxInt

	r := bufio.NewReader(file)
	for line, _, err := r.ReadLine(); err != io.EOF; line, _, err = r.ReadLine() {
		if err != nil {
			log.Fatalf("can't read bank: %v", err)
		}

		parts := bytes.Split(line, []byte(","))
		if len(parts) != 2 {
			log.Fatalf("can't parse line; expected X,Y but got: %s", line)
		}

		x, err := strconv.Atoi(string(parts[0]))
		if err != nil {
			log.Fatalf("can't parse X coordinate from line %q: %v", line, err)
		}
		y, err := strconv.Atoi(string(parts[1]))
		if err != nil {
			log.Fatalf("can't parse Y coordinate from line %q: %v", line, err)
		}

		point := Point{X: x, Y: y}
		if point.X > maxPoint.X && point.Y > maxPoint.Y {
			maxPoint = point
		}

		points[point] = true
		redTiles = append(redTiles, point)
		cells[point] = Day9CellRed

		minX = min(minX, point.X)
		minY = min(minY, point.Y)
		maxX = max(maxX, point.X)
		maxY = max(maxY, point.Y)
	}

	return Day9Grid{
		cells:    cells,
		redTiles: redTiles,
		minX:     minX,
		minY:     minY,
		maxX:     maxX,
		maxY:     maxY,
	}
}

func (grid Day9Grid) IsRed(x, y int) bool {
	return grid.cells[Point{X: x, Y: y}] == Day9CellRed
}

func (grid Day9Grid) Show() {
	fmt.Println()

	maxX := grid.maxX + 3
	maxY := grid.maxY + 2

	fmt.Print(" ")
	for x := range maxX {
		fmt.Print(x % 10)
	}

	for y := range maxY {
		for x := range maxX {
			if x == 0 {
				fmt.Printf("\n%d", y%10)
			}
			if grid.IsRed(x, y) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
	}
	fmt.Println()
	fmt.Println()
}
