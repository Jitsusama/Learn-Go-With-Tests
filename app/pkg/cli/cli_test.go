package cli_test

import (
	"jitsusama/lgwt/app/pkg/cli"
	test "jitsusama/lgwt/app/pkg/testing"
	"strings"
	"testing"
)

func TestCli(t *testing.T) {
	t.Run("records chris winning", func(t *testing.T) {
		stdin := strings.NewReader("Chris wins\n")
		store := test.StubbedPlayerStore{}

		cli := cli.NewCli(&store, stdin)
		cli.PlayPoker()

		test.AssertPlayerWin(t, &store, "Chris")
	})
	t.Run("records cleo winning", func(t *testing.T) {
		stdin := strings.NewReader("Cleo wins\n")
		store := test.StubbedPlayerStore{}

		cli := cli.NewCli(&store, stdin)
		cli.PlayPoker()

		test.AssertPlayerWin(t, &store, "Cleo")
	})
}
