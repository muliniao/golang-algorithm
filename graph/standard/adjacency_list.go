package standard

import (
	"container/list"
	"fmt"
)

const numOfVertex = 5

// [0]Bill ---> Rocket ---> Linda
// [1]Rocket ---> Bruce ---> Shirly ---> Bill
// [2]Linda ---> Bill ---> Shirly
// [3]Shirly ---> Bill
// [4]Bruce ---> Rocket

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

func NewAdjNode(vertex int, vertexInfo *VertexInfo) *AdjVNode {
	return &AdjVNode{
		vertex:     vertex,
		vertexINfo: vertexInfo,
	}
}

// VNode : 头结点
type VNode struct {
	//firstEdge *AdjVNode
	data interface{}
}

func NewVNode(data interface{}) *VNode {
	return &VNode{
		data: data,
	}
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

func (g *Graph) InsertVertex(vertex *VNode) error {
	if len(g.vNodeLists) >= numOfVertex {
		return fmt.Errorf("cannot exceed 5 vertex")
	}

	for _, v := range g.vNodeLists {
		if v.Len() == 0 {
			v.PushBack(vertex)
		}
	}

	g.numOfVertex++
	return nil
}

func (g *Graph) InsertEdge(edge *Edge) error {
	if edge.leftVertex < 0 || edge.leftVertex > 5 {
		return fmt.Errorf("")
	}

	if edge.rightVertex < 0 || edge.rightVertex > 5 {
		return fmt.Errorf("")
	}

	newAdjNode := NewAdjNode(edge.rightVertex, edge.vertexInfo)
	g.vNodeLists[edge.leftVertex].PushBack(newAdjNode)

	g.numOfEdge++
	return nil
}

func (g *Graph) GetNumOfVertex() int {
	return g.numOfVertex
}

func (g *Graph) GetNumOfEdge() int {
	return g.numOfEdge
}

func (g *Graph) GetWeight() {

}

func (g *Graph) ShowGraph() {

}
