package main

import (
	"github.com/rewolf/wordle-golf-linebot/internal/daemon"
	"log"
)

func main() {
	var err error

	if botDaemon, err := daemon.New(); err == nil {
		err = botDaemon.Run()
	}

	if err != nil {
		log.Fatalln(err)
	}
}
