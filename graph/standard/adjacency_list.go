package standard

import (
	"container/list"
	"fmt"
)

const numOfVertex = 5

// Bill ---> Rocket ---> Linda
// Rocket ---> Bruce ---> Shirly ---> Bill
// Linda ---> Bill ---> Shirly
// Shirly ---> Bill
// Bruce ---> Rocket

// Edge : 边定义
type Edge struct {
	leftVertex  int
	rightVertex int
	vertexInfo  *VertexInfo
}

// AdjVNode : 边表结点
type AdjVNode struct {
	vertex     int         // 顶点下标
	vertexINfo *VertexInfo // 结点信息
	//nextAdjVNode *AdjVNode   // 下一个边表结点
}

// VNode : 头结点
type VNode struct {
	firstEdge *AdjVNode
	data      interface{}
}

// VertexInfo : 结点信息
type VertexInfo struct {
	weight int
}

// Graph : 图
type Graph struct {
	numOfVertex int
	numOfEdge   int
	vNodeLists  []*list.List
}

func NewGraph() *Graph {
	graph := new(Graph)
	graph.numOfEdge = 0
	graph.numOfVertex = numOfVertex

	for i := 0; i < numOfVertex-1; i++ {
		graph.vNodeLists[i] = list.New()
	}

	return graph
}

func (g *Graph) InsertVertex(vertex VNode) error {
	if len(g.vNodeLists) >= numOfVertex {
		return fmt.Errorf(" ")
	}

	for _, v := range g.vNodeLists {
		if v.Len() == 0 {
			v.PushBack(vertex)
		}
	}

	return nil
}

func (g *Graph) InsertEdge(edge *Edge) error {

	return nil
}

func (g *Graph) GetNumOfVertex() {

}

func (g *Graph) GetNumOfEdge() {

}

func (g *Graph) GetWeight() {

}

func (g *Graph) ShowGraph() {

}
