package countdown

import (
	"bytes"
	"reflect"
	"testing"
)

type Spy struct {
	Calls  []string
	Output *bytes.Buffer
}

func (s *Spy) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func (s *Spy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return s.Output.Write(p)
}

func TestCountdown(t *testing.T) {
	spy := &Spy{Output: &bytes.Buffer{}}

	Countdown(spy, spy)

	wantedOutput := "3\n2\n1\nGo!"
	if spy.Output.String() != wantedOutput {
		t.Errorf("got %q want %q", spy.Output, wantedOutput)
	}
	wantedCalls := []string{
		"sleep", "write", "sleep", "write",
		"sleep", "write", "sleep", "write",
	}
	if !reflect.DeepEqual(wantedCalls, spy.Calls) {
		t.Errorf("wanted calls %v got %v", wantedCalls, spy.Calls)
	}
}
