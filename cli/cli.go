package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Test214-dev/mooncascade-task/models"
	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "sport event cli",
}

func Execute() {
	pointCmd := &cobra.Command{
		Use:   "timings",
		Short: "timing point functions",
	}
	athleteCmd := &cobra.Command{
		Use:   "athlete",
		Short: "athlete functions",
	}
	athleteListCmd := &cobra.Command{
		Use:   "list",
		Short: "list all athletes",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := http.Get("http://localhost:8080/api/athletes/")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				fmt.Printf("Request failed with code %d\n", resp.StatusCode)
				os.Exit(1)
			}
			body, err := ioutil.ReadAll(resp.Body)
			res := make([]models.Athlete, 0)
			if err := json.Unmarshal(body, &res); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			PrettyPrintAthletes(res)
		},
	}
	athleteGetCmd := &cobra.Command{
		Use:   "get",
		Short: "get athlete by ID",
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			resp, err := http.Get(fmt.Sprintf("http://localhost:8080/api/athletes/%s", id))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				fmt.Printf("Request failed with code %d\n", resp.StatusCode)
				os.Exit(1)
			}
			body, err := ioutil.ReadAll(resp.Body)
			res := models.Athlete{}
			if err := json.Unmarshal(body, &res); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			PrettyPrintAthlete(res, true)
		},
	}
	athleteAddCmd := &cobra.Command{
		Use:   "add [athlete full name] [athlete start number]",
		Short: "add a new athlete",
		Run: func(cmd *cobra.Command, args []string) {
			fullName := args[0]
			startNumber, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Printf("Unable to parse start number. Is it an integer number?")
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

			resp, err := http.Post("http://localhost:8080/api/athletes/", "application/json", bytes.NewReader(bbody))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusCreated {
				fmt.Printf("Request failed with code %d %s\n", resp.StatusCode, GetResponseError(resp))
				os.Exit(1)
			}

			fmt.Printf("Athlete chip ID is: %s\n", resp.Header.Get("Location"))
		},
	}

	pointGetCmd := &cobra.Command{
		Use:   "get",
		Short: "get timing by ID",
		Run: func(cmd *cobra.Command, args []string) {
			id := args[0]
			resp, err := http.Get(fmt.Sprintf("http://localhost:8080/api/timings/%s", id))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				fmt.Printf("Request failed with code %d\n", resp.StatusCode)
				os.Exit(1)
			}
			body, err := ioutil.ReadAll(resp.Body)
			res := models.Timing{}
			if err := json.Unmarshal(body, &res); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			PrettyPrintPoint(res, true)
		},
	}
	pointListCmd := &cobra.Command{
		Use:   "list",
		Short: "list all points",
		Run: func(cmd *cobra.Command, args []string) {
			resp, err := http.Get("http://localhost:8080/api/timings/")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				fmt.Printf("Request failed with code %d\n", resp.StatusCode)
				os.Exit(1)
			}
			body, err := ioutil.ReadAll(resp.Body)
			res := make([]models.Timing, 0)
			if err := json.Unmarshal(body, &res); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			PrettyPrintPints(res)
		},
	}
	pointAddCmd := &cobra.Command{
		Use:   "add [athlete chip id] [point type]",
		Short: "add a timing point for athlete",
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

			resp, err := http.Post("http://localhost:8080/api/timings/", "application/json", bytes.NewReader(bbody))
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusCreated {
				fmt.Printf("Request failed with code %d %s\n", resp.StatusCode, GetResponseError(resp))
				os.Exit(1)
			}

			fmt.Printf("Point ID is: %s\n", resp.Header.Get("Location"))
		},
	}
	pointCmd.AddCommand(pointListCmd)
	athleteCmd.AddCommand(athleteListCmd)
	athleteCmd.AddCommand(athleteGetCmd)
	athleteCmd.AddCommand(athleteAddCmd)
	pointCmd.AddCommand(pointGetCmd)
	pointCmd.AddCommand(pointAddCmd)
	rootCmd.AddCommand(athleteCmd)
	rootCmd.AddCommand(pointCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func PrettyPrintAthletes(athletes []models.Athlete) {
	fmt.Println("|------------------------------------------------------------------------|")
	for _, a := range athletes {
		PrettyPrintAthlete(a, false)
	}
}

func PrettyPrintPoint(p models.Timing, first bool) {
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

func PrettyPrintPints(points []models.Timing) {
	fmt.Println("|------------------------------------------------------------------------|")
	for _, p := range points {
		PrettyPrintPoint(p, false)
	}
}

func PrettyPrintAthlete(a models.Athlete, first bool) {
	if first {
		fmt.Println("|------------------------------------------------------------------------|")
	}
	fmt.Printf("|%21s |\t%40s |\n", "Athlete chip ID", a.ChipID)
	fmt.Printf("|%21s |\t%40s |\n", "Athlete name", a.FullName)
	fmt.Printf("|%21s |\t%40d |\n", "Athlete start number", a.StartNumber)
	fmt.Println("|------------------------------------------------------------------------|")
}

func GetResponseError(response *http.Response) string {
	data, _ := ioutil.ReadAll(response.Body)
	m := models.AppError{}
	if err := json.Unmarshal(data, &m); err != nil {
		return ""
	}

	return m.Error
}

func main() {
	Execute()
}
