package firstsearch

import (
	"fmt"

	"golang-algorithm/graph/standard"
	"golang-algorithm/queue/standard/arrayqueue"
)

// 1.访问初始结点v,并标记结点v为已访问
// 2.结点v入队列
// 3.当队列为空时,继续执行,否则算法结束
// 4.出队列,取队列头结点u
// 5.查找结点u第一个临接结点u
// 6.若u的邻接节点w不存在,则跳转到步骤3; 否则循环以下三个步骤
// -6.1 结点w未访问,将w标记为已访问
// -6.2 结点w入列
// -6.3 查找结点u的继w邻接节点后的下一个邻接节点w,跳转到步骤6

type BroadFirstSearch struct {
	adjacencyMatrix *standard.AdjacencyMatrix
}

func (b *BroadFirstSearch) bfs(isVisited []bool, i int) {

	//var u int
	//var w int
	queue := arrayqueue.NewArrayQueue()

	fmt.Println(b.adjacencyMatrix.GetVertexValueByIndex(i) + "=>")
	isVisited[i] = true
	queue.Add(i)

	for len(queue.Elements) != 0 {

	}

}
