package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

// --- Day 10: Factory ---
//
// Just across the hall, you find a large factory. Fortunately, the Elves here
// have plenty of time to decorate. Unfortunately, it's because the factory
// machines are all offline, and none of the Elves can figure out the
// initialization procedure.
//
// The Elves do have the manual for the machines, but the section detailing the
// initialization procedure was eaten by a Shiba Inu. All that remains of the
// manual are some indicator light diagrams, button wiring schematics, and
// joltage requirements for each machine.
//
// For example:
//
// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
// [...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
// [.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
//
// The manual describes one machine per line. Each line contains a single
// indicator light diagram in [square brackets], one or more button wiring
// schematics in (parentheses), and joltage requirements in {curly braces}.
//
// To start a machine, its indicator lights must match those shown in the
// diagram, where . means off and # means on. The machine has the number of
// indicator lights shown, but its indicator lights are all initially off.
//
// So, an indicator light diagram like [.##.] means that the machine has four
// indicator lights which are initially off and that the goal is to
// simultaneously configure the first light to be off, the second light to be
// on, the third to be on, and the fourth to be off.
//
// You can toggle the state of indicator lights by pushing any of the listed
// buttons. Each button lists which indicator lights it toggles, where 0 means
// the first light, 1 means the second light, and so on. When you push a
// button, each listed indicator light either turns on (if it was off) or turns
// off (if it was on). You have to push each button an integer number of times;
// there's no such thing as "0.5 presses" (nor can you push a button a negative
// number of times).
//
// So, a button wiring schematic like (0,3,4) means that each time you push
// that button, the first, fourth, and fifth indicator lights would all toggle
// between on and off. If the indicator lights were [#.....], pushing the
// button would change them to be [...##.] instead.
//
// Because none of the machines are running, the joltage requirements are
// irrelevant and can be safely ignored.
//
// You can push each button as many times as you like. However, to save on
// time, you will need to determine the fewest total presses required to
// correctly configure all indicator lights for all machines in your list.
//
// There are a few ways to correctly configure the first machine:
//
// [.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
//
// - You could press the first three buttons once each, a total of 3 button presses.
// - You could press (1,3) once, (2,3) once, and (0,1) twice, a total of 4 button presses.
// - You could press all of the buttons except (1,3) once each, a total of 5 button presses.
//
// However, the fewest button presses required is 2. One way to do this is by
// pressing the last two buttons ((0,2) and (0,1)) once each.
//
// The second machine can be configured with as few as 3 button presses:
//
// [...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
//
// One way to achieve this is by pressing the last three buttons ((0,4),
// (0,1,2), and (1,2,3,4)) once each.
//
// The third machine has a total of six indicator lights that need to be
// configured correctly:
//
// [.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
//
// The fewest presses required to correctly configure it is 2; one way to do
// this is by pressing buttons (0,3,4) and (0,1,2,4,5) once each.
//
// So, the fewest button presses required to correctly configure the indicator
// lights on all of the machines is 2 + 3 + 2 = 7.
//
// Analyze each machine's indicator light diagram and button wiring schematics.
// What is the fewest button presses required to correctly configure the
// indicator lights on all of the machines?
func Day10Part1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file %q: %v", filename, err)
	}

	result := 0

	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			return result
		}
		if err != nil {
			log.Fatalf("can't read line from file %q: %v", filename, err)
		}

		machine := NewDay10MachineFromLine(line)

		fmt.Println()
		fmt.Printf("running: %q\n", machine)
		presses := machine.FindFewestButtonPresses()
		fmt.Printf("presses: %d\n", presses)
		fmt.Println()

		result += presses
	}
}

func Day10Part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("can't open file %q: %v", filename, err)
	}

	r := bufio.NewReader(file)
	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			return 0
		}
		if err != nil {
			log.Fatalf("can't read line from file %q: %v", filename, err)
		}

		fmt.Println(string(line))
	}
}

type Day10Machine struct {
	state               []bool
	endState            []bool
	buttonWirings       [][]int
	joltageRequirements []int
}

type ReadDay10MachineState byte

const (
	UndefinedReadDay10MachineState = ReadDay10MachineState(0)
	EndStateReadDay10MachineState  = ReadDay10MachineState(1)
	WiringsReadDay10MachineState   = ReadDay10MachineState(2)
	JoltageReadDay10MachineState   = ReadDay10MachineState(3)
)

