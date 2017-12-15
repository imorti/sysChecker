package main

import (
	"os"

	"github.com/mitchellh/colorstring"
)

func main() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		colorstring.Println("[red]Looks like Go isn't configured on your machine. GOPATH is not set.")
	} else {
		colorstring.Println("[green] Success! Your go environment is configured properly.")
	}

}
