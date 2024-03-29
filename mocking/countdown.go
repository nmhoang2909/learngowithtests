package mocking

import (
	"fmt"
	"io"
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

type SpyCountdownOperations struct {
	Calls []string
}

var (
	Sleep = "sleep"
	Write = "write"
)

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, Sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, Write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
