package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/guiaramos/learn-go-with-tests/application"
)

const dbFileName = "game.db.json"

func main() {
	store, close, err := poker.FileSystemPlayerStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Let's play poker")
	fmt.Println("Type {name} wins to record a win")

	game := poker.BlindAlerterFunc(poker.StdOutAlerter)

	cli := poker.NewCLI(store, os.Stdin, game)
	cli.PlayPoker()
}
