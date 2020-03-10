package main

import (
	"encoding/json"
	"fmt"
	"github.com/Test214-dev/mooncascade-task/models"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var athleteListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all athletes",
	Run: func(cmd *cobra.Command, args []string) {
		body, err := getAsset("/athletes/")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		res := make([]models.Athlete, 0)
		if err := json.Unmarshal(body, &res); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		prettyPrintAthletes(res)
	},
}
var athleteGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get athlete by ID",
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		body, err := getAsset(fmt.Sprintf("/athletes/%s", id))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		res := models.Athlete{}
		if err := json.Unmarshal(body, &res); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		prettyPrintAthlete(res, true)
	},
}
var athleteAddCmd = &cobra.Command{
	Use:   "add [athlete full name] [athlete start number]",
	Short: "add a new athlete",
	Run: func(cmd *cobra.Command, args []string) {
		fullName := args[0]
		startNumber, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Unable to parse start number. Is it an integer?")
			os.Exit(1)
		}
		athlete := models.Athlete{
			FullName:    fullName,
			StartNumber: startNumber,
		}

		bbody, err := json.Marshal(&athlete)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		chipID, err := createAsset("/athletes", bbody)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Athlete chip ID is: %s\n", chipID)
	},
}

func prettyPrintAthlete(a models.Athlete, first bool) {
	if first {
		fmt.Println("|------------------------------------------------------------------------|")
	}
	fmt.Printf("|%21s |\t%40s |\n", "Athlete chip ID", a.ChipID)
	fmt.Printf("|%21s |\t%40s |\n", "Athlete name", a.FullName)
	fmt.Printf("|%21s |\t%40d |\n", "Athlete start number", a.StartNumber)
	fmt.Println("|------------------------------------------------------------------------|")
}

func prettyPrintAthletes(athletes []models.Athlete) {
	fmt.Println("|------------------------------------------------------------------------|")
	for _, a := range athletes {
		prettyPrintAthlete(a, false)
	}
}
