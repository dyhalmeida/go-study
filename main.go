package main

import (
	"command-line/app"
	"log"
	"os"
)

func main() {
	app := app.Main()

	error := app.Run(os.Args)
	if error != nil {
		log.Fatal(error)
	}
}