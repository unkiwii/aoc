package main

import (
	"github.com/unkiwii/aoc/lib/time"
)

func main() {
	time.It("Day1Part1", func() int { return Day1Part1("input/day1") })
	time.It("Day1Part2", func() int { return Day1Part2("input/day1") })
	time.It("Day2Part1", func() int { return Day2Part1("input/day2") })
	time.It("Day2Part2", func() int { return Day2Part2("input/day2") })
}
