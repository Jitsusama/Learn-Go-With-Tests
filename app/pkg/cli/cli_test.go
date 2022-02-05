package cli_test

import (
	"jitsusama/lgwt/app/pkg/cli"
	test "jitsusama/lgwt/app/pkg/testing"
	"strings"
	"testing"
	"time"
)

var dummyAlerter = &SpiedBlindAlerter{}

func TestCli(t *testing.T) {
	t.Run("records chris winning", func(t *testing.T) {
		stdin := strings.NewReader("Chris wins\n")
		store := test.StubbedPlayerStore{}

		cli := cli.NewCli(&store, stdin, dummyAlerter)
		cli.PlayPoker()

		test.AssertPlayerWin(t, &store, "Chris")
	})
	t.Run("records cleo winning", func(t *testing.T) {
		stdin := strings.NewReader("Cleo wins\n")
		store := test.StubbedPlayerStore{}

		cli := cli.NewCli(&store, stdin, dummyAlerter)
		cli.PlayPoker()

		test.AssertPlayerWin(t, &store, "Cleo")
	})
	t.Run("schedules printing of updated blind values", func(t *testing.T) {
		stdin := strings.NewReader("Chris wins\n")
		store := test.StubbedPlayerStore{}
		alerter := &SpiedBlindAlerter{}

		cli := cli.NewCli(&store, stdin, alerter)
		cli.PlayPoker()

		if len(alerter.alerts) != 1 {
			t.Fatal("no blind alert was scheduled")
		}
	})
}

type SpiedBlindAlerter struct {
	alerts []struct {
		scheduledAt time.Duration
		amount      int
	}
}

func (s *SpiedBlindAlerter) ScheduleAlertAt(duration time.Duration, amount int) {
	s.alerts = append(s.alerts, struct {
		scheduledAt time.Duration
		amount      int
	}{duration, amount})
}
