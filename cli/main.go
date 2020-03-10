package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "sport event cli",
}

func Execute() {
	timingCmd := &cobra.Command{
		Use:   "timing",
		Short: "timing functions",
	}
	athleteCmd := &cobra.Command{
		Use:   "athlete",
		Short: "athlete functions",
	}

	timingCmd.AddCommand(timingListCmd)
	timingCmd.AddCommand(timingGetCmd)
	timingCmd.AddCommand(timingAddCmd)

	athleteCmd.AddCommand(athleteListCmd)
	athleteCmd.AddCommand(athleteGetCmd)
	athleteCmd.AddCommand(athleteAddCmd)

	rootCmd.AddCommand(athleteCmd)
	rootCmd.AddCommand(timingCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
