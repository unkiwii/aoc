package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2Part1(filename string) int {
	return day2(filename, func(id int) bool {
		// any invalid id is one that has the same repeated digits twice
		s := strconv.FormatInt(int64(id), 10)
		if len(s)%2 == 1 {
			// ignore any odd digit number
			return false
		}
		h := len(s) / 2
		left, right := s[:h], s[h:]
		return strings.Compare(left, right) == 0
	})
}

func Day2Part2(filename string) int {
	return day2(filename, func(id int) bool {
		isInvalidParts := func(parts []string) bool {
			for i := range parts {
				if parts[0] != parts[i] {
					return false
				}
			}
			return len(parts) != 0
		}

		// any invalid id is one that has the same repeated digits at least twice
		s := strconv.FormatInt(int64(id), 10)
		half := len(s) / 2
		for i := 1; i <= half; i++ {
			parts := SplitN(s, i)
			if isInvalidParts(parts) {
				return true
			}
		}
		return false
	})
}

func day2(filename string, isInvalidID func(int) bool) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file %q: %v", filename, err)
	}

	result := 0

	r := bufio.NewReader(file)
	for {
		ran, isEOF := readRange(r)

		for n := range RangeOver(ran) {
			if isInvalidID(n) {
				result += n
			}
		}

		if isEOF {
			return result
		}
	}
}

func SplitN(s string, n int) []string {
	if n < 1 {
		return nil
	}
	var l []string
	if len(s)%n != 0 {
		return nil
	}
	left, right := 0, 0
	for i := n; i <= len(s); i += n {
		right = i
		l = append(l, s[left:right])
		left = right
	}
	return l
}

const Comma = byte(',')

var Hyphen = []byte("-")

func readRange(r *bufio.Reader) (Range, bool) {
	data, err := r.ReadBytes(Comma)
	isEOF := err == io.EOF
	if !isEOF && err != nil {
		log.Fatalf("can't read range: %v", err)
	}

	// remove Comma or newline from data
	data = data[:len(data)-1]
	parts := bytes.Split(data, Hyphen)
	return Range{
		low:  parts[0],
		high: parts[1],
	}, isEOF
}

type Range struct {
	low, high []byte
}

type RangeOverYield func(int) bool

func RangeOver(r Range) func(RangeOverYield) {
	noop := func(yield RangeOverYield) {}

	low, err := strconv.Atoi(string(r.low))
	if err != nil {
		return noop
	}
	high, err := strconv.Atoi(string(r.high))
	if err != nil {
		return noop
	}

	return func(yield RangeOverYield) {
		for n := low; n <= high; n++ {
			if !yield(n) {
				return
			}
		}
	}
}
