package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/unkiwii/aoc/lib/heap"
)

// --- Day 8: Playground ---
//
// Equipped with a new understanding of teleporter maintenance, you confidently
// step onto the repaired teleporter pad.
//
// You rematerialize on an unfamiliar teleporter pad and find yourself in a
// vast underground space which contains a giant playground!
//
// Across the playground, a group of Elves are working on setting up an
// ambitious Christmas decoration project. Through careful rigging, they have
// suspended a large number of small electrical junction boxes.
//
// Their plan is to connect the junction boxes with long strings of lights.
// Most of the junction boxes don't provide electricity; however, when two
// junction boxes are connected by a string of lights, electricity can pass
// between those two junction boxes.
//
// The Elves are trying to figure out which junction boxes to connect so that
// electricity can reach every junction box. They even have a list of all of
// the junction boxes' positions in 3D space (your puzzle input).
//
// For example:
//
// 162,817,812
// 57,618,57
// 906,360,560
// 592,479,940
// 352,342,300
// 466,668,158
// 542,29,236
// 431,825,988
// 739,650,466
// 52,470,668
// 216,146,977
// 819,987,18
// 117,168,530
// 805,96,715
// 346,949,466
// 970,615,88
// 941,993,340
// 862,61,35
// 984,92,344
// 425,690,689
//
// This list describes the position of 20 junction boxes, one per line. Each
// position is given as X,Y,Z coordinates. So, the first junction box in the
// list is at X=162, Y=817, Z=812.
//
// To save on string lights, the Elves would like to focus on connecting pairs
// of junction boxes that are as close together as possible according to
// straight-line distance. In this example, the two junction boxes which are
// closest together are 162,817,812 and 425,690,689.
//
// By connecting these two junction boxes together, because electricity can
// flow between them, they become part of the same circuit. After connecting
// them, there is a single circuit which contains two junction boxes, and the
// remaining 18 junction boxes remain in their own individual circuits.
//
// Now, the two junction boxes which are closest together but aren't already
// directly connected are 162,817,812 and 431,825,988. After connecting them,
// since 162,817,812 is already connected to another junction box, there is now
// a single circuit which contains three junction boxes and an additional 17
// circuits which contain one junction box each.
//
// The next two junction boxes to connect are 906,360,560 and 805,96,715. After
// connecting them, there is a circuit containing 3 junction boxes, a circuit
// containing 2 junction boxes, and 15 circuits which contain one junction box
// each.
//
// The next two junction boxes are 431,825,988 and 425,690,689. Because these
// two junction boxes were already in the same circuit, nothing happens!
//
// This process continues for a while, and the Elves are concerned that they
// don't have enough extension cables for all these circuits. They would like
// to know how big the circuits will be.
//
// After making the ten shortest connections, there are 11 circuits: one
// circuit which contains 5 junction boxes, one circuit which contains 4
// junction boxes, two circuits which contain 2 junction boxes each, and seven
// circuits which each contain a single junction box. Multiplying together the
// sizes of the three largest circuits (5, 4, and one of the circuits of size
// 2) produces 40.
//
// Your list contains many junction boxes; connect together the 1000 pairs of
// junction boxes which are closest together. Afterward, what do you get if you
// multiply together the sizes of the three largest circuits?
func Day8Part1(maxConnections int, filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file %q: %v", filename, err)
	}

	var points []Point3D

	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("can't read line from file %q: %v", filename, err)
		}

		points = append(points, NewPoint3DFromLine(line))
	}

	distances := heap.NewWithLess(PointPairLess)
	for i := 0; i < len(points)-1; i++ {
		a := points[i]
		for j := i + 1; j < len(points); j++ {
			b := points[j]
			pair := NewPointPair(a, b)
			distances.PushItem(pair)
		}
	}

	circuits := map[Point3D][]Point3D{}

	for _, p := range points {
		circuits[p] = []Point3D{p}
	}

	i := 1
mainLoop:
	for distances.Len() > 0 && i <= maxConnections {
		pair := distances.PopItem()

		a, b := pair.a, pair.b
		la := len(circuits[pair.a])
		lb := len(circuits[pair.b])

		switch {
		case la == 1 && lb > 1:
			// MERGE a into b (swap a with b)
			a, b = b, a
		case la == 1 && lb == 1:
		case la > 1 && lb == 1:
		case la > 1 && lb > 1:
			// MERGE b into a (keep a and b as they are)
		}

		i++

		for _, p := range circuits[a] {
			if p.Equal(b) {
				// ignore it, is already there
				continue mainLoop
			}
		}
		for _, p := range circuits[b] {
			if p.Equal(a) {
				// ignore it, is already there
				continue mainLoop
			}
		}

		circuits[a] = append(circuits[a], circuits[b]...)

		for _, p := range circuits[a] {
			circuits[p] = circuits[a]
		}
	}

	circuitSizes := map[int]int{}
	for _, circuit := range circuits {
		size := len(circuit)
		circuitSizes[size]++
	}

	maxSizeOfCircuits := func() int {
		maxSize := 0
		for size := range circuitSizes {
			if size > maxSize {
				maxSize = size
			}
		}
		return maxSize
	}

	result := 1
	for range 3 {
		size := maxSizeOfCircuits()
		result *= size
		delete(circuitSizes, size)
	}

	return result
}

func Day8Part2(filename string) int {
	return 0
}

type Point3D struct {
	X, Y, Z int
}

func NewPoint3DFromLine(line []byte) Point3D {
	parts := bytes.Split(line, []byte(","))
	if len(parts) != 3 {
		log.Fatalf("can't parse line as a 3d point; expected \"X,Y,Z\", but got: %q", line)
	}

	x, err := strconv.Atoi(string(parts[0]))
	if err != nil {
		log.Fatalf("can't parse %q as X coordinate: %v", parts[0], err)
	}
	y, err := strconv.Atoi(string(parts[1]))
	if err != nil {
		log.Fatalf("can't parse %q as Y coordinate: %v", parts[1], err)
	}
	z, err := strconv.Atoi(string(parts[2]))
	if err != nil {
		log.Fatalf("can't parse %q as Z coordinate: %v", parts[2], err)
	}

	return Point3D{X: x, Y: y, Z: z}
}

func (p Point3D) String() string {
	return fmt.Sprintf("(%3d, %3d, %3d)", p.X, p.Y, p.Z)
}

func (p Point3D) Equal(o Point3D) bool {
	return p.X == o.X && p.Y == o.Y && p.Z == o.Z
}

func (p Point3D) SquaredLinearDistanceTo(o Point3D) int {
	dx := o.X - p.X
	dy := o.Y - p.Y
	dz := o.Z - p.Z
	return dx*dx + dy*dy + dz*dz
}

type PointPair struct {
	a, b     Point3D
	distance int
}

func NewPointPair(a, b Point3D) PointPair {
	return PointPair{
		a:        a,
		b:        b,
		distance: a.SquaredLinearDistanceTo(b),
	}
}

func PointPairLess(a, b PointPair) bool {
	return a.distance < b.distance
}

func (pp PointPair) String() string {
	return fmt.Sprintf("%s -> %s = %d", pp.a, pp.b, pp.distance)
}
