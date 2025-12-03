package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
)

// possible rotation: Ln, Rn
// 0 <= n <= 99 (n is distance, not absolute)
// L: go to lower number
// R: go to higher number
// dial is a circle: R1 from 99 is 0, L1 from 0 is 99

// password is how many times the dial is pointing at 0 after any rotation
func Day1Part1(start int, filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file %q: %v", filename, err)
	}

	pos := start
	password := 0

	r := bufio.NewReader(file)
	for {
		sign := readSign(r)
		if sign == 999 {
			return password
		}

		offset := readOffset(r)

		pos = (pos + sign*offset) % 100
		if pos < 0 {
			pos += 100
		}
		if pos == 0 {
			password++
		}
	}
}

func readSign(r *bufio.Reader) int {
	dir, size, err := r.ReadRune()
	if err == io.EOF {
		return 999
	}
	if size != 1 {
		log.Fatalf("invalid rune size %d: expected 'L' or 'R' of size 1", size)
	}
	if err != nil {
		log.Fatalf("can't read rune 'L' or 'R': %v", err)
	}
	switch dir {
	case 'R':
		return 1
	case 'L':
		return -1
	}

	log.Fatalf("invalid direction %c", dir)
	return 0
}

func readOffset(r *bufio.Reader) int {
	line, _, err := r.ReadLine()
	if err == io.EOF {
		return 0
	}
	if err != nil {
		log.Fatalf("can't read offset: %v", err)
	}
	n, err := strconv.Atoi(string(line))
	if err != nil {
		log.Fatalf("invalid offset: %v", err)
	}
	return n
}

func Day1Part2(start int, filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file %q: %v", filename, err)
	}

	pos := start
	password := 0

	r := bufio.NewReader(file)
	for {
		sign := readSign(r)
		if sign == 999 {
			return password
		}

		offset := readOffset(r)

		for range offset {
			pos += sign
			if pos < 0 {
				pos += 100
			}
			if pos > 99 {
				pos -= 100
			}
			if pos == 0 {
				password++
			}
		}
	}
}
