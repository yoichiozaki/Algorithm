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
	time_in := time.Now()
	a = Sort.InsertionSort(a)
	time_out := time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", time_out.Sub(time_in).String())

	// selection sort
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("---------- selection sort ----------")
	// fmt.Println("before:\t", a)
	time_in = time.Now()
	a = Sort.SelectionSort(a)
	time_out = time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", time_out.Sub(time_in).String())

	// heap sort
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("------------ heap sort -------------")
	// fmt.Println("before:\t", a)
	time_in = time.Now()
	a = Sort.HeapSort(a)
	time_out = time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", time_out.Sub(time_in).String())

	// quick sort
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("------------ quick sort ------------")
	// fmt.Println("before:\t", a)
	time_in = time.Now()
	Sort.QuickSort(&a, 0, len(a)-1)
	time_out = time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", time_out.Sub(time_in).String())

	// bucket sort
	a = make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(100 * SIZE)
	}
	fmt.Println("----------- bucket sort ------------")
	// fmt.Println("before:\t", a)
	time_in = time.Now()
	a = Sort.BucketSort(a, 100 * SIZE)
	time_out = time.Now()
	// fmt.Println("after:\t", a)
	fmt.Println("time:\t", time_out.Sub(time_in).String())
}
