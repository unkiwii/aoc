package main

import "github.com/unkiwii/aoc/lib/time"

func main() {
	time.It("Day  1.1", func() int { return Day1Part1(50, "input/day1") })
	time.It("Day  1.2", func() int { return Day1Part2(50, "input/day1") })
	time.It("Day  2.1", func() int { return Day2Part1("input/day2") })
	time.It("Day  2.2", func() int { return Day2Part2("input/day2") })
	time.It("Day  3.1", func() int { return Day3Part1("input/day3") })
	time.It("Day  3.2", func() int { return Day3Part2("input/day3") })
	time.It("Day  4.1", func() int { return Day4Part1("input/day4") })
	time.It("Day  4.2", func() int { return Day4Part2("input/day4") })
	time.It("Day  5.1", func() int { return Day5Part1("input/day5") })
	time.It("Day  5.2", func() int { return Day5Part2("input/day5") })
	time.It("Day  6.1", func() int { return Day6Part1("input/day6") })
	time.It("Day  6.2", func() int { return Day6Part2("input/day6") })
	time.It("Day  7.1", func() int { return Day7Part1("input/day7") })
	// time.It("Day 7.2", func() int { return Day7Part2("input/day7") })
	time.It("Day  8.1", func() int { return Day8Part1(1000, "input/day8") })
	time.It("Day  8.2", func() int { return Day8Part2("input/day8") })
	time.It("Day  9.1", func() int { return Day9Part1("input/day9") })
	// time.It("Day 9.2", func() int { return Day9Part2("input/day9") })
	time.It("Day 10.1", func() int { return Day10Part1("input/day10") })
	// time.It("Day 10.2", func() int { return Day10Part2("input/day10") })
}
