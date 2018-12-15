package main

import (
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
	fmt.Println("---------- insertion sort ----------")
	// fmt.Println("before:\t", a)
	timeIn := time.Now()
	a = Sort.InsertionSort(a)
	timeOut := time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", timeOut.Sub(timeIn).String())

	// selection sort
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("---------- selection sort ----------")
	// fmt.Println("before:\t", a)
	timeIn = time.Now()
	a = Sort.SelectionSort(a)
	timeOut = time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", timeOut.Sub(timeIn).String())

	// heap sort
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("------------ heap sort -------------")
	// fmt.Println("before:\t", a)
	timeIn = time.Now()
	a = Sort.HeapSort(a)
	timeOut = time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", timeOut.Sub(timeIn).String())

	// quick sort
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("------------ quick sort ------------")
	// fmt.Println("before:\t", a)
	timeIn = time.Now()
	Sort.QuickSort(&a, 0, len(a)-1)
	timeOut = time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", timeOut.Sub(timeIn).String())

	// bucket sort
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("----------- bucket sort ------------")
	// fmt.Println("before:\t", a)
	timeIn = time.Now()
	a = Sort.BucketSort(a, 100 * SIZE)
	timeOut = time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", timeOut.Sub(timeIn).String())

	// counting sort
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("----------- count sort -----------")
	// fmt.Println("before:\t", a)
	timeIn = time.Now()
	a = Sort.CountingSort(a, 100 * SIZE)
	timeOut = time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", timeOut.Sub(timeIn).String())

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
}
