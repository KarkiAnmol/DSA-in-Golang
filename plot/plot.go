package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Attempt to open the CSV file containing the performance data.
	file, err := os.Open("../performanceMeasurement/performance_data.csv")
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	// Create a CSV reader for the opened file.
	reader := csv.NewReader(file)

	// Read all the records.
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV data: %v", err)
	}

	// Initialize the plot.
	p := plot.New()

	p.Title.Text = "Sorting Algorithm Performance"
	p.X.Label.Text = "Input Size"
	p.Y.Label.Text = "Execution Time (ms)"

	// Prepare plot points for insertion and merge sort.
	ptsInsertion := make(plotter.XYs, 0)
	ptsMerge := make(plotter.XYs, 0)

	// Process the records and populate the plot points.
	for i, record := range records {
		// Skip the header row.
		if i == 0 {
			continue
		}

		size, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatalf("Failed to parse size from CSV: %v", err)
		}
		time, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			log.Fatalf("Failed to parse execution time from CSV: %v", err)
		}

		if record[0] == "Insertion Sort" {
			ptsInsertion = append(ptsInsertion, plotter.XY{X: float64(size), Y: time})
		} else if record[0] == "Merge Sort" {
			ptsMerge = append(ptsMerge, plotter.XY{X: float64(size), Y: time})
		}
	}

	// Add line plots to the plot for each sorting algorithm.
	err = plotutil.AddLinePoints(p,
		"Insertion Sort", ptsInsertion,
		"Merge Sort", ptsMerge)
	if err != nil {
		log.Fatalf("Failed to add line points to plot: %v", err)
	}

	// Save the plot to a file.
	if err := p.Save(10*vg.Inch, 8*vg.Inch, "sorting_performance.png"); err != nil {
		log.Fatalf("Failed to save the plot to a PNG file: %v", err)
	}
}
