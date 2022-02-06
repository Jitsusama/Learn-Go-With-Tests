package game_test

import (
	"fmt"
	"io"
	"jitsusama/lgwt/app/pkg/game"
	test "jitsusama/lgwt/app/pkg/testing"
	"testing"
	"time"
)

func TestPoker(t *testing.T) {
	t.Run("schedules blind alerts for 5 player game", func(t *testing.T) {
		alerter := &spiedBlindAlerter{}
		store := &test.StubbedPlayerStore{}

		game := game.NewPokerGame(alerter, store)
		game.Start(5, io.Discard)

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
		assertSchedule(t, alerter, cases)
	})
	t.Run("schedules blind alerts for 7 player game", func(t *testing.T) {
		alerter := &spiedBlindAlerter{}
		store := &test.StubbedPlayerStore{}

		game := game.NewPokerGame(alerter, store)
		game.Start(7, io.Discard)

		cases := []scheduledAlert{
			{0 * time.Second, 100},
			{12 * time.Minute, 200},
			{24 * time.Minute, 300},
			{36 * time.Minute, 400},
		}
		assertSchedule(t, alerter, cases)
	})
	t.Run("records final winner", func(t *testing.T) {
		alerter := &spiedBlindAlerter{}
		store := &test.StubbedPlayerStore{}

		game := game.NewPokerGame(alerter, store)
		game.Finish("Ruth")

		test.AssertPlayerWin(t, store, "Ruth")
	})
}

func assertSchedule(t *testing.T, alerter *spiedBlindAlerter, cases []scheduledAlert) {
	for i, expected := range cases {
		t.Run(fmt.Sprint(expected), func(t *testing.T) {
			if len(alerter.alerts) <= i {
				t.Fatalf("alert %d was not scheduled: %v", i, alerter.alerts)
			}
			assertScheduledAlert(t, alerter.alerts[i], expected)
		})
	}
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

func (s *spiedBlindAlerter) ScheduleAlertAt(d time.Duration, a int, t io.Writer) {
	s.alerts = append(s.alerts, scheduledAlert{d, a})
}
