package Graph

import (
	"strings"
)

// # 深さ優先探索
// グラフの探索戦略のうちのもっとも基本的なもののうちの一つ。もう一つは幅優先探索。
// ある頂点からゴールとなる頂点までのたどり着くことができるかがわかる。ただ、その経路が最短経路であるという保証はない。
// 深さ優先探索は関数の再帰呼び出しもしくはスタックを用いて記述することができる。

var (
	maze    [][]string
	h, w    int
	visited [][]bool
)

func DepthFirstSearch(input string) bool {
	splittedRows := strings.Split(input, "\n")
	maze = make([][]string, len(splittedRows))
	for i := range maze {
		splitted := strings.Split(splittedRows[i], "")
		maze[i] = make([]string, len(splitted))
		for j := range splitted {
			maze[i][j] = splitted[j]
		}
	}
	h, w = len(maze), len(maze[0])
	visited = make([][]bool, h)
	for i := range visited {
		visited[i] = make([]bool, w)
	}
	sx, sy, gx, gy := 0, 0, 0, 0
	for i := range maze {
		for j := range maze[i] {
			switch maze[i][j] {
			case "s":
				sx, sy = j, i
			case "g":
				gx, gy = j, i
			}
		}
	}
	dfs(sx, sy)
	if visited[gy][gx] {
		return true
	} else {
		return false
	}
}

func dfs(x, y int) {
	if x < 0 || w <= x || y < 0 || h <= y || maze[y][x] == "#" { // 現在地が壁の中もしくは迷路の外である
		return
	}
	if visited[y][x] { // すでに訪問済みである
		return
	}
	visited[y][x] = true
	dfs(x+1, y)
	dfs(x-1, y)
	dfs(x, y+1)
	dfs(x, y-1)
}
