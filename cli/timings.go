package main

import (
	"encoding/json"
	"fmt"
	"github.com/Test214-dev/mooncascade-task/models"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var timingGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get timing by ID",
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		body, err := getAsset(fmt.Sprintf("/timings/%s", id))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		res := models.Timing{}
		if err := json.Unmarshal(body, &res); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		prettyPrintTiming(res, true)
	},
}
var timingListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all timings",
	Run: func(cmd *cobra.Command, args []string) {
		body, err := getAsset("/timings/")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		res := make([]models.Timing, 0)
		if err := json.Unmarshal(body, &res); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		prettyPrintTimings(res)
	},
}
var timingAddCmd = &cobra.Command{
	Use:   "add [athlete chip id] [point type]",
	Short: "Add a timing point for athlete. Timestamp is set automatically to current time.",
	Run: func(cmd *cobra.Command, args []string) {
		chipID := args[0]
		pointID := args[1]
		point := models.AddTimingRequest{
			PointID:   pointID,
			Timestamp: time.Now(),
			ChipID:    chipID,
		}

		bbody, err := json.Marshal(&point)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		timingID, err := createAsset("/timings", bbody)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Timing ID is: %s\n", timingID)
	},
}

func prettyPrintTiming(p models.Timing, first bool) {
	if first {
		fmt.Println("|------------------------------------------------------------------------|")
	}
	fmt.Printf("|%21s |\t%40s |\n", "Athlete chip ID", p.ChipID)
	fmt.Printf("|%21s |\t%40s |\n", "Athlete name", p.FullName)
	fmt.Printf("|%21s |\t%40d |\n", "Athlete start number", p.StartNumber)
	fmt.Printf("|%21s |\t%40s |\n", "Point ID", p.PointID)
	fmt.Printf("|%21s |\t%40s |\n", "Timing ID", p.TimingID)
	fmt.Printf("|%21s |\t%40s |\n", "Timestamp", p.Timestamp)
	fmt.Println("|------------------------------------------------------------------------|")
}

func prettyPrintTimings(points []models.Timing) {
	fmt.Println("|------------------------------------------------------------------------|")
	for _, p := range points {
		prettyPrintTiming(p, false)
	}
}
