package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCountDown(t *testing.T) {
	t.Run("print 3 times to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		CountDown(buffer, &SpyCountdownOperations{})
		want := `3
2
1
Go!`
		got := buffer.String()
		assert.Equal(t, want, got)
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			Write,
			Sleep,

			Write,
			Sleep,

			Write,
			Sleep,

			Write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("want: %v but got: %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{
		duration: sleepTime,
		sleep:    spyTime.Sleep,
	}
	sleeper.Sleep()

	assert.Equal(t, sleepTime, spyTime.durationSlept)
}
