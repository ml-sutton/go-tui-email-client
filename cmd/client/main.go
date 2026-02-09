package main

import (
	"log"

	"github.com/ml-sutton/go-tui-email-client/lifecycle"
)

func main() {
	err := lifecycle.RunClient()
	log.Fatal(err)
}
