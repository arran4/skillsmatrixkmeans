package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/example/skills-matrix-kmeans/internal/kmeans"
)

func main() {
	inputFile := flag.String("file", "", "Path to the input CSV file")
	k := flag.Int("k", 3, "Number of clusters")
	outputFile := flag.String("out", "", "Path to the output JSON file (optional)")
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Error: Please provide an input file using -file")
		flag.Usage()
		os.Exit(1)
	}

	points, err := readCSV(*inputFile)
	if err != nil {
		fmt.Printf("Error reading CSV: %v\n", err)
		os.Exit(1)
	}

	clusters, err := kmeans.KMeans(points, *k, 100)
	if err != nil {
		fmt.Printf("Error running K-Means: %v\n", err)
		os.Exit(1)
	}

	output, err := json.MarshalIndent(clusters, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling output: %v\n", err)
		os.Exit(1)
	}

	if *outputFile != "" {
		err := os.WriteFile(*outputFile, output, 0644)
		if err != nil {
			fmt.Printf("Error writing to output file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Clustering complete. Results written to %s\n", *outputFile)
	} else {
		fmt.Println(string(output))
	}
}

func readCSV(filename string) ([]kmeans.Point, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 2 {
		return nil, fmt.Errorf("CSV file must have at least a header and one data row")
	}

	// Assume first row is header, skip it.
	// Assume first column is identifier (Name).
	// Subsequent columns are numeric skills.

	var points []kmeans.Point

	for i, record := range records {
		if i == 0 {
			continue // Skip header
		}

		if len(record) < 2 {
			continue // Skip invalid rows
		}

		id := record[0]
		var vector []float64

		for j := 1; j < len(record); j++ {
			val, err := strconv.ParseFloat(record[j], 64)
			if err != nil {
				return nil, fmt.Errorf("invalid number at row %d, col %d: %v", i+1, j+1, err)
			}
			vector = append(vector, val)
		}

		points = append(points, kmeans.Point{
			ID:     id,
			Vector: vector,
		})
	}

	return points, nil
}
