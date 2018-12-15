package Sort

import (
	"math/rand"
)

// # クイックソート
// 値を個別に比較して整列するというアプローチではなくて、「この値より小さいものは左に、大きいものは右に」
// という分割操作を再帰的に実行することでこれ以上分割できない要素数まで分割が進めば、結果的に全体が整列している
// のがクイックソートである。
// 全体を整列するという問題を、「右側を整列する」「左側を整列する」という独立な部分問題に分けてそれぞれを解決する
// アプローチで、このように複雑な問題を独立な部分問題に分けてそれぞれを解決することで全体を解決するアプローチを
// 分割統治法という。
// ここで注目すべきは「右側と左側の境目の選び方」である。賢明に境界を選べば右側と左側の要素数が大体同じくらいになる。
// つまり境界として配列内の要素の中央値を採用するということである。何故ならばこうすることで解決すべき部分問題の大きさが
// 元の問題の大凡半分になるからである。「部分問題のサイズを小さくしたいならもっと小さくできるではないか」という反論は
// 一方を半分より小さくするともう一方が半分より大きくなることに気がつくべきだ。
// しかし上でば「境界として配列内の要素の中央値を採用する」とさらっと書いたが、これはいささかチャレンジングである。
// 何故ならば「未整列」の配列からその中央値を弾き出すことが必要になるからである。
// 何も考えないで「中央値など事前にわからんのだからランダムに選んでやる」というアプローチを取ると素朴なクイックソートとなる。
// この時「境界として配列内の要素の中央値を採用」できるかは運任せになる。が、それでもそれなりに速い整列アルゴリズムである。
// クイックソートの最適化としては、分割していって整列しなければならない部分配列のサイズが閾値を下回ったらそこからは挿入ソートに
// 切り替えるというものがある。
// ピボット選択にランダムな値を用いたクイックソートでさえ、他の整列アルゴリズムよりも優れた平均時性能を発揮する。
// クイックソートはこれまで数多くの最適化の研究がなされてきた。理想的な場合ではpartition関数がその時の整列対象を
// 大凡半分に分割してO(n log n)である。最悪の場合は毎回のピボットとして最大値もしくは最小値をとった場合でO(n^2)である。
// ほとんどのシステムで整列に用いるクイックソートが用意されている。OSでは最適化の施されたクイックソートを用いることが多い。
// Linuxのあるバージョンではヒープソートを用いてqsortを実装しているものもあるらしい。
// 最適化には次のようなものがある。
// - 再帰処理を省くために部分タスクを蓄えておくスタックを作る
// - median-of-three戦略に基づいて出来るだけ最適なピボットを選択する
// - 最小分割サイズを設定して、整列対象の配列がそれ以下のサイズならば以降は挿入ソートに切り替える（上述）
// - 部分問題の処理において、サイズの小さい方の部分問題を先に処理して再帰スタックの全体サイズを小さくする
// これらの最適化は最適時性能を発揮する可能性を上げるための最適化であり、クイックソートの最悪時性能を軽減するものではないことに
// 注意が必要である。最悪時であってもO(n log n)の性能を発揮することを保証するためには、毎回のピボットの選択において
// 整列したい配列の中央値（の近似値）を選択することを保証することが必要である。このための理論もあるらしい。-> BFPRT分割アルゴリズム
// クイックソートの発展形としてイントロソートと言うものが提案されている。これはクイックソートの再帰の深さを監視し、それがlog(n)を超えた
// 段階でヒープソートに切り替えるというもので、C++の標準ライブラリのSGI実装ではイントロソートを元にした整列メカニズムを採用しているという。

func QuickSort(a *[]int, left, right int) {
	if left < right {
		pivotIndex := randomSelectPivotIndex(left, right)
		pivotIndex = partition(a, pivotIndex, left, right)
		QuickSort(a, left, pivotIndex-1)  // 左側を整列させる
		QuickSort(a, pivotIndex+1, right) // 右側を整列させる
	}
}

// 配列aを初期状態でのa[pivotIndex]の値を境界値, storeを分割後の配列a内でのpivotの位置とすると
// a[left, store)の全ての要素がpivot以下でありかつ、a[store, right]の全てに要素がpivotより大きい
// ように分割し、storeを返す関数
func partition(a *[]int, pivotIndex, left, right int) int {
	pivot := (*a)[pivotIndex]
	(*a)[pivotIndex], (*a)[right] = (*a)[right], (*a)[pivotIndex]
	store := left
	for idx := left; idx < right; idx++ {
		if (*a)[idx] < pivot {
			(*a)[idx], (*a)[store] = (*a)[store], (*a)[idx]
			store++
		}
	}
	(*a)[store], (*a)[right] = (*a)[right], (*a)[store]
	return store
}

