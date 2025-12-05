package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/unkiwii/aoc/lib/interval"
	"github.com/unkiwii/aoc/lib/list"
)

// --- Day 5: Cafeteria ---
//
// As the forklifts break through the wall, the Elves are delighted to discover
// that there was a cafeteria on the other side after all.
//
// You can hear a commotion coming from the kitchen. "At this rate, we won't
// have any time left to put the wreaths up in the dining hall!" Resolute in
// your quest, you investigate.
//
// "If only we hadn't switched to the new inventory management system right
// before Christmas!" another Elf exclaims. You ask what's going on.
//
// The Elves in the kitchen explain the situation: because of their complicated
// new inventory management system, they can't figure out which of their
// ingredients are fresh and which are spoiled. When you ask how it works, they
// give you a copy of their database (your puzzle input).
//
// The database operates on ingredient IDs. It consists of a list of fresh
// ingredient ID ranges, a blank line, and a list of available ingredient IDs.
// For example:
//
// 3-5
// 10-14
// 16-20
// 12-18
//
// 1
// 5
// 8
// 11
// 17
// 32
//
// The fresh ID ranges are inclusive: the range 3-5 means that ingredient IDs
// 3, 4, and 5 are all fresh. The ranges can also overlap; an ingredient ID is
// fresh if it is in any range.
//
// The Elves are trying to determine which of the available ingredient IDs are
// fresh. In this example, this is done as follows:
//
//	Ingredient ID 1 is spoiled because it does not fall into any range.
//	Ingredient ID 5 is fresh because it falls into range 3-5.
//	Ingredient ID 8 is spoiled.
//	Ingredient ID 11 is fresh because it falls into range 10-14.
//	Ingredient ID 17 is fresh because it falls into range 16-20 as well as range 12-18.
//	Ingredient ID 32 is spoiled.
//
// So, in this example, 3 of the available ingredient IDs are fresh.
//
// Process the database file from the new inventory management system. How many
// of the available ingredient IDs are fresh?
func Day5Part1(filename string) int {
	freshIntervals, ingredients := readDatabase(filename, false)

	count := 0
	for _, in := range ingredients {
		if isFresh(freshIntervals, in) {
			count++
		}
	}
	return count
}

// --- Part Two ---
//
// The Elves start bringing their spoiled inventory to the trash chute at the
// back of the kitchen.
//
// So that they can stop bugging you when they get new inventory, the Elves
// would like to know all of the IDs that the fresh ingredient ID ranges
// consider to be fresh. An ingredient ID is still considered fresh if it is in
// any range.
//
// Now, the second section of the database (the available ingredient IDs) is
// irrelevant. Here are the fresh ingredient ID ranges from the above example:
//
// 3-5
// 10-14
// 16-20
// 12-18
//
// The ingredient IDs that these ranges consider to be fresh are 3, 4, 5, 10,
// 11, 12, 13, 14, 15, 16, 17, 18, 19, and 20. So, in this example, the fresh
// ingredient ID ranges consider a total of 14 ingredient IDs to be fresh.
//
// Process the database file again. How many ingredient IDs are considered to
// be fresh according to the fresh ingredient ID ranges?
func Day5Part2(filename string) int {
	freshIntervals, _ := readDatabase(filename, true)

	mergeIntervals := func(intervals *list.List[interval.Interval]) int {
		amountMerged := 0

		for a := intervals; a != nil; a = a.Next {
			for b := a.Next; b != nil; b = b.Next {
				if b.Value.IsInside(a.Value) {
					a.Next = b.Next
				} else if a.Value.Overlaps(b.Value) {
					a.Value = a.Value.Merge(b.Value)
					amountMerged++
				}
			}
		}

		return amountMerged
	}

	intervals := list.FromSlice(freshIntervals)
	for amountMerged := -1; amountMerged != 0; {
		amountMerged = mergeIntervals(intervals)
	}

	count := 0

	for i := intervals; i != nil; i = i.Next {
		count += i.Value.Distance()
	}

	return count
}

func readDatabase(filename string, stopAtIntervals bool) ([]interval.Interval, []int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file %q: %v", filename, err)
	}

	readingFreshIntervals := true
	var freshIntervals []interval.Interval
	var ingredients []int

	r := bufio.NewReader(file)
	for {
		if readingFreshIntervals {
			i, eof := interval.Read(r, '\n', []byte("-"))
			if eof {
				log.Fatal("unexpected EOF")
			}
			freshIntervals = append(freshIntervals, i)

			next, _ := r.Peek(1)
			if next[0] == '\n' {
				r.Discard(1) // discard empty line
				readingFreshIntervals = false
				if stopAtIntervals {
					return freshIntervals, ingredients
				}
			}

		} else {
			line, _, err := r.ReadLine()
			if err == io.EOF {
				return freshIntervals, ingredients
			}
			if err != nil {
				log.Fatalf("can't read line: %v", err)
			}
			n, err := strconv.Atoi(string(line))
			if err != nil {
				log.Fatalf("can't parse ingredient: %v", err)
			}
			ingredients = append(ingredients, n)
		}
	}
}

func isFresh(freshIntervals []interval.Interval, ingredient int) bool {
	for _, in := range freshIntervals {
		if in.Contains(ingredient) {
			return true
		}
	}
	return false
}
