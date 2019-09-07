package bloomfilter

import (
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	bf := NewBloomFilter(1024, 3)

	paragraph := `Alabama Alaska Arizona Arkansas California Colorado Connecticut
    Delaware Florida Georgia Hawaii Idaho Illinois Indiana Iowa Kansas
    Kentucky Louisiana Maine Maryland Massachusetts Michigan Minnesota
    Mississippi Missouri Montana Nebraska Nevada NewHampshire NewJersey
    NewMexico NewYork NorthCarolina NorthDakota Ohio Oklahoma Oregon
    Pennsylvania RhodeIsland SouthCarolina SouthDakota Tennessee Texas Utah
    Vermont Virginia Washington WestVirginia Wisconsin Wyoming`
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