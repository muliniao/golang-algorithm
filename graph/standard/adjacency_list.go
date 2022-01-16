package standard

import "container/list"

const m = 5

// Edge : 边定义
type Edge struct {
	leftVertex  int
	rightVertex int
	vertexInfo  *VertexInfo
}

// AdjVNode : 边表结点
type AdjVNode struct {
	vertex       int         // 顶点下标
	vertexINfo   *VertexInfo // 结点信息
	nextAdjVNode *AdjVNode   // 下一个边表结点
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
