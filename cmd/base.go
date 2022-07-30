package cmd

import (
	"github.com/spf13/cobra"
)

var base = &cobra.Command{
	Use:     "main",
	Short:   "Music Tribe Interview Task",
	Long:    "A small CLI to learn / consume data from a text file and generate random text using trigrams",
	Version: "0.0.1",
}

func Execute() error {
	return base.Execute()
}
