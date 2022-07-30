package cmd

import "github.com/spf13/cobra"

var learn = &cobra.Command{
	Use:   "learn",
	Short: "Learn the given text file",
	Run: func(cmd *cobra.Command, args []string) {
		// use learn func from package
	},
}

func init() {
	base.AddCommand(learn)
}
