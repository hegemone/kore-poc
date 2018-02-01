package main

import (
	"log"

	"github.com/dahendel/kore-poc/buffalo_poc/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
