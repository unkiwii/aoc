package main

import (
	"testing"
)

func TestDay3Part1(t *testing.T) {
	filename := "input/day3.test"
	want := 357
	got := Day3Part1(filename)
	if got != want {
		t.Errorf("Day3Part1(%q) got %d; want: %d", filename, got, want)
	}
}

func TestDay3Part2(t *testing.T) {
	filename := "input/day3.test"
	want := 3121910778619
	got := Day3Part2(filename)
	if got != want {
		t.Errorf("Day3Part2(%q) got %d; want: %d", filename, got, want)
	}
}
