package cli

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/example/skills-matrix-kmeans/cmd"
	"github.com/example/skills-matrix-kmeans/internal/kmeans"
)

type OutputCluster struct {
	Label    string             `json:"label"`
	Center   map[string]float64 `json:"center"`
	Members  []string           `json:"members"`
	Cohesion float64            `json:"cohesion"`
}

// Run is a subcommand `skills-kmeans`
//
// Flags:
//
//	inputFile:  -file (default: "") Path to the input CSV file
//	k:          -k    (default: 3)  Number of clusters
//	outputFile: -out  (default: "") Path to the output JSON file (optional)
func Run(inputFile string, k int, outputFile string) error {
	if inputFile == "" {
		fmt.Println("Error: Please provide an input file using -file")
		return cmd.ErrPrintHelp
	}
	return run(inputFile, outputFile, k, nil)
}

func run(inputFile, outputFile string, k int, rng *rand.Rand) error {
	points, headers, err := readCSV(inputFile)
	if err != nil {
		return fmt.Errorf("reading CSV: %w", err)
	}

	clusters, err := kmeans.KMeans(points, k, 100, rng)
	if err != nil {
		return fmt.Errorf("running K-Means: %w", err)
	}

	var outputClusters []OutputCluster
	for _, c := range clusters {
		center := make(map[string]float64)
		for j, val := range c.Centroid {
			if j < len(headers) {
				center[headers[j]] = val
			}
		}

		var members []string
		for _, p := range c.Points {
			members = append(members, p.ID)
		}

		outputClusters = append(outputClusters, OutputCluster{
			Label:    generateLabel(center),
			Center:   center,
			Members:  members,
			Cohesion: c.Cohesion,
		})
	}

	output, err := json.MarshalIndent(outputClusters, "", "  ")
	if err != nil {
		return fmt.Errorf("marshaling output: %w", err)
	}

	if outputFile != "" {
		err := os.WriteFile(outputFile, output, 0644)
		if err != nil {
			return fmt.Errorf("writing to output file: %w", err)
		}
		fmt.Printf("Clustering complete. Results written to %s\n", outputFile)
	} else {
		fmt.Println(string(output))
	}
	return nil
}

func readCSV(filename string) ([]kmeans.Point, []string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	if len(records) < 2 {
		return nil, nil, fmt.Errorf("CSV file must have at least a header and one data row")
	}

	// Assume first row is header.
	// Assume first column is identifier (Name).
	// Subsequent columns are numeric skills.
	headers := records[0][1:]

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
				return nil, nil, fmt.Errorf("invalid number at row %d, col %d: %v", i+1, j+1, err)
			}
			vector = append(vector, val)
		}

		points = append(points, kmeans.Point{
			ID:     id,
			Vector: vector,
		})
	}

	return points, headers, nil
}

func generateLabel(center map[string]float64) string {
	type kv struct {
		Key   string
		Value float64
	}

	var ss []kv
	for k, v := range center {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			return ss[i].Key < ss[j].Key
		}
		return ss[i].Value > ss[j].Value
	})

	var topSkills []string
	for i, kv := range ss {
		if i >= 3 {
			break
		}
		topSkills = append(topSkills, kv.Key)
	}

	return strings.Join(topSkills, ", ")
}
