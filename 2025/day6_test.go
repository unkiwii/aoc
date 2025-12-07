package main

import (
	"testing"
)

func TestDay6Part1(t *testing.T) {
	filename := "input/day6.test"
	want := 4277556
	got := Day6Part1(filename)
	if got != want {
		t.Errorf("Day6Part1(%q) got %d; want: %d", filename, got, want)
	}
}

func TestDay6Part2(t *testing.T) {
	filename := "input/day6.test"
	want := 3263827
	got := Day6Part2(filename)
	if got != want {
		t.Errorf("Day6Part2(%q) got %d; want: %d", filename, got, want)
	}
}
