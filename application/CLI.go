package poker

import (
	"bufio"
	"io"
	"strings"
)

// CLI is a collection of methods for use the CLI
type CLI struct {
	playerStore PlayerStore
	in          *bufio.Scanner
}

// NewCLI creates a new CLI collection
func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playerStore: store,
		in:          bufio.NewScanner(in),
	}
}

// PlayPoker triggers the start of the game
func (c *CLI) PlayPoker() {
	userInput := c.readLine()
	c.playerStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (c *CLI) readLine() string {
	c.in.Scan()
	return c.in.Text()
}
