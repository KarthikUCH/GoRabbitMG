package main

import (
	"os"
	"strings"

	"github.com/KarthikUCH/goproducer/produce"
)

func main() {
	body := bodyFrom(os.Args) // To receive the message sent from command line
	produce.PostMessage(body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
