package Search

import (
	"log"
	"os"
)

// # ハッシュに基づいた探索
// 任意長のビット列から規則性のない固定長のビット列を生成する関数であるハッシュ関数を用いることでデータを事前に「分類」しておくことで、
// ほぼダイレクトに探索対象にたどり着くことができる。ハッシュ関数の出力が衝突した時にどのように対処するかによっていくつかのバリエーションが
// ある。具体的にはリンク付きリストを用いる方法やオープンアドレス法がある。ハッシュに基づいた探索は時間計算量については優秀であるが、
// 空間計算量の観点からでは無駄が多いとされる。無限に広いハッシュテーブルを用意すれば探索自体はO(1)の時間で済むが、現実的ではない。

func HashBasedSearch(a []int, table Table, t int, method CollisionAvoidanceMethod) bool {
	h := hash(t)
	switch method {
	case LinkedList:
		if table[h] == nil {
			return false
		}
		for tmp := table[h]; tmp != nil; tmp = tmp.next {
			if tmp.element == t {
				return true
			}
		}
		return false
	// TODO: open address method is not yet completed.
	case OpenAddress:
		now := h
		for b := 0; b < 100; b++ {
			if table[now].element == t {
				return true
			} else {
				now = (now + b + b*b) % 100
			}
		}
		return false
	default: // never reach here
		log.Println("Error: No such method")
		return false
	}
}

type CollisionAvoidanceMethod int

const (
	LinkedList CollisionAvoidanceMethod = iota
	OpenAddress
)

// const SIZE = 262143 // int(math.Pow(2, 18) - 1)
const SIZE = 1000

type entry struct {
	element int
	next    *entry
}

func newEntry(e int) *entry {
	return &entry{element: e, next: nil}
}

type Table []*entry

func NewHashTable() Table {
	return make([]*entry, SIZE)
}

func (t *Table) Load(a []int, method CollisionAvoidanceMethod) {
	switch method {
	case LinkedList:
		for _, e := range a {
			h := hash(e)
			if (*t)[h] == nil {
				(*t)[h] = newEntry(e)
			} else {
				ne := newEntry(e)
				ne.next = (*t)[h]
				(*t)[h] = ne
			}
		}
	// TODO: open address method is not yet completed.
	case OpenAddress:
	NextElement:
		for _, e := range a {
			h := hash(e)
			if (*t)[h] == nil {
				(*t)[h] = newEntry(e)
			} else {
				i := 0
				for (*t)[h] != nil {
					h = (h + i + i*i) % 100
					if (*t)[h] == nil {
						(*t)[h] = newEntry(e)
						continue NextElement
					}
					i++
				}
			}
		}
	default:
		log.Println("Error: No such method")
		os.Exit(0)
	}
}

func hash(x int) int {
	return x % SIZE
}
