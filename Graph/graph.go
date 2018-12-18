package Graph

import (
	"errors"
	"github.com/awalterschulze/gographviz"
	"os"
	"os/exec"
	"strconv"
)

// # グラフ
// グラフとは複雑な構造を持った情報を表現するためのデータ構造である。グラフは節点と辺の集合である。グラフは辺に向きがあるかないか、
// 辺に重みがあるかないかで分類することができる。
// グラフの実装は様々ある。グラフの形をそのまま書き下したような形での実装もできるし、隣接頂点のリストで保持する、もしくは隣接行列の形で
// 頂点間の接続を保持することでグラフを表現することもできる。密なグラフは隣接行列で表現するのが良い。疎なグラフは隣接リストで表現すると良い。

// 頂点
type Vertex uint

// 辺
type Edge struct {
	From Vertex
	To   Vertex
}

type Graph struct {
	vertices      map[Vertex]int      // 頂点集合 = 頂点: 何番目に追加されたか
	verticesCount int                 // 頂点数
	edges         map[Edge]int        // 辺集合 = 始点と終点のペア: 重み
	edgesCount    int                 // 辺数
	isDirected    bool                // 有向か無向か
	neighbours    map[Vertex][]Vertex // 各頂点に隣接する頂点の集合 = 頂点: 隣接頂点集合
}

func NewGraph() *Graph {
	return &Graph{
		vertices:      map[Vertex]int{},
		verticesCount: 0,
		edges:         map[Edge]int{},
		edgesCount:    0,
		isDirected:    false,
		neighbours:    map[Vertex][]Vertex{},
	}
}

func (g *Graph) SetDir(isDir bool) {
	g.isDirected = isDir
}

func (g *Graph) GetVertices() []Vertex {
	vertices := make([]Vertex, g.verticesCount)
	for v, i := range g.vertices {
		vertices[i] = v
	}
	return vertices
}

func (g *Graph) ExistsVertex(v Vertex) bool {
	if _, ok := g.vertices[v]; !ok {
		return false
	}
	return true
}

func (g *Graph) AddVertex(v Vertex) error {
	if g.ExistsVertex(v) {
		return errors.New("g already has the vertex v")
	}
	g.vertices[v] = g.verticesCount
	g.verticesCount++
	return nil
}

func (g *Graph) RemoveVertex(v Vertex) error {
	if !g.ExistsVertex(v) {
		return errors.New("input v does not exist in g")
	}
	delete(g.vertices, v)
	g.verticesCount--
	return nil
}

func (g *Graph) GetEdges() []Edge {
	edges := make([]Edge, g.edgesCount)
	i := 0
	for e, _ := range g.edges {
		edges[i] = e
		i++
	}
	return edges
}

func (g *Graph) ExistsEdge(from, to Vertex) bool {
	e := Edge{From: from, To: to}
	if _, ok := g.edges[e]; ok {
		return true
	}
	if !g.isDirected {
		e = Edge{From: to, To: from}
		if _, ok := g.edges[e]; ok {
			return true
		}
	}
	return false
}

func (g *Graph) AddEdge(from, to Vertex, weight int) error {
	if from == to {
		return errors.New("can not add an edge from and to the same vertex")
	}
	if !g.ExistsVertex(from) || !g.ExistsVertex(to) {
		return errors.New("such edge does not exist")
	}
	if g.ExistsEdge(from, to) {
		return errors.New("such edge already exists in g")
	}
	g.edges[Edge{From: from, To: to}] = weight
	g.neighbours[from] = append(g.neighbours[from], to)
	g.edgesCount++
	if !g.isDirected {
		g.neighbours[to] = append(g.neighbours[to], from)
	}
	return nil
}

func (g *Graph) RemoveEdge(from, to Vertex) error {
	if from == to {
		return errors.New("can not remove an edge from and to the same vertex")
	}
	if !g.ExistsEdge(from, to) {
		return errors.New("such edge does not exit in g")
	}
	delete(g.edges, Edge{From: from, To: to})
	g.edgesCount--
	return nil
}

func (g *Graph) GetWeight(from, to Vertex) (int, error) {
	if from == to {
		return 0, errors.New("can not get an edge from and to the same vertex")
	}
	if !g.ExistsEdge(from, to) {
		return 0, errors.New("such edge does not exist")
	}
	weight := g.edges[Edge{From: from, To: to}]
	return weight, nil
}

func (g *Graph) SetWeight(from, to Vertex, weight int) error {
	if from == to {
		return errors.New("can not change an edge from and to the same vertex")
	}
	if !g.ExistsEdge(from, to) {
		return errors.New("such edge does not exist")
	}
	g.edges[Edge{From: from, To: to}] = weight
	return nil
}

func (g *Graph) GetNeighbours(v Vertex) []Vertex {
	vertices := make([]Vertex, g.vertices[v])
	return vertices
}

func (g *Graph) Visualize() error {
	gv := gographviz.NewGraph()
	if err := gv.SetName("G"); err != nil {
		return err
	}
	if err := gv.SetDir(g.isDirected); err != nil {
		return err
	}
	if err := gv.AddAttr("G", "bgcolor", "\"#343434\""); err != nil {
		return err
	}
	if err := gv.AddAttr("G", "layout", "circo"); err != nil {
		return err
	}
	nodeAttrs := make(map[string]string)
	nodeAttrs["colorscheme"] = "rdylgn11"
	nodeAttrs["style"] = "\"solid,filled\""
	nodeAttrs["fontcolor"] = "6"
	nodeAttrs["fontname"] = "\"Migu 1M\""
	nodeAttrs["color"] = "7"
	nodeAttrs["fillcolor"] = "11"
	nodeAttrs["shape"] = "doublecircle"
	for _, v := range g.GetVertices() {
		if err := gv.AddNode("G", strconv.FormatUint(uint64(v), 10), nodeAttrs); err != nil {
			return err
		}
	}
	edgeAttrs := make(map[string]string)
	edgeAttrs["color"] = "white"
	for _, e := range g.GetEdges() {
		if err := gv.AddEdge(
			strconv.FormatUint(uint64(e.From), 10),
			strconv.FormatUint(uint64(e.To), 10),
			g.isDirected, edgeAttrs); err != nil {
			return err
		}
	}
	s := gv.String()
	dotfile := "./Graph/img/gv.dot"
	pngfile := "./Graph/img/gv.png"
	file, err := os.Create(dotfile)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write([]byte(s))
	if err = exec.Command("dot", "-T", "png",
		dotfile, "-o", pngfile).Run(); err != nil {
		return err
	}
	return nil
}
