package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var base = &cobra.Command{
	Use:     "trigram",
	Short:   "Generate new text using trigrams from existing text",
	Long:    "A small CLI to learn / consume data from a text file and generate random text using trigrams",
	Version: "0.0.1",
}

func Execute() {
	if err := base.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
}
