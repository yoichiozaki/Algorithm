package main

import (
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
}
