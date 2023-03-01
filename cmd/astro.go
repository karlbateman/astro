// Package main contains the application entrypoint.
package main

import (
	"fmt"
	"os"

	"github.com/karlbateman/astro"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func run() error {
	a := astro.New()
	return a.Run()
}
