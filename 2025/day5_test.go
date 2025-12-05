package main

import (
	"testing"
)

func TestDay5Part1(t *testing.T) {
	filename := "input/day5.test"
	want := 3
	got := Day5Part1(filename)
	if got != want {
		t.Errorf("Day5Part1(%q) got %d; want: %d", filename, got, want)
	}
}

func TestDay5Part2(t *testing.T) {
	filename := "input/day5.test"
	want := 14
	got := Day5Part2(filename)
	if got != want {
		t.Errorf("Day5Part2(%q) got %d; want: %d", filename, got, want)
	}
}
