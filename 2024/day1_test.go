package main

import (
	"testing"
)

func TestDay1Part1(t *testing.T) {
	filename := "input/day1.test"
	want := 11
	got := Day1Part1(filename)
	if got != want {
		t.Errorf("Day1Part1(%q) got %d; want: %d", filename, got, want)
	}
}

func TestDay1Part2(t *testing.T) {
	filename := "input/day1.test"
	want := 31
	got := Day1Part2(filename)
	if got != want {
		t.Errorf("Day1Part2(%q) got %d; want: %d", filename, got, want)
	}
}
