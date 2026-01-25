# Skills Matrix K-Means

A command-line tool that leverages the K-Means algorithm to cluster entities (such as team members or job candidates) based on their skills matrix.

## Why use this tool?

Clustering skill profiles can uncover patterns in your workforce or candidate pool that aren't immediately obvious from raw data.

### Use Cases

*   **Team Formation**: ensuring agile squads have a balanced mix of skills by identifying different developer archetypes (e.g., "Frontend Specialists", "Backend Specialists", "Generalists").
*   **Learning & Development**: Grouping employees with similar skill gaps to provide targeted training sessions.
*   **Recruitment Analysis**: Categorizing candidates into buckets to see which roles are easiest or hardest to fill based on the current pipeline.

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

## Detailed Example

### Scenario

Imagine you have a list of developers with self-assessed scores (1-5) for various technologies. You want to group them into 2 clusters to identify the primary "Backend" and "Frontend" groups.

### 1. Prepare Data (`skills.csv`)

```csv
Name,Go,Python,React,SQL,CSS
Alice,5,4,1,5,2
Bob,4,5,2,4,1
Charlie,1,2,5,2,5
Diana,2,1,4,1,4
Eve,3,3,3,3,3
```

### 2. Run the Tool

```bash
./skills-kmeans -file skills.csv -k 2
```

### 3. Analyze Output

The tool outputs a JSON array where each object is a cluster containing a `Centroid` (the "average" profile of that cluster) and the list of `Points` (people) in that cluster.

```json
[
  {
    "Centroid": [
      2,
      2,
      4,
      2,
      4
    ],
    "Points": [
      {
        "ID": "Charlie",
        "Vector": [1, 2, 5, 2, 5]
      },
      {
        "ID": "Diana",
        "Vector": [2, 1, 4, 1, 4]
      },
      {
        "ID": "Eve",
        "Vector": [3, 3, 3, 3, 3]
      }
    ]
  },
  {
    "Centroid": [
      4.5,
      4.5,
      1.5,
      4.5,
      1.5
    ],
    "Points": [
      {
        "ID": "Alice",
        "Vector": [5, 4, 1, 5, 2]
      },
      {
        "ID": "Bob",
        "Vector": [4, 5, 2, 4, 1]
      }
    ]
  }
]
```

**Interpretation:**

*   **Cluster 1 (Top)**: The centroid has high values for the 3rd and 5th skills (React, CSS). This group contains Charlie, Diana, and Eve (who is a generalist but closer to this group). This represents the **Frontend/Fullstack** group.
*   **Cluster 2 (Bottom)**: The centroid has high values for the 1st, 2nd, and 4th skills (Go, Python, SQL). This group contains Alice and Bob. This represents the **Backend** group.
