# Skills Matrix K-Means

A simple command-line tool to cluster entities (e.g., people) based on a skills matrix using the K-Means algorithm.

## Description

This tool reads a CSV file containing a skills matrix where the first column is an identifier (e.g., Name) and subsequent columns represent skill levels (e.g., 1-5). It then groups these entities into `k` clusters based on the similarity of their skill profiles.

## Installation

### Prerequisites

- Go 1.24 or later

### Build

```bash
go build -o skills-kmeans ./cmd/skills-kmeans
```

## Usage

```bash
./skills-kmeans -file <input_csv> -k <number_of_clusters> [-out <output_json>]
```

- `-file`: Path to the input CSV file (required).
- `-k`: Number of clusters (default: 3).
- `-out`: Path to the output JSON file. If omitted, prints to stdout.

### Example CSV (`data.csv`)

```csv
Name,Go,Python,SQL,Communication
Alice,5,3,4,4
Bob,2,5,3,3
Charlie,4,2,5,2
Diana,1,5,2,4
Eve,5,4,4,5
Frank,2,2,2,2
```

### Running the example

```bash
./skills-kmeans -file data.csv -k 2
```

## Output

The output is a JSON representation of the clusters, showing the centroid (average skill profile of the cluster) and the points (entities) belonging to that cluster.

```json
[
  {
    "Centroid": [ ... ],
    "Points": [
      {
        "ID": "Alice",
        "Vector": [5, 3, 4, 4]
      },
      ...
    ]
  },
  ...
]
```