func NewDay10MachineFromLine(line []byte) Day10Machine {
	var endState []bool
	var accum []byte
	var buttonWirings [][]int
	var buttonWiringIndex int
	var joltageRequirements []int

	addWiringButton := func() {
		wiring, err := strconv.Atoi(string(accum))
		if err != nil {
			log.Fatalf("can't parse wiring from %q: %v", accum, err)
		}
		buttonWirings[buttonWiringIndex] = append(buttonWirings[buttonWiringIndex], wiring)
		accum = accum[0:0]
	}

	addJoltageRequirement := func() {
		joltage, err := strconv.Atoi(string(accum))
		if err != nil {
			log.Fatalf("can't parse joltage from %q: %v", accum, err)
		}
		joltageRequirements = append(joltageRequirements, joltage)
		accum = accum[0:0]
	}

	readState := UndefinedReadDay10MachineState
	for _, b := range line {
		switch b {
		case '[':
			readState = EndStateReadDay10MachineState
		case '(':
			buttonWirings = append(buttonWirings, []int{})
			buttonWiringIndex = len(buttonWirings) - 1
			readState = WiringsReadDay10MachineState
		case '{':
			readState = JoltageReadDay10MachineState
		case ']':
			endState = make([]bool, len(accum))
			for i, v := range accum {
				if v == '#' {
					endState[i] = true
				} else {
					endState[i] = false
				}
			}
			accum = accum[0:0]
			readState = UndefinedReadDay10MachineState
		case ')':
			addWiringButton()
			readState = UndefinedReadDay10MachineState
		case '}':
			addJoltageRequirement()
			readState = UndefinedReadDay10MachineState
		case ',':
			switch readState {
			case WiringsReadDay10MachineState:
				addWiringButton()
			case JoltageReadDay10MachineState:
				addJoltageRequirement()
			}
		default:
			switch readState {
			case UndefinedReadDay10MachineState:
				// ignore byte
			case WiringsReadDay10MachineState, JoltageReadDay10MachineState, EndStateReadDay10MachineState:
				accum = append(accum, b)
			}
		}
	}

	return Day10Machine{
		state:               make([]bool, len(endState)),
		endState:            endState,
		buttonWirings:       buttonWirings,
		joltageRequirements: joltageRequirements,
	}
}

func (m Day10Machine) String() string {
	var buf strings.Builder
	buf.WriteString(day10MachineStateString(m.state))
	buf.WriteString(" | ")
	buf.WriteString(day10MachineStateString(m.endState))
	for _, wiring := range m.buttonWirings {
		buf.WriteRune(' ')
		buf.WriteString(day10MachineWiringString(wiring))
	}
	buf.WriteString(" {")
	for i, joltage := range m.joltageRequirements {
		if i != 0 {
			buf.WriteRune(',')
		}
		buf.WriteString(strconv.Itoa(joltage))
	}
	buf.WriteRune('}')
	return buf.String()
}

func day10MachineStateString(state []bool) string {
	var buf strings.Builder
	buf.WriteRune('[')
	for _, light := range state {
		if light {
			buf.WriteByte('#')
		} else {
			buf.WriteByte('.')
		}
	}
	buf.WriteRune(']')
	return buf.String()
}

func day10MachineWiringString(wiring []int) string {
	var buf strings.Builder
	buf.WriteRune('(')
	for i, button := range wiring {
		if i != 0 {
			buf.WriteRune(',')
		}
		buf.WriteString(strconv.Itoa(button))
	}
	buf.WriteRune(')')
	return buf.String()
}

func (m Day10Machine) FindFewestButtonPresses() int {
	// initialState := make([]bool, len(m.state))
	oldState := make([]bool, len(m.state))

	// TODO: make this a bit vector so we can turn on/off things just by AND them?
	minPresses := -1
	presses := 0
	for _, wiring := range m.buttonWirings {
		copy(oldState, m.state)
		for _, button := range wiring {
			m.state[button] = !m.state[button]
		}
		presses++

		fmt.Printf("  %s -> %s, pressed %s\n",
			day10MachineStateString(oldState),
			day10MachineStateString(m.state),
			day10MachineWiringString(wiring),
		)

		if slices.Equal(m.state, m.endState) {
			fmt.Printf("    FOUND END STATE!! IN %d PRESSES\n", presses)
			if presses < minPresses || minPresses < 0 {
				minPresses = presses
			}
			presses = 0
		}
	}

	return minPresses
}
