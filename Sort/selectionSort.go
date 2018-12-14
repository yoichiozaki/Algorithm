package Sort

// # 選択ソート
// A[0, n)で最大値を選択してA[n-1]と交換してn--、という動作をnが0になるまで繰り返す。
// 貪欲戦略の例の一つ。
func SelectionSort(a []int) []int {
	n := len(a)
	for i := n - 1; 1 <= i; i-- {
		maxPos := selectMax(&a, 0, i)
		if maxPos != i {
			a[i], a[maxPos] = a[maxPos], a[i]
		}
	}
	return a
}

func selectMax(a *[]int, left, right int) int {
	maxPos := left
	for i := left; i <= right; i++ {
		if (*a)[maxPos] < (*a)[i] {
			maxPos = i
		}
	}
	return maxPos
}
