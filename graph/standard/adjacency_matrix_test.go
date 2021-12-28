package standard

import "testing"

var adjacencyMatrix *AdjacencyMatrix

func init() {
	adjacencyMatrix = NewAdjacencyMatrix(6)
}

func TestAdjacencyMatrix_InsertVertex(t *testing.T) {
	adjacencyMatrix.InsertVertex("A")
	adjacencyMatrix.InsertVertex("B")
	adjacencyMatrix.InsertVertex("C")
	adjacencyMatrix.InsertVertex("D")
	adjacencyMatrix.InsertVertex("E")
}

func TestAdjacencyMatrix_InsertEdge(t *testing.T) {

	adjacencyMatrix.InsertEdge(0, 0, 0)
	adjacencyMatrix.InsertEdge(0, 1, 1)
	adjacencyMatrix.InsertEdge(0, 2, 1)
	adjacencyMatrix.InsertEdge(0, 3, 1)
	adjacencyMatrix.InsertEdge(0, 4, 1)
	adjacencyMatrix.InsertEdge(0, 5, 0)

	adjacencyMatrix.InsertEdge(1, 0, 1)

}
