package main

import (
	"testing"
)

func TestDay4Part1(t *testing.T) {
	filename := "input/day4.test"
	want := 13
	got := Day4Part1(filename)
	if got != want {
		t.Errorf("Day4Part1(%q) got %d; want: %d", filename, got, want)
	}
}

func TestDay4Part2(t *testing.T) {
	filename := "input/day4.test"
	want := 43
	got := Day4Part2(filename)
	if got != want {
		t.Errorf("Day4Part2(%q) got %d; want: %d", filename, got, want)
	}
}
