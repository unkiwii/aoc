package main

import (
	"testing"
)

func TestDay8Part1(t *testing.T) {
	filename := "input/day8.test"
	want := 40
	got := Day8Part1(10, filename)
	if got != want {
		t.Errorf("Day8Part1(%q) got %d; want: %d", filename, got, want)
	}
}

func TestDay8Part2(t *testing.T) {
	filename := "input/day8.test"
	want := -1
	got := Day8Part2(filename)
	if got != want {
		t.Errorf("Day8Part2(%q) got %d; want: %d", filename, got, want)
	}
}
