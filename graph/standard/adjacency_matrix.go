package standard

type AdjacencyMatrix struct {
	VertexList []string
	Edges      [][]int
	NumOfEdges int
}

func NewAdjacencyMatrix(n int) *AdjacencyMatrix {

	edges := make([][]int, 0, n)
	for i := range edges {
		edges[i] = make([]int, 0, n)
	}

	return &AdjacencyMatrix{
		VertexList: make([]string, 0, n),
		Edges:      edges,
		NumOfEdges: 0,
	}
}

// InsertVertex ... 插入节点
func (a *AdjacencyMatrix) InsertVertex(vertex string) {
	a.VertexList = append(a.VertexList, vertex)
}

// InsertEdge ... 添加边
func (a *AdjacencyMatrix) InsertEdge(v1, v2, weight int) {
	a.Edges[v1][v2] = weight
	a.Edges[v2][v1] = weight
	a.NumOfEdges++
}

// GetNumOfVertex ... 返回节点个数
func (a *AdjacencyMatrix) GetNumOfVertex() int {
	return len(a.VertexList)
}

// GetNumOfEdges ... 返回边的数目
func (a *AdjacencyMatrix) GetNumOfEdges() int {
	return a.NumOfEdges
}

// GetVertexValueByIndex ... 返回VertexValue
// 0 -> "A"
// 1 -> "B"
// 2 -> "C"
func (a *AdjacencyMatrix) GetVertexValueByIndex(index int) string {
	return a.VertexList[index]
}

// GetWeight ... 返回权值
func (a *AdjacencyMatrix) GetWeight(v1, v2 int) int {
	return a.Edges[v1][v2]
}

// ShowGraph ... 显示图对应的矩阵
func (a *AdjacencyMatrix) ShowGraph() {

}
