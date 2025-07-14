package main

import (
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

var s []int

func main() {
	if len(os.Args) < 2 {
		panic("output file argument is required")
	}
	f, err := os.Create(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	rng := rand.New(rand.NewSource(42))

	start := time.Now()
	for time.Since(start) < 5*time.Second {
		s = make([]int, 1000)
		for i := range s {
			s[i] = rng.Intn(1000)
		}
		bubblesort(s)
	}
}

func bubblesort(s []int) {
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[i] > s[j] {
				s[i], s[j] = s[j], s[i]
			}
		}
	}
}
