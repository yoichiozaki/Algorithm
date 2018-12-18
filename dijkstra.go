package main

import (
	"errors"
	"fmt"
)

// # ダイクストラ法
// 単一始点最短経路問題の解を与えるアルゴリズムの一つ。単一始点最短経路問題とはグラフにおいてスタートとゴールがそれぞれ一つに決められている
// 時に、どの経路を辿っていくと最小のコストでゴールにたどり着けるかという問題のこと。
// ダイクストラ法は優先度付きキューによる実装と、密グラフ向けの二次元隣接行列での実装が考えられる。

func main() {
	g := NewDirectedGraph()
	g.Add("s", "a", 2)
	g.Add("s", "b", 5)
	g.Add("a", "b", 2)
	g.Add("a", "c", 5)
	g.Add("b", "c", 4)
	g.Add("b", "d", 2)
	g.Add("c", "z", 7)
	g.Add("d", "c", 5)
	g.Add("d", "z", 2)
	path, err := g.Dijgstra("s", "z")
	if err != nil {
		fmt.Println("Goal not found")
		return
	}
	for _, node := range path {
		fmt.Printf("node: %v, cost: %v\n", node.name, node.costFromStart)
	}
}
type Node struct {
	name          string
	edges         []*Edge
	isDone        bool
	costFromStart int
	prev          *Node
}

func NewNode(name string) *Node {
	return &Node{name: name, edges: []*Edge{}, isDone: false, costFromStart:  -1, prev: nil}
}
func (n *Node) AddEdge(e *Edge) {
	n.edges = append(n.edges, e)
}
type Edge struct {
	next *Node
	cost int
}
func NewEdge(next *Node, cost int) *Edge {
	return  &Edge{next: next, cost: cost}
}
type DirectedGraph struct {
	nodes map[string]*Node
}
func NewDirectedGraph() *DirectedGraph {
	return  &DirectedGraph{map[string]*Node{}}
}
func (dg *DirectedGraph) Add(src, dst string, cost int) {
	srcNode, exist := dg.nodes[src]
	if !exist {
		srcNode = NewNode(src)
		dg.nodes[src] = srcNode
	}
	dstNode, exist := dg.nodes[dst]
	if !exist {
		dstNode = NewNode(dst)
		dg.nodes[dst] = dstNode
	}
	edge := NewEdge(dstNode, cost)
	srcNode.AddEdge(edge)
}
func (dg *DirectedGraph) Dijgstra(start, goal string) (ret []*Node, err error) {
	startNode := dg.nodes[start]
	startNode.costFromStart = 0
	for {
		node, err := dg.nextNode()
		if err != nil {
			return nil, errors.New("Goal not found")
		}
		if node.name == goal {
			break
		}
		for _, edge := range node.edges {
			nextNode := edge.next
			if nextNode.isDone {
				continue
			}
			cost := node.costFromStart + edge.cost
			if  nextNode.costFromStart == -1 || cost < nextNode.costFromStart {
				nextNode.costFromStart = cost
				nextNode.prev = node
			}
		}
		node.isDone = true
	}
	n := dg.nodes[goal]
	for {
		ret = append(ret, n)
		if n.name == start {
			break
		}
		n = n.prev
	}
	return ret, nil
}

func (dg *DirectedGraph) nextNode() (next *Node, err error) {
	for _, node := range dg.nodes {
		if node.isDone {
			continue
		}
		if node.costFromStart == -1 {
			continue
		}
		if next == nil {
			next = node
		}
		if node.costFromStart < next.costFromStart {
			next = node
		}
	}
	if next == nil {
		return nil, errors.New("Untreated node not found")
	}
	return
}