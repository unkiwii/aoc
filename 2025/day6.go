package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
)

// --- Day 6: Trash Compactor ---
//
// After helping the Elves in the kitchen, you were taking a break and helping
// them re-enact a movie scene when you over-enthusiastically jumped into the
// garbage chute!
//
// A brief fall later, you find yourself in a garbage smasher. Unfortunately,
// the door's been magnetically sealed.
//
// As you try to find a way out, you are approached by a family of cephalopods!
// They're pretty sure they can get the door open, but it will take some time.
// While you wait, they're curious if you can help the youngest cephalopod with
// her math homework.
//
// Cephalopod math doesn't look that different from normal math. The math
// worksheet (your puzzle input) consists of a list of problems; each problem
// has a group of numbers that need to be either added (+) or multiplied (*)
// together.
//
// However, the problems are arranged a little strangely; they seem to be
// presented next to each other in a very long horizontal list. For example:
//
//	123 328  51 64
//	 45 64  387 23
//	  6 98  215 314
//	*   +   *   +
//
// Each problem's numbers are arranged vertically; at the bottom of the problem
// is the symbol for the operation that needs to be performed. Problems are
// separated by a full column of only spaces. The left/right alignment of
// numbers within each problem can be ignored.
//
// So, this worksheet contains four problems:
//
//	123 * 45 * 6 = 33210
//	328 + 64 + 98 = 490
//	51 * 387 * 215 = 4243455
//	64 + 23 + 314 = 401
//
// To check their work, cephalopod students are given the grand total of adding
// together all of the answers to the individual problems. In this worksheet,
// the grand total is 33210 + 490 + 4243455 + 401 = 4277556.
//
// Of course, the actual worksheet is much wider. You'll need to make sure to
// unroll it completely so that you can read the problems clearly.
//
// Solve the problems on the math worksheet. What is the grand total found by
// adding together all of the answers to the individual problems?
func Day6Part1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file %q: %v", filename, err)
	}

	space := []byte(" ")
	plus := byte('+')
	star := byte('*')

	operands := map[int][]int{}
	var operations []byte

	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			return calculateTotal(operands, operations)
		}

		splitted := bytes.Split(line, space)
		i := 0
		for _, s := range splitted {
			if len(s) == 0 {
				continue
			}

			if s[0] == plus || s[0] == star {
				operations = append(operations, s[0])
				continue
			}

			n, err := strconv.Atoi(string(s))
			if err != nil {
				log.Fatalf("can't parse operand %q: %v", s, err)
			}

			operands[i] = append(operands[i], n)
			i++
		}
	}
}

// --- Part Two ---
//
// The big cephalopods come back to check on how things are going. When they
// see that your grand total doesn't match the one expected by the worksheet,
// they realize they forgot to explain how to read cephalopod math.
//
// Cephalopod math is written right-to-left in columns. Each number is given in
// its own column, with the most significant digit at the top and the least
// significant digit at the bottom. (Problems are still separated with a column
// consisting only of spaces, and the symbol at the bottom of the problem is
// still the operator to use.)
//
// Here's the example worksheet again:
//
//	123 328  51 64
//	 45 64  387 23
//	  6 98  215 314
//	*   +   *   +
//
// Reading the problems right-to-left one column at a time, the problems are
// now quite different:
//
// - The rightmost problem is 4 + 431 + 623 = 1058
// - The second problem from the right is 175 * 581 * 32 = 3253600
// - The third problem from the right is 8 + 248 + 369 = 625
// - Finally, the leftmost problem is 356 * 24 * 1 = 8544
//
// Now, the grand total is 1058 + 3253600 + 625 + 8544 = 3263827.
//
// Solve the problems on the math worksheet again. What is the grand total
// found by adding together all of the answers to the individual problems?
func Day6Part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file %q: %v", filename, err)
	}

	var grid [][]byte

	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		row := make([]byte, len(line))
		copy(row, line)
		grid = append(grid, row)
	}

	rotatedGrid := rotate(grid)

	var operands []int
	var operand []byte
	var operator func([]int) int

	result := 0

	makeOperand := func() {
		if len(operand) != 0 {
			n, err := strconv.Atoi(string(operand))
			if err != nil {
				log.Fatalf("can't parse number %s: %v", operand, err)
			}
			operand = operand[0:0]
			operands = append(operands, n)
		}
	}

	for _, line := range rotatedGrid {
		trimmed := bytes.Trim(line, " ")
		if len(trimmed) == 0 && operator != nil {
			result += operator(operands)
			operands = operands[0:0]
			operator = nil
		}
		for _, value := range line {
			switch value {
			case ' ':
				// ignore spaces
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				operand = append(operand, value)
			case '*':
				operator = mul
			case '+':
				operator = sum
			}
		}
		makeOperand()
	}
	if operator != nil {
		result += operator(operands)
		operands = operands[0:0]
		operator = nil
	}

	return result
}

func calculateTotal(operands map[int][]int, operations []byte) int {
	total := 0
	for i, op := range operations {
		switch op {
		case '+':
			total += sum(operands[i])
		case '*':
			total += mul(operands[i])
		}
	}
	return total
}

func sum(l []int) int {
	r := 0
	for _, n := range l {
		r += n
	}
	return r
}

func mul(l []int) int {
	r := 1
	for _, n := range l {
		r *= n
	}
	return r
}

func rotate(grid [][]byte) [][]byte {
	if len(grid) < 1 {
		return grid
	}

	result := make([][]byte, len(grid[0]))
	for i := range result {
		result[i] = make([]byte, len(grid))
	}

	for y, row := range grid {
		for x, v := range row {
			result[x][y] = v
		}
	}

	return result
}
