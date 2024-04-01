package main

import (
	"log"

	"github.com/hsmtkk/aukabucomgo/cmd"
)

func main() {
	rootCmd := cmd.RootCmd
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
