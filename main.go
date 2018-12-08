package main

import (
	"Algorithm/Sort"
	"fmt"
	"math/rand"
)

const (
	SIZE = 100
)
func main() {
	a := make([]int, SIZE)
	for i := range a {
		a[i] = rand.Intn(1000*SIZE)
	}
	Sort.InsertionSort(a)
	fmt.Println(a)
}
