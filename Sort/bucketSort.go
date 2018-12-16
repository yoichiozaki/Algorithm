package Sort

// # バケットソート
// 要素を比較することで整列するアルゴリズムではn個の要素を整列するのにO(n log n)よりも良い性能を出すことはできないことが数学的に証明されている。
// 一方で、「比較することなしに」整列させるアルゴリズムではO(n)で実現することができる。比較することなしに整列するためには要素の集まりを
// 一様に分割するハッシュ関数があれば実現できる。このハッシュ関数は順序のついたハッシュ関数でi<jの時バケツiに挿入された要素はバケツjに挿入された
// 要素よりも小さくなければならない。

func BucketSort(a []int, max int) []int {
	num := numBuckets(max)
	buckets := make([]bucket, num)
	for i := 0; i < len(a); i++ {
		k := hash(a[i])
		buckets[k].elements = append(buckets[k].elements, a[i])
	}
	return extract(buckets, &a, max)
}

type bucket struct {
	elements []int
}

func extract(buckets []bucket, a *[]int, max int) []int {
	idx := 0
	num := numBuckets(max)
	for i := 0; i < num; i++ {
		InsertionSort(buckets[i].elements)
		for _, e := range buckets[i].elements {
			(*a)[idx] = e
			idx++
		}
	}
	return *a
}

func numBuckets(numElements int) int {
	num := numElements
	return num
}

func hash(ele int) int {
	return ele
}
