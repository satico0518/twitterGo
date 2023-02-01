package main

import (
	"log"

	"github.com/satico0518/twitterGo/bd"
	"github.com/satico0518/twitterGo/handlers"
)

func main() {
	if !bd.CheckConnetion() {
		log.Fatal("Db no connected")
		return
	}

	handlers.Handlers()
}
