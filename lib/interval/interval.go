package interval

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
)

// Read an interval from the reader
//
// Starting at the current position of the reader, read bytes until it found
// delim, once there it will separate the buffer in two parts using the
// separator
//
// For example, with this data available in the reader: "123,456=789,abc"
//
// This call:
//
//	interval.Read(r, '=', []byte{','})
//
// # Will return an interval with values 123 and 456
//
// If the end of file (EOF) is reached then the interval and true are returned
func Read(r *bufio.Reader, delim byte, separator []byte) (Interval, bool) {
	data, err := r.ReadBytes(delim)
	isEOF := err == io.EOF
	if !isEOF && err != nil {
		log.Fatalf("can't read range: %v", err)
	}

	// remove delim from data
	data = data[:len(data)-1]
	parts := bytes.Split(data, separator)

	low, err := strconv.Atoi(string(parts[0]))
	if err != nil {
		log.Fatalf("can't parse range low value: %v", err)
	}
	high, err := strconv.Atoi(string(parts[1]))
	if err != nil {
		log.Fatalf("can't parse range high value: %v", err)
	}

	return Interval{
		low:  low,
		high: high,
	}, isEOF
}

type Interval struct {
	low, high int
}

func (i Interval) String() string {
	return fmt.Sprintf("[%d, %d]", i.low, i.high)
}

func (i Interval) Contains(n int) bool {
	return i.low <= n && n <= i.high
}

type RangeYield func(int) bool

func Range(r Interval) func(RangeYield) {
	return func(yield RangeYield) {
		for n := r.low; n <= r.high; n++ {
			if !yield(n) {
				return
			}
		}
	}
}
