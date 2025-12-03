package main

import (
	"fmt"
	"testing"
)

func TestDay2Part1(t *testing.T) {
	filename := "input/day2.test"
	want := 1227775554
	got := Day2Part1(filename)
	if got != want {
		t.Errorf("Day2Part1(%q) got %d; want: %d", filename, got, want)
	}
}

func TestDay2Part2(t *testing.T) {
	filename := "input/day2.test"
	want := 4174379265
	got := Day2Part2(filename)
	if got != want {
		t.Errorf("Day2Part2(%q) got %d; want: %d", filename, got, want)
	}
}

func TestSplitN(t *testing.T) {
	equal := func(got, want []string) bool {
		if len(got) != len(want) {
			return false
		}
		for i := range want {
			if want[i] != got[i] {
				return false
			}
		}
		return true
	}

	for _, tc := range []struct {
		s    string
		n    int
		want []string
	}{
		{s: "", n: 0, want: nil},
		{s: "", n: 1, want: nil},
		{s: "", n: 2, want: nil},
		{s: "", n: 3, want: nil},
		{s: "1", n: 0, want: nil},
		{s: "1", n: 1, want: []string{"1"}},
		{s: "1", n: 2, want: nil},
		{s: "1", n: 3, want: nil},
		{s: "1", n: 4, want: nil},
		{s: "1", n: 5, want: nil},
		{s: "12", n: 0, want: nil},
		{s: "12", n: 1, want: []string{"1", "2"}},
		{s: "12", n: 2, want: []string{"12"}},
		{s: "12", n: 3, want: nil},
		{s: "12", n: 4, want: nil},
		{s: "12", n: 5, want: nil},
		{s: "123", n: 0, want: nil},
		{s: "123", n: 1, want: []string{"1", "2", "3"}},
		{s: "123", n: 2, want: nil},
		{s: "123", n: 3, want: []string{"123"}},
		{s: "123", n: 4, want: nil},
		{s: "123", n: 5, want: nil},
		{s: "1234", n: 0, want: nil},
		{s: "1234", n: 1, want: []string{"1", "2", "3", "4"}},
		{s: "1234", n: 2, want: []string{"12", "34"}},
		{s: "1234", n: 3, want: nil},
		{s: "1234", n: 4, want: []string{"1234"}},
		{s: "12345", n: 0, want: nil},
		{s: "12345", n: 1, want: []string{"1", "2", "3", "4", "5"}},
		{s: "12345", n: 2, want: nil},
		{s: "12345", n: 3, want: nil},
		{s: "12345", n: 4, want: nil},
		{s: "12345", n: 5, want: []string{"12345"}},
		{s: "123456", n: 0, want: nil},
		{s: "123456", n: 1, want: []string{"1", "2", "3", "4", "5", "6"}},
		{s: "123456", n: 2, want: []string{"12", "34", "56"}},
		{s: "123456", n: 3, want: []string{"123", "456"}},
		{s: "123456", n: 4, want: nil},
		{s: "123456", n: 5, want: nil},
		{s: "123456", n: 6, want: []string{"123456"}},
		{s: "123456789", n: -40, want: nil},
		{s: "123456789", n: 0, want: nil},
		{s: "123456789", n: 1, want: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}},
		{s: "123456789", n: 2, want: nil},
		{s: "123456789", n: 3, want: []string{"123", "456", "789"}},
		{s: "123456789", n: 4, want: nil},
		{s: "123456789", n: 5, want: nil},
		{s: "123456789", n: 6, want: nil},
		{s: "123456789", n: 9, want: []string{"123456789"}},
		{s: "123456789", n: 100, want: nil},
	} {
		t.Run(fmt.Sprintf("SplitN(%q,%d)", tc.s, tc.n), func(t *testing.T) {
			got := SplitN(tc.s, tc.n)
			if !equal(got, tc.want) {
				t.Errorf("want: %v, got: %v", tc.want, got)
			}
		})
	}
}
