package cmd

import (
	"fmt"

	"github.com/adambkaplan/driedel/pkg/driedel"
	"github.com/spf13/cobra"
)

var spinCmd = &cobra.Command{
	Use:   "spin",
	Short: "Spin the driedel",
	Long:  "Spins the driedel and prints the result",
	Run:   runSpinCmd,
}

func runSpinCmd(cmd *cobra.Command, args []string) {
	driedel := driedel.NewDriedel()
	result := driedel.Spin()
	fmt.Printf("You spun %s!\n", result)
}
