package main

import (
	"os"
	"time"

	"github.com/jitsusama/lgwt/maths/clockface"
)

func main() {
	t := time.Now()
	clockface.SvgWriter(os.Stdout, t)
}
