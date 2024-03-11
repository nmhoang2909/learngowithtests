package mocking

import (
	"fmt"
	"io"
	"log"
	"time"
)

const DefaultWaiting time.Duration = 3 * time.Second

func CountDown(w io.Writer, sleeper Sleeper) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(w, i)
		// time.Sleep(2 * time.Second)
		sleeper.Sleep()
	}
	fmt.Fprint(w, "Go!")
}

type Sleeper interface {
	Sleep()
}

// Mock SpySleeper
type SpySleeper struct {
	Call int
}

func (s *SpySleeper) Sleep() {
	s.Call++
}

type DefaultSleeper struct{}

// Real Sleeper
func (d *DefaultSleeper) Sleep() {
	log.Printf("Sleeping in %v second", DefaultWaiting)
	time.Sleep(DefaultWaiting)
}
