package main

import (
	"os"

	"github.com/mitchellh/colorstring"
)

func main() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		colorstring.Println("[red] Uh oh, looks like GOPATH is not set.")
	} else {
		colorstring.Println("[green] Success! Well looks like GOPATH is configured properly. FYI it's here: " + gopath)
	}

	goroot := os.Getenv("GOROOT")

	if goroot == "" {
		colorstring.Println("[red] Uh oh, Looks like GOROOT is not set.")
	} else {
		colorstring.Println("[green] Well looks like GOROOT is configured properly. FYI it's here: " + goroot)
	}

}
