package main

import (
	"testing"
)

func TestDay10Part1(t *testing.T) {
	filename := "input/day10.test"
	want := -1
	got := Day10Part1(filename)
	if got != want {
		t.Errorf("Day10Part1(%q) got %d; want: %d", filename, got, want)
	}
}

func TestDay10Part2(t *testing.T) {
	filename := "input/day10.test"
	want := -1
	got := Day10Part2(filename)
	if got != want {
		t.Errorf("Day10Part2(%q) got %d; want: %d", filename, got, want)
	}
}
