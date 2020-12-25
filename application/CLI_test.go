package poker_test

import (
	"strings"
	"testing"

	poker "github.com/guiaramos/learn-go-with-tests/application"
)

func TestCLI(t *testing.T) {
	t.Run("record gui win from user input", func(t *testing.T) {
		in := strings.NewReader("Gui wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Gui")
	})

	t.Run("record claire win from user input", func(t *testing.T) {
		in := strings.NewReader("Claire wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Claire")
	})
}
