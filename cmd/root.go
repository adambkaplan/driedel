package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "driedel",
		Short: "A command line to run a driedel game from your terminal!",
		Long: `Driedel is a command line that lets you run the Hanukah driedel game from an
interactive terminal`,
	}
)

func init() {
	rootCmd.AddCommand(spinCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
