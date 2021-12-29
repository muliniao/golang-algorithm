package standard

import (
	"testing"
)

var adjacencyMatrix *AdjacencyMatrix

func init() {
	adjacencyMatrix = NewAdjacencyMatrix(5)
}

func TestAdjacencyMatrix_InsertVertex(t *testing.T) {
	adjacencyMatrix.InsertVertex("A")
	adjacencyMatrix.InsertVertex("B")
	adjacencyMatrix.InsertVertex("C")
	adjacencyMatrix.InsertVertex("D")
	adjacencyMatrix.InsertVertex("E")
}

func TestAdjacencyMatrix_InsertEdge(t *testing.T) {

	// A-B, A-C, B-C, B-D, B-E
	adjacencyMatrix.InsertEdge(0, 1, 1)
	adjacencyMatrix.InsertEdge(0, 2, 1)
	adjacencyMatrix.InsertEdge(1, 2, 1)
	adjacencyMatrix.InsertEdge(1, 3, 1)
	adjacencyMatrix.InsertEdge(1, 4, 1)

	adjacencyMatrix.ShowGraph()
}
