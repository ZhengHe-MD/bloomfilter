package bloomfilter

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := NewBloomFilter(1024, 3)

	paragraph := `A Bloom filter is a space-efficient probabilistic data structure, conceived by Burton Howard Bloom in 1970, that is used to test whether an element is a member of a set. False positive matches are possible, but false negatives are not – in other words, a query returns either "possibly in set" or "definitely not in set". Elements can be added to the set, but not removed (though this can be addressed with a "counting" filter); the more elements that are added to the set, the larger the probability of false positives`
	var words []string
	for _, w := range strings.Split(paragraph, " ") {
		w = strings.TrimSpace(w)
		if len(w) > 0 {
			words = append(words, w)
		}
	}

	for _, word := range words {
		bf.AddKey(word)
	}

	for _, word := range words {
		assert.True(t, bf.HasKey(word), word)
	}
	assert.False(t, bf.HasKey("world"), "world")
	log.Println(bf.FalsePositiveRate(uint64(len(words))))
}
