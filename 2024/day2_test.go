package main

import (
	"testing"
)

func TestDay2Part1(t *testing.T) {
	filename := "input/day2.test"
	want := 2
	got := Day2Part1(filename)
	if got != want {
		t.Errorf("Day2Part1(%q) got %d; want: %d", filename, got, want)
	}
}

func TestDay2Part2(t *testing.T) {
	filename := "input/day2.test"
	want := 4
	got := Day2Part2(filename)
	if got != want {
		t.Errorf("Day2Part2(%q) got %d; want: %d", filename, got, want)
	}
}
