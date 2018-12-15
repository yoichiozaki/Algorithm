package Sort

// # マージソート
// 記憶領域をたくさん使う代わりに時間的には高速な整列が可能。空間計算量と時間計算量のトレードオフが垣間見える。

func MergeSort(a []int) []int {
	n := len(a)
	copied := make([]int, n)
	copy(copied, a)
	mergeSort(copied, a, 0, n)
	return a
}

func mergeSort(a, result []int, start, end int) {
	if end - start < 2 {
		return
	}
	if end - start == 2 {
		if result[start] > result[start+1] {
			result[start], result[start+1] = result[start+1], result[start]
		}
		return
	}
	mid := (end + start)/2
	mergeSort(result, a, start, mid)
	mergeSort(result, a, mid, end)
	i := start
	j := mid
	idx := start
	for idx < end {
		if end <= j || (i < mid && a[i] < a[j]) {
			result[idx] = a[i]
			i++
		} else {
			result[idx] = a[j]
			j++
		}
		idx++
	}
}
