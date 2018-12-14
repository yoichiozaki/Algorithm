package Sort

// # ヒープソート
// 選択ソートでは「A[0, n)で最大値を選択」するために少なくともn-1回比較が必要であったが、
// 直接比較する要素の個数を最小化することはできないだろうかという考えのものとで出てくるのがヒープソート
// 一般に「親の要素より子の要素が小さい」ように構築される木構造をヒープと呼ぶが、
// このヒープという木構造を用いることで「A[0, n)で最大値を選択」するのに比較をする必要がなくなる（何故ならば根が最大値であるから）
// という性質を利用した整列法がヒープソート
// ヒープは抽象データ構造であって、その実現方法は様々に考えられることに注意。
// 「木構造だから木構造をそのまま実装しよう」というやり方でもいいが、ヒープはその形状の特性から配列でも実現することが可能である。

func HeapSort(a []int) []int {
	buildHeap(&a) // 初期化: sortしたい対象の配列をヒープに変換する
	for i := len(a) - 1; 1 < i; i-- {
		a[0], a[i] = a[i], a[0] // ヒープの根は[0, i]で最大値になっているので移動して取っておく
		heapify(&a, 0, i) // もう一度ヒープ化する
	}
	return a
}

func buildHeap(a *[]int) {
	n := len(*a)/2-1
	for i := n; 0 < i; i-- {
		heapify(a, i, n)
	}
}

// a[idx, max)をヒープ化する
func heapify(a *[]int, idx, max int) {
	largest := idx // 親a[idx]が自身の子供のどちらよりも大きいことを仮定する
	left := 2*idx + 1;
	right := 2*idx + 2
	if left < max && (*a)[idx] < (*a)[left] {
		largest = left // 左の子が親よりも大きかった
	}
	if right < max && (*a)[largest] < (*a)[right] {
		largest = right // 右の子が親と左の兄弟のどちらよりも大きかった
	}
	if largest != idx {
		(*a)[idx], (*a)[largest] = (*a)[largest], (*a)[idx]
		heapify(a, largest, max) // 再帰的にheapifyを呼び出すことで大きい値を一つ上に持ち上げていく
	}
}
