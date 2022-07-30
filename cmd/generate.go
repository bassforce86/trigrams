package cmd

import "github.com/spf13/cobra"

var generate = &cobra.Command{
	Use:   "gen",
	Short: "Generate from learned text",
	Run: func(cmd *cobra.Command, args []string) {
		// use generate func from package
	},
}

func init() {
	base.AddCommand(generate)
}
