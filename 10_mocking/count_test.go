package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

const sleep = "sleep"
const write = "write"

type CountdownOperationSpy struct {
	Calls []string
}

func (c *CountdownOperationSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("print data", func(t *testing.T) {

		buffer := bytes.Buffer{}

		Countdown(&buffer, &CountdownOperationSpy{})

		got := buffer.String()
		want :=
			`321 Go!`

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}

	})

	t.Run("sleep before every write", func(t *testing.T) {
		spySleeperCounter := &CountdownOperationSpy{}
		Countdown(spySleeperCounter, spySleeperCounter)

		want := []string{
			sleep,
			write,
			sleep,
			write,

			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleeperCounter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleeperCounter.Calls)
		}

	})

}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}

	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
