package time

import (
	"fmt"
	"time"
)

func It(name string, f func() int) {
	start := time.Now()
	r := f()
	fmt.Printf("[%14s] %s: %d\n", time.Since(start), name, r)
}
