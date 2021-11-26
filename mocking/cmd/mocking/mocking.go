package main

import (
	"jitsusama/lgwt/mocking/countdown"
	"os"
)

func main() {
	sleeper := &countdown.DefaultSleeper{}
	countdown.Countdown(os.Stdout, sleeper)
}
