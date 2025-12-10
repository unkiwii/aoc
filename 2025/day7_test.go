package main

import (
	"testing"
)

func TestDay7Part1(t *testing.T) {
	filename := "input/day7.test"
	want := 21
	got := Day7Part1(filename)
	if got != want {
		t.Errorf("Day7Part1(%q) got %d; want: %d", filename, got, want)
	}
}

// func TestDay7Part2(t *testing.T) {
// 	filename := "input/day7.test"
// 	want := 40
// 	got := Day7Part2(filename)
// 	if got != want {
// 		t.Errorf("Day7Part2(%q) got %d; want: %d", filename, got, want)
// 	}
// }
