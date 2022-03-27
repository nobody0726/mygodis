package main

import (
	"log"

	"github.com/nobody0726/mygodis/tcp"
)

func main() {
	err := tcp.ListenAndServeWithSignal(&tcp.Config{
		Address: ":8000",
	}, tcp.NewEchoHandler())
	if err != nil {
		log.Fatal(err)
	}
}
