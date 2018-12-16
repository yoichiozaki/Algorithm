package Search

import (
	"crypto/md5"
	"encoding/hex"
	"math/big"
	"strconv"
)

// # ブルームフィルタ
// ブルームフィルタとは確率的データ構造の一種で、あるデータが集合に属するかどうかを判定する際に使われる。メモリの空間消費が少なくすみ、
// 非常に効率的にデータの存在判定ができる。
// ブルームフィルタはこれまでの探索とは毛色が違う。ブルームフィルタでわかるのは「その要素はあるかもしれない」ということだけである。
// これは偽陽性と呼ばれる。つまりブルームフィルタを用いれば、存在するものについては確実にtrueが返ってくるが、集合に存在しない要素に対して
// trueが返ってくる可能性があるということである。上記の誤りを許容できる環境であれば使える探索法である。

const (
	k = 3
)

type BloomFilter []bool

func NewBloomFilter(size int) BloomFilter {
	return make([]bool, size)
}

func (bf *BloomFilter) Add(element string) {
	hash := md5Hash(element)
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]
	i64HashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64HashB, _ := strconv.ParseInt(hashB, 16, 64)

	for i := 0; i < k; i++ {
		(*bf)[doubleHashing(i64HashA, i64HashB, i, len(*bf))] = true
	}
}

func (bf *BloomFilter) Exists(element string) bool {
	hash := md5Hash(element)
	hashA := hash[:int(len(hash)/2)]
	hashB := hash[int(len(hash)/2):]
	i64HashA, _ := strconv.ParseInt(hashA, 16, 64)
	i64HashB, _ := strconv.ParseInt(hashB, 16, 64)
	for i := 0; i < k; i++ {
		if !(*bf)[doubleHashing(i64HashA, i64HashB, i, len(*bf))] {
			return false
		}
	}
	return true // いるかもしれない
}

func md5Hash(str string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(str))
	return hex.EncodeToString(md5Hash.Sum(nil))
}

func doubleHashing(hashA, hashB int64, n, length int) (hash int64) {
	h := new(big.Int).Mul(big.NewInt(int64(n)), big.NewInt(hashB))
	h = new(big.Int).Add(big.NewInt(hashA), h)
	h = new(big.Int).Rem(h, big.NewInt(int64(length)))
	// if the rem is negative, make it positive.
	hash = h.Int64()
	if hash < 0 {
		hash += int64(length)
	}
	return
}