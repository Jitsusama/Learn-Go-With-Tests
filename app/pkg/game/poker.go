package game

import (
	"io"
	"time"
)

func NewPokerGame(alerter BlindAlerter, store PlayerStore) *Poker {
	return &Poker{alerter, store}
}

type Poker struct {
	alerter BlindAlerter
	store   PlayerStore
}

func (p *Poker) Start(players int, alertDestination io.Writer) {
	blindIncrement := time.Duration(5+players) * time.Minute
	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
	blindTime := 0 * time.Second
	for _, blind := range blinds {
		p.alerter.ScheduleAlertAt(blindTime, blind, alertDestination)
		blindTime = blindTime + blindIncrement
	}
}

func (p *Poker) Finish(winner string) {
	p.store.IncrementScore(winner)
}
