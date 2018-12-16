package Sort

// # 数え上げソート
// ソート対象のデータをキーとして、キーの出現回数とその累積度数分布を計算して利用することで整列を行うアルゴリズム。
// データがとりうる値の範囲をあらかじめ知っている必要がある。バケットソートと違い、キーの重複を許す。
// 作業領域はソート対象のデータの範囲分だけ必要なので、あまりに大きな範囲のデータを取り扱うには物理的に不可能になる。
// しかし大きな作業領域と引き換えにデータの大小を比較しなくてよいため高速な整列が可能。

func CountingSort(a []int, max int) []int {
	ret := make([]int, len(a))
	counting := make([]int, max+1)
	for i := range a {
		counting[a[i]]++
	}
	for i := 0; i < max; i++ {
		counting[i+1] += counting[i] // 累積
	}
	for i := len(a) - 1; 0 <= i; i-- {
		counting[a[i]]--
		ret[counting[a[i]]] = a[i]
	}
	return ret
}
