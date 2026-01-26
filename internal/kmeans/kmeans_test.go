package kmeans

import (
	"testing"
)

func TestEuclideanDistance(t *testing.T) {
	tests := []struct {
		name string
		a    []float64
		b    []float64
		want float64
	}{
		{"Same point", []float64{1, 1}, []float64{1, 1}, 0.0},
		{"Distance 1", []float64{0, 0}, []float64{1, 0}, 1.0},
		{"Distance 5", []float64{0, 0}, []float64{3, 4}, 5.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := euclideanDistance(tt.a, tt.b); got != tt.want {
				t.Errorf("euclideanDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKMeans(t *testing.T) {
	// Simple test: 2 obvious clusters
	// Cluster 1: (0,0), (1,1)
	// Cluster 2: (10,10), (11,11)
	points := []Point{
		{ID: "A", Vector: []float64{0, 0}},
		{ID: "B", Vector: []float64{1, 1}},
		{ID: "C", Vector: []float64{10, 10}},
		{ID: "D", Vector: []float64{11, 11}},
	}

	k := 2
	clusters, err := KMeans(points, k, 10, nil)
	if err != nil {
		t.Fatalf("KMeans failed: %v", err)
	}

	if len(clusters) != k {
		t.Errorf("Expected %d clusters, got %d", k, len(clusters))
	}

	// Verify points separation
	// We expect A & B in one cluster, C & D in another.

	// Since we don't know which cluster is which (random init),
	// we just check if A and B are together, and C and D are together.

	// Helper to find cluster index for a point ID
	findCluster := func(id string) int {
		for i, c := range clusters {
			for _, p := range c.Points {
				if p.ID == id {
					return i
				}
			}
		}
		return -1
	}

	idxA := findCluster("A")
	idxB := findCluster("B")
	idxC := findCluster("C")
	idxD := findCluster("D")

	if idxA == -1 || idxB == -1 || idxC == -1 || idxD == -1 {
		t.Fatalf("Some points were not assigned to any cluster")
	}

	if idxA != idxB {
		t.Errorf("Expected A and B to be in the same cluster, but got %d and %d", idxA, idxB)
	}
	if idxC != idxD {
		t.Errorf("Expected C and D to be in the same cluster, but got %d and %d", idxC, idxD)
	}
	if idxA == idxC {
		t.Errorf("Expected A and C to be in different clusters, but they are in the same one")
	}
}
