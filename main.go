package main

import (
	"fmt"
	"os"
)

var (
	// Version is the current version of the application.
	Version = "main"
)

func main() {
	fmt.Fprintf(os.Stdout, "GoReleaser supports ko! Version: %s", Version)
}
