package main

import (
	"log"

	"github.com/ml-sutton/go-tui-email-client/internal/lifecycle"
)

func main() {
	err := lifecycle.RunClient()
	log.Fatal(err)
}
