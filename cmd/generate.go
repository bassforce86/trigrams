package cmd

import (
	"log"
	"os"
	"sort"

	"github.com/bassforce86/trigrams/pkg/ngram"
	"github.com/spf13/cobra"
)

var generate = &cobra.Command{
	Use:   "gen",
	Short: "Generate from learned text",
	Run: func(cmd *cobra.Command, args []string) {
		bytes, err := os.ReadFile("files/pride-and-prejudice.txt")
		if err != nil {
			log.Fatal(err)
		}
		trigram := ngram.BuildTrigram(string(bytes))
		sort.Sort(trigram)
		for _, trgrm := range trigram {
			log.Printf("%d [ %s ]", trgrm.Frequency, trgrm.Words)
		}
	},
}

func init() {
	base.AddCommand(generate)
}
