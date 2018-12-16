package Search

// # 逐次探索
// 先頭からしらみつぶして探しに行く検索法。イテレータを介さないと集合の要素にアクセスできない場合にはこのパターンになる。

func SequentialSearch(a []int, t int) bool {
	n := len(a)
	for i := 0; i < n; i++ {
		if a[i] == t {
			return true
		}
	}
	return false
}

func SequentialSearchWithIterator(a []int, t int) bool {
	intset := NewIntSet(a)
	it := intset.GetIterator()
	for it.HasNext() {
		e := it.Next()
		if e == t {
			return true
		}
	}
	return false
}

type Aggregate interface {
	GetIterator() *Iterator
}

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type IntSet struct {
	Aggregate
	set  []int
	last int
}

func NewIntSet(a []int) *IntSet {
	return &IntSet{set: a, last: len(a)}
}

func (is *IntSet) IntSet(max int) {
	is.set = make([]int, max)
}

func (is *IntSet) GetIntAt(index int) int {
	return is.set[index]
}

func (is *IntSet) Append(num int) {
	is.set[is.last] = num
	is.last++
}

func (is *IntSet) GetLength() int {
	return is.last
}

func (is *IntSet) GetIterator() *IntIterator {
	return NewIntIterator(is)
}

type IntIterator struct {
	Iterator
	set   *IntSet
	index int
}

func NewIntIterator(set *IntSet) *IntIterator {
	return &IntIterator{set: set, index: 0}
}
func (it *IntIterator) HasNext() bool {
	if it.index < it.set.GetLength() {
		return true
	} else {
		return false
	}
}

func (it *IntIterator) Next() int {
	x := it.set.GetIntAt(it.index)
	it.index++
	return x
}
