package cli_test

import (
	"bytes"
	"fmt"
	"jitsusama/lgwt/app/pkg/cli"
	test "jitsusama/lgwt/app/pkg/testing"
	"strings"
	"testing"
	"time"
)

var dummyAlerter = &spiedBlindAlerter{}
var dummyStdout = &bytes.Buffer{}

func TestCli(t *testing.T) {
	t.Run("records chris winning", func(t *testing.T) {
		stdin := strings.NewReader("Chris wins\n")
		store := test.StubbedPlayerStore{}

		cli := cli.NewCli(&store, stdin, dummyStdout, dummyAlerter)
		cli.PlayPoker()

		test.AssertPlayerWin(t, &store, "Chris")
	})
	t.Run("records cleo winning", func(t *testing.T) {
		stdin := strings.NewReader("Cleo wins\n")
		store := test.StubbedPlayerStore{}

		cli := cli.NewCli(&store, stdin, dummyStdout, dummyAlerter)
		cli.PlayPoker()

		test.AssertPlayerWin(t, &store, "Cleo")
	})
	t.Run("schedules printing of updated blind values", func(t *testing.T) {
		stdin := strings.NewReader("Chris wins\n")
		store := test.StubbedPlayerStore{}
		alerter := &spiedBlindAlerter{}

		cli := cli.NewCli(&store, stdin, dummyStdout, alerter)
		cli.PlayPoker()

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{10 * time.Minute, 200},
			{20 * time.Minute, 300},
			{30 * time.Minute, 400},
			{40 * time.Minute, 500},
			{50 * time.Minute, 600},
			{60 * time.Minute, 800},
			{70 * time.Minute, 1000},
			{80 * time.Minute, 2000},
			{90 * time.Minute, 4000},
			{100 * time.Minute, 8000},
		}
		for i, expected := range cases {
			t.Run(fmt.Sprint(expected), func(t *testing.T) {
				if len(alerter.alerts) <= i {
					t.Fatalf("alert %d was not scheduled: %v", i, alerter.alerts)
				}
				assertScheduledAlert(t, alerter.alerts[i], expected)
			})
		}
	})
	t.Run("requests number of players from user", func(t *testing.T) {
		stdout := &bytes.Buffer{}
		stdin := &bytes.Buffer{}
		alerter := &spiedBlindAlerter{}
		store := &test.StubbedPlayerStore{}

		cli := cli.NewCli(store, stdin, stdout, alerter)
		cli.PlayPoker()

		actual := stdout.String()
		expected := "Please enter the number of players: "
		if actual != expected {
			t.Errorf("got %q want %q", actual, expected)
		}
	})
}

func assertScheduledAlert(t testing.TB, actual scheduledAlert, expected scheduledAlert) {
	t.Helper()
	if actual.amount != expected.amount {
		t.Errorf("amount: got %d want %d", actual.amount, expected.amount)
	}
	if actual.at != expected.at {
		t.Errorf("time: got %v want %v", actual.at, expected.at)
	}
}

type scheduledAlert struct {
	at     time.Duration
	amount int
}

func (s scheduledAlert) String() string {
	return fmt.Sprintf("%d chips at %v", s.amount, s.at)
}

type spiedBlindAlerter struct {
	alerts []scheduledAlert
}

func (s *spiedBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, scheduledAlert{duration, amount})
}
