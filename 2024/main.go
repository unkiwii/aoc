package main

import (
	"github.com/unkiwii/aoc/lib/time"
)

func main() {
	time.It("Day 1.1", func() int { return Day1Part1("input/day1") })
	time.It("Day 1.2", func() int { return Day1Part2("input/day1") })
	time.It("Day 2.1", func() int { return Day2Part1("input/day2") })
	time.It("Day 2.2", func() int { return Day2Part2("input/day2") })
	time.It("Day 3.1", func() int { return Day3Part1("input/day3") })
	time.It("Day 3.2", func() int { return Day3Part2("input/day3") })
}
