package main

import (
	"log"

	swissknife "github.com/Sagleft/swiss-knife"
	"github.com/Sagleft/tgfun"
)

func main() {
	_, err := tgfun.NewFunnel(
		getData(),
		getScript(),
	)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("app started")
	swissknife.RunInBackground()
}
