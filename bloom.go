package bloomfilter

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"math"
	"strconv"
)

type BloomFilter struct {
	kHash  int
	bitset BitSet
}

func NewBloomFilter(s, k uint64) *BloomFilter {
	return &BloomFilter{
		kHash:  int(k),
		bitset: NewBitSet(s),
	}
}

func (m *BloomFilter) hashToIndex(key string, ki uint64) uint64 {
	return (h1(key) + h2(key)*ki) % uint64(len(m.bitset)*elemLen)
}

func (m *BloomFilter) AddKey(key string) {
	for i := 0; i < m.kHash; i++ {
		hi := m.hashToIndex(key, uint64(i))
		m.bitset.Set(hi)
	}
}

func (m *BloomFilter) HasKey(key string) bool {
	for i := 0; i < m.kHash; i++ {
		hi := m.hashToIndex(key, uint64(i))
		if !m.bitset.HasSet(hi) {
			return false
		}
	}
	return true
}

func (m *BloomFilter) FalsePositiveRate(total uint64) float64 {
	exp := math.Pow(math.E, -float64(m.kHash)*float64(total)/float64(len(m.bitset))/float64(elemLen))
	return math.Pow(1-exp, float64(m.kHash))
}

func h1(key string) uint64 {
	hash := sha1.Sum([]byte(key))
	val, _ := strconv.ParseUint(hex.EncodeToString(hash[:8]), 16, 64)
	return val
}

func h2(key string) uint64 {
	hash := md5.Sum([]byte(key))
	val, _ := strconv.ParseUint(hex.EncodeToString(hash[:8]), 16, 64)
	return val
}
