package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var references = &cobra.Command{
	Use:   "references",
	Short: "Lists all references for this cli tool",
	Run: func(cmd *cobra.Command, args []string) {
		for _, ref := range refs() {
			fmt.Printf("URL: %s\t- %s\n", ref.URL, ref.Reason)
		}
	},
}

func init() {
	base.AddCommand(references)
}

type Reference struct {
	URL    string
	Reason string
}

func refs() []Reference {
	return []Reference{
		{URL: "https://github.com/msandim/trigram", Reason: "Used to aid base understanding of trigram parsing"},
		{URL: "https://github.com/spf13/cobra", Reason: "Used for creating cli"},
	}
}
