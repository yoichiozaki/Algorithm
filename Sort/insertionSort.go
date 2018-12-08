package Sort

func InsertionSort(a []int) []int {
	for pos := 1; pos < len(a); pos++ {
		insert(&a, pos, a[pos])
	}
	return a
}

func insert(a *[]int, pos, value int) {
	i := pos - 1
	for 0 <= i && value < (*a)[i] {
		(*a)[i+1] = (*a)[i]
		i--
	}
	(*a)[i+1] = value
}
