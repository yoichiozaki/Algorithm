package Graph

import (
	"strings"
)

// # 幅優先探索
// グラフの探索戦略のうちのもっとも基本的なもののうちの一つ。もう一つは深さ優先探索。
// ある頂点からゴールとなる頂点までのたどり着くことができるかがわかる。ただ、その経路が最短経路であるという保証はない。
// 幅優先探索はキューを用いて記述することができる。

func BreadthFirstSearch(input string) bool {
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
	dx := []int{0, -1, 0, 1}
	dy := []int{-1, 0, 1, 0}
	if sx == gx && sy == gy {
		return true
	}
	visited[sy][sx] = true
	que := make([]Point, 0)
	que = append(que, Point{sx, sy})
	for len(que) > 0 {
		current := que[0]
		que = que[1:]
		for i := 0; i < 4; i++ {
			next := Point{current.x + dx[i], current.y + dy[i]}
			if next.x == gx && next.y == gy {
				visited[next.y][next.x] = true
				break
			}
			if 0 <= next.x && next.x < w && 0 <= next.y && next.y < h &&
				maze[next.y][next.x] == "." && !visited[next.y][next.x] {
				visited[next.y][next.x] = true
				que = append(que, next)
			} else {
				continue
			}
		}
	}
	if visited[gy][gx] {
		return true
	} else {
		return false
	}
}

type Point struct {
	x, y int
}