// 何も考えずにランダムにピボット位置を与える
func randomSelectPivotIndex(left, right int) int {
	return rand.Intn(right-left) + left
}

// golangの標準ライブラリでのソートの実装は以下の通り。
// func quickSort(data Interface, a, b, maxDepth int) {
// 	for b-a > 12 { // Use ShellSort for slices <= 12 elements
// 		if maxDepth == 0 {
// 			heapSort(data, a, b)
// 			return
// 		}
// 		maxDepth--
// 		mlo, mhi := doPivot(data, a, b)
// 		// Avoiding recursion on the larger subproblem guarantees
// 		// a stack depth of at most lg(b-a).
// 		if mlo-a < b-mhi {
// 			quickSort(data, a, mlo, maxDepth)
// 			a = mhi // i.e., quickSort(data, mhi, b)
// 		} else {
// 			quickSort(data, mhi, b, maxDepth)
// 			b = mlo // i.e., quickSort(data, a, mlo)
// 		}
// 	}
// 	if b-a > 1 {
// 		// Do ShellSort pass with gap 6
// 		// It could be written in this simplified form cause b-a <= 12
// 		for i := a + 6; i < b; i++ {
// 			if data.Less(i, i-6) {
// 				data.Swap(i, i-6)
// 			}
// 		}
// 		insertionSort(data, a, b)
// 	}
// }
//
// ピボット選択にはmedian-of-three-medians-of-three戦略をとっている
// func doPivot(data Interface, lo, hi int) (midlo, midhi int) {
// 	m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
// 	if hi-lo > 40 {
// 		// Tukey's ``Ninther,'' median of three medians of three.
// 		s := (hi - lo) / 8
// 		medianOfThree(data, lo, lo+s, lo+2*s)
// 		medianOfThree(data, m, m-s, m+s)
// 		medianOfThree(data, hi-1, hi-1-s, hi-1-2*s)
// 	}
// 	medianOfThree(data, lo, m, hi-1)
//
// 	// Invariants are:
// 	//	data[lo] = pivot (set up by ChoosePivot)
// 	//	data[lo < i < a] < pivot
// 	//	data[a <= i < b] <= pivot
// 	//	data[b <= i < c] unexamined
// 	//	data[c <= i < hi-1] > pivot
// 	//	data[hi-1] >= pivot
// 	pivot := lo
// 	a, c := lo+1, hi-1
//
// 	for ; a < c && data.Less(a, pivot); a++ {
// 	}
// 	b := a
// 	for {
// 		for ; b < c && !data.Less(pivot, b); b++ { // data[b] <= pivot
// 		}
// 		for ; b < c && data.Less(pivot, c-1); c-- { // data[c-1] > pivot
// 		}
// 		if b >= c {
// 			break
// 		}
// 		// data[b] > pivot; data[c-1] <= pivot
// 		data.Swap(b, c-1)
// 		b++
// 		c--
// 	}
// 	// If hi-c<3 then there are duplicates (by property of median of nine).
// 	// Let be a bit more conservative, and set border to 5.
// 	protect := hi-c < 5
// 	if !protect && hi-c < (hi-lo)/4 {
// 		// Lets test some points for equality to pivot
// 		dups := 0
// 		if !data.Less(pivot, hi-1) { // data[hi-1] = pivot
// 			data.Swap(c, hi-1)
// 			c++
// 			dups++
// 		}
// 		if !data.Less(b-1, pivot) { // data[b-1] = pivot
// 			b--
// 			dups++
// 		}
// 		// m-lo = (hi-lo)/2 > 6
// 		// b-lo > (hi-lo)*3/4-1 > 8
// 		// ==> m < b ==> data[m] <= pivot
// 		if !data.Less(m, pivot) { // data[m] = pivot
// 			data.Swap(m, b-1)
// 			b--
// 			dups++
// 		}
// 		// if at least 2 points are equal to pivot, assume skewed distribution
// 		protect = dups > 1
// 	}
// 	if protect {
// 		// Protect against a lot of duplicates
// 		// Add invariant:
// 		//	data[a <= i < b] unexamined
// 		//	data[b <= i < c] = pivot
// 		for {
// 			for ; a < b && !data.Less(b-1, pivot); b-- { // data[b] == pivot
// 			}
// 			for ; a < b && data.Less(a, pivot); a++ { // data[a] < pivot
// 			}
// 			if a >= b {
// 				break
// 			}
// 			// data[a] == pivot; data[b-1] < pivot
// 			data.Swap(a, b-1)
// 			a++
// 			b--
// 		}
// 	}
// 	// Swap pivot into middle
// 	data.Swap(pivot, b-1)
// 	return b - 1, c
// }