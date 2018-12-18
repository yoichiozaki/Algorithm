package main

import (
	"Algorithm/Graph"
	"Algorithm/Search"
	"Algorithm/Sort"
	"fmt"
	"math/rand"
	"time"
)

const (
	SIZE = 10000
)

func main() {
	// insertion sort
	rand.Seed(time.Now().UnixNano())
	a := make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("--------- insertion sort ---------")
	// fmt.Println("before:\t", a)
	timeIn := time.Now()
	a = Sort.InsertionSort(a)
	timeOut := time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", timeOut.Sub(timeIn).String())
	fmt.Println()

	// selection sort
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("--------- selection sort ---------")
	// fmt.Println("before:\t", a)
	timeIn = time.Now()
	a = Sort.SelectionSort(a)
	timeOut = time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", timeOut.Sub(timeIn).String())
	fmt.Println()

	// heap sort
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("----------- heap sort ------------")
	// fmt.Println("before:\t", a)
	timeIn = time.Now()
	a = Sort.HeapSort(a)
	timeOut = time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", timeOut.Sub(timeIn).String())
	fmt.Println()

	// quick sort
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("----------- quick sort -----------")
	// fmt.Println("before:\t", a)
	timeIn = time.Now()
	Sort.QuickSort(&a, 0, len(a)-1)
	timeOut = time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", timeOut.Sub(timeIn).String())
	fmt.Println()

	// bucket sort
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("---------- bucket sort -----------")
	// fmt.Println("before:\t", a)
	timeIn = time.Now()
	a = Sort.BucketSort(a, 100*SIZE)
	timeOut = time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", timeOut.Sub(timeIn).String())
	fmt.Println()

	// counting sort
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("----------- count sort -----------")
	// fmt.Println("before:\t", a)
	timeIn = time.Now()
	a = Sort.CountingSort(a, 100*SIZE)
	timeOut = time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", timeOut.Sub(timeIn).String())
	fmt.Println()

	// merge sort
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("----------- merge sort -----------")
	// fmt.Println("before:\t", a)
	timeIn = time.Now()
	a = Sort.MergeSort(a)
	timeOut = time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", timeOut.Sub(timeIn).String())
	fmt.Println()

	// (direct) sequential search
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	var target int
	fmt.Println("--- (direct) sequential search ---")
	timeIn = time.Now()
	for i := 0; i < 1000; i++ {
		target = rand.Intn(100 * SIZE)
		if Search.SequentialSearch(a, target) {
			fmt.Printf("Found! - Target: %d\n", target)
		}
	}
	timeOut = time.Now()
	fmt.Println("time for 1000 searches:\t", timeOut.Sub(timeIn).String())
	fmt.Println()

	// sequential search with iterator
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("--- sequential search with itr ---")
	timeIn = time.Now()
	for i := 0; i < 1000; i++ {
		target = rand.Intn(100 * SIZE)
		if Search.SequentialSearchWithIterator(a, target) {
			fmt.Printf("Found! - Target: %d\n", target)
		}
	}
	timeOut = time.Now()
	fmt.Println("time for 1000 searches:\t", timeOut.Sub(timeIn).String())
	fmt.Println()

	// bianry search
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("--------- binary search ---------")
	Sort.QuickSort(&a, 0, len(a)-1) // this is my own implementation.
	// sort.Ints(a)
	timeIn = time.Now()
	for i := 0; i < 1000; i++ {
		target = rand.Intn(100 * SIZE)
		if Search.BinarySearch(a, target, false) {
			fmt.Printf("Found! - Target: %d\n", target)
		}
	}
	timeOut = time.Now()
	fmt.Println("time for 1000 searches:\t", timeOut.Sub(timeIn).String())
	fmt.Println()

	// hash based search
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("-------- hash based search --------")
	table := Search.NewHashTable()
	table.Load(a, Search.LinkedList)
	timeIn = time.Now()
	for i := 0; i < 1000; i++ {
		target = rand.Intn(100 * SIZE)
		if Search.HashBasedSearch(a, table, target, Search.LinkedList) {
			fmt.Printf("Found! - Target: %d\n", target)
		}
		// TODO: open address method is not yet completed.
		// if Search.HashBasedSearch(a, target, Search.OpenAddress) {
		// 	fmt.Printf("Found! - Target: %d\n", target)
		// }
	}
	timeOut = time.Now()
	fmt.Println("time for 1000 searches:\t", timeOut.Sub(timeIn).String())
	fmt.Println()

	// bloom filter
	// bf := Search.NewBloomFilter(3)
	// bf.Add("1")
	// bf.Add("2")
	// bf.Add("3")
	// bf.Add("4")
	// bf.Add("5")
	// fmt.Println(bf.Exists("1"))
	// fmt.Println(bf.Exists("2"))
	// fmt.Println(bf.Exists("3"))
	// fmt.Println(bf.Exists("4"))
	// fmt.Println(bf.Exists("5"))
	// fmt.Println(bf.Exists("10"))
	// fmt.Println(bf.Exists("999"))

	// search with binary search tree
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("- search with binary search tree -")
	bst := Search.NewTree()
	for _, i := range a {
		bst.Add(i)
	}
	// bst.Inorder()
	timeIn = time.Now()
	for i := 0; i < 1000; i++ {
		target = rand.Intn(100 * SIZE)
		if bst.Search(target) {
			fmt.Printf("Found! - Target: %d\n", target)
		}
	}
	timeOut = time.Now()
	fmt.Println("time for 1000 searches:\t", timeOut.Sub(timeIn).String())
	fmt.Println()

	maze1 := `s.........
#########.
#.......#.
#..####.#.
##....#.#.
#####.#.#.
g.#.#.#.#.
#.#.#.#.#.
#.#.#.#.#.
#.....#...`
	maze2 := `s.........
#########.
#.......#.
#..####.#.
##....#.#.
#####.#.#.
g.#.#.#.#.
#.#.#.#.#.
###.#.#.#.
#.....#...`

	// depth first search
	fmt.Println(maze1)
	if Graph.DepthFirstSearch(maze1) {
		fmt.Println("we can reach the goal from start.")
	} else {
		fmt.Println("we can NOT reach the goal from start.")
	}

	fmt.Println(maze2)
	if Graph.DepthFirstSearch(maze2) {
		fmt.Println("we can reach the goal from start.")
	} else {
		fmt.Println("we can NOT reach the goal from start.")
	}

	// breadth first search
	fmt.Println(maze1)
	if Graph.BreadthFirstSearch(maze1) {
		fmt.Println("we can reach the goal from start.")
	} else {
		fmt.Println("we can NOT reach the goal from start.")
	}

	fmt.Println(maze2)
	if Graph.BreadthFirstSearch(maze2) {
		fmt.Println("we can reach the goal from start.")
	} else {
		fmt.Println("we can NOT reach the goal from start.")
	}

	// graph
	// nodes := make([]Graph.Vertex, 100)
	// for i := range nodes {
	// 	nodes[i] = Graph.Vertex(i)
	// }
	// sample := Graph.NewGraph()
	// for _, node := range nodes {
	// 	sample.AddVertex(node)
	// }
	// for i := 0; i < 100; i++ {
	// 	if err := sample.AddEdge(Graph.Vertex(rand.Intn(100)), Graph.Vertex(rand.Intn(100)), rand.Intn(10));
	// 	err != nil {
	// 		continue
	// 	}
	// }
	// err := sample.Visualize()
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
