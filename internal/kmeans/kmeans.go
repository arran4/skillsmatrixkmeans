package kmeans

import (
	"math"
	"math/rand"
)

// Point represents a data point with an identifier and a vector of values (skills).
type Point struct {
	ID     string
	Vector []float64
}

// Cluster represents a cluster of points with a centroid.
type Cluster struct {
	Centroid []float64
	Points   []Point
	Cohesion float64
}

// KMeans performs the K-Means clustering algorithm.
// data: slice of points to cluster.
// k: number of clusters.
// maxIter: maximum number of iterations.
// rng: random number generator source (optional, uses global rand if nil).
func KMeans(data []Point, k int, maxIter int, rng *rand.Rand) ([]Cluster, error) {
	if k <= 0 {
		return nil, nil
	}
	if k > len(data) {
		k = len(data)
	}

	// Initialize centroids randomly
	centroids := make([][]float64, k)
	var perm []int
	if rng != nil {
		perm = rng.Perm(len(data))
	} else {
		perm = rand.Perm(len(data))
	}

	for i := 0; i < k; i++ {
		centroids[i] = make([]float64, len(data[perm[i]].Vector))
		copy(centroids[i], data[perm[i]].Vector)
	}

	var clusters []Cluster

	for iter := 0; iter < maxIter; iter++ {
		// Reset clusters
		clusters = make([]Cluster, k)
		for i := 0; i < k; i++ {
			clusters[i].Centroid = centroids[i]
		}

		// Assign points to nearest centroid
		for _, p := range data {
			nearestIndex := 0
			minDist := math.MaxFloat64

			for i, c := range centroids {
				dist := euclideanDistance(p.Vector, c)
				if dist < minDist {
					minDist = dist
					nearestIndex = i
				}
			}
			clusters[nearestIndex].Points = append(clusters[nearestIndex].Points, p)
		}

		// Update centroids
		converged := true
		for i := 0; i < k; i++ {
			if len(clusters[i].Points) == 0 {
				continue
			}

			newCentroid := make([]float64, len(centroids[i]))
			for _, p := range clusters[i].Points {
				for j, val := range p.Vector {
					newCentroid[j] += val
				}
			}
			for j := range newCentroid {
				newCentroid[j] /= float64(len(clusters[i].Points))
			}

			// Check for convergence
			if euclideanDistance(centroids[i], newCentroid) > 1e-6 {
				converged = false
			}
			centroids[i] = newCentroid
			clusters[i].Centroid = centroids[i]
		}

		if converged {
			break
		}
	}

	// Calculate cohesion for each cluster
	for i := range clusters {
		if len(clusters[i].Points) == 0 {
			clusters[i].Cohesion = 0
			continue
		}
		totalDist := 0.0
		for _, p := range clusters[i].Points {
			totalDist += euclideanDistance(p.Vector, clusters[i].Centroid)
		}
		clusters[i].Cohesion = totalDist / float64(len(clusters[i].Points))
	}

	return clusters, nil
}

func euclideanDistance(a, b []float64) float64 {
	sum := 0.0
	for i := range a {
		diff := a[i] - b[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}
