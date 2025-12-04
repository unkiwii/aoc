package main

import (
	"fmt"
	"time"
)

func timeit(name string, f func() int) {
	start := time.Now()
	r := f()
	fmt.Printf("[%14s] %s: %d\n", time.Since(start), name, r)
}

func main() {
	timeit("Day 1.1", func() int { return Day1Part1(50, "input/day1") })
	timeit("Day 1.2", func() int { return Day1Part2(50, "input/day1") })
	timeit("Day 2.1", func() int { return Day2Part1("input/day2") })
	timeit("Day 2.2", func() int { return Day2Part2("input/day2") })
	timeit("Day 3.1", func() int { return Day3Part1("input/day3") })
	timeit("Day 3.2", func() int { return Day3Part2("input/day3") })
}
