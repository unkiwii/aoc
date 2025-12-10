package main

import (
	"testing"
)

func TestDay9Part1(t *testing.T) {
	filename := "input/day9.test"
	want := 50
	got := Day9Part1(filename)
	if got != want {
		t.Errorf("Day9Part1(%q) got %d; want: %d", filename, got, want)
	}
}

func TestDay9Part2(t *testing.T) {
	filename := "input/day9.test"
	want := 24
	got := Day9Part2(filename)
	if got != want {
		t.Errorf("Day9Part2(%q) got %d; want: %d", filename, got, want)
	}
}
