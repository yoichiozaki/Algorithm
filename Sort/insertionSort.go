package Sort

// # 挿入ソート
// 「トランプを並べ替える」動作に似ている。
// 挿入ソートは要素数が少ないか、最初から集まりがほぼ整列している時に使う。
// 配列が整列済みの時に最適性能を示し、逆順に整列されている時に最悪性能を示す。
// 挿入ソートは余分なスペースをほとんど必要としない。
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
