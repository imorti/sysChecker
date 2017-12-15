package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		log.Fatal("Your GOPATH has not been set!")
	}

	path := os.Getenv("PATH")
	gobin := filepath.Join(gopath, "bin")
	if !strings.Contains(path, gobin) {
		log.Fatalf("Your PATH does not contain %s", gobin)
	} else {
		log.Println("Looks like Go isn't configured on your machine")
	}

	log.Println("Success! Your go environment is configured properly.")
}
