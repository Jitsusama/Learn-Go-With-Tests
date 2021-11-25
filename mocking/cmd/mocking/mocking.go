package main

import (
	"jitsusama/lgwt/mocking/countdown"
	"os"
)

func main() {
	countdown.Countdown(os.Stdout)
}
