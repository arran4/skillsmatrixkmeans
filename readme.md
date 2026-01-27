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

The tool outputs a JSON array containing the identified clusters. Each cluster object provides the following details:

-   `label`: An identifier for the group derived from its top skills (e.g., "Go, Python, SQL").
-   `center`: A map showing the average skill levels for the group.
-   `members`: A list of the individuals belonging to the group.
-   `cohesion`: The average distance of members from the group's center. Lower values indicate a more cohesive group (members are more similar to each other).

```json
[
  {
    "label": "Go, Python, SQL",
    "center": {
      "CSS": 1.5,
      "Go": 4.5,
      "Python": 4.5,
      "React": 1.5,
      "SQL": 4.5
    },
    "members": [
      "Alice",
      "Bob"
    ],
    "cohesion": 1.12
  },
  {
    "label": "CSS, React, Go",
    "center": {
      "CSS": 4,
      "Go": 2,
      "Python": 2,
      "React": 4,
      "SQL": 2
    },
    "members": [
      "Charlie",
      "Diana",
      "Eve"
    ],
    "cohesion": 1.79
  }
]
```

**Interpretation:**

*   **Go, Python, SQL**: The center shows high values for Go, Python, and SQL. This group contains Alice and Bob. This represents the **Backend** group.
*   **CSS, React, Go**: The center shows high values for React and CSS. This group contains Charlie, Diana, and Eve (who is a generalist but closer to this group). This represents the **Frontend/Fullstack** group.
