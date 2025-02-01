package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/montanaflynn/stats"
)

// Read CSV file and extract data in the right format
func readCSV(filePath string) []stats.Series {
	file, _ := os.Open(filePath)
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	// Create stats.Series which is what stats.LinearRegression takes as an input for each part of the CSV
	datasets := make([]stats.Series, 4)

	//Iterate through rows
	for i, row := range records {

		// Skip header row
		if i == 0 {
			continue
		}

		//For each row, convert each column to a float
		xVal, _ := strconv.ParseFloat(row[0], 64)
		yVal1, _ := strconv.ParseFloat(row[1], 64)
		yVal2, _ := strconv.ParseFloat(row[2], 64)
		yVal3, _ := strconv.ParseFloat(row[3], 64)
		xVal4, _ := strconv.ParseFloat(row[4], 64)
		yVal4, _ := strconv.ParseFloat(row[5], 64)

		//append varible to store them
		datasets[0] = append(datasets[0], stats.Coordinate{X: xVal, Y: yVal1})
		datasets[1] = append(datasets[1], stats.Coordinate{X: xVal, Y: yVal2})
		datasets[2] = append(datasets[2], stats.Coordinate{X: xVal, Y: yVal3})
		datasets[3] = append(datasets[3], stats.Coordinate{X: xVal4, Y: yVal4})
	}

	return datasets
}

func main() {
	//Set file path
	filePath := "Anscombe_quartet_data.csv"

	// Read CSV and get datasets
	datasets := readCSV(filePath)

	// create slices with lenth of 4 for the results and timing
	results := make([][2]float64, 4)
	timings := make([]float64, 4)

	//Iterate through each row
	for i := range datasets {
		//start time
		start := time.Now()

		//perform linear regression
		params, _ := stats.LinearRegression(datasets[i])
		//end time
		elapsed := time.Since(start).Seconds()
		//store results
		results[i] = [2]float64{params[1].Y, params[0].Y} // slope and intercept
		timings[i] = elapsed
	}

	// Print results
	for i, res := range results {
		fmt.Printf("Dataset %d: Slope = %.5f, Intercept = %.5f\n", i+1, res[0], res[1])
		fmt.Printf("Time = %.9f seconds\n", timings[i])
	}
}
