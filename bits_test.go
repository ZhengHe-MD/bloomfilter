package bloomfilter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBits(t *testing.T) {
	cases := []struct {
		size         uint64
		expectedSize int
	}{
		{16, 2},
		{17, 3},
	}

	for _, c := range cases {
		bits := NewBits(c.size)
		assert.Equal(t, c.expectedSize, len(bits))
	}
}

func TestBits_Set(t *testing.T) {
	cases := []struct {
		idx          uint64
		expectedBits []uint8
	}{
		{0, []uint8{1 << 7, 0}},
		{1, []uint8{1 << 6, 0}},
		{2, []uint8{1 << 5, 0}},
		{4, []uint8{1 << 3, 0}},
		{7, []uint8{1 << 0, 0}},
		{8, []uint8{0, 1<<7}},
		{9, []uint8{0, 1<<6}},
		{15, []uint8{0, 1<<0}},
	}

	for _, c := range cases {
		bits := NewBits(16)
		bits.Set(c.idx)
		assert.Equal(t, []uint8(bits), c.expectedBits)
	}
}

func TestBits_HasSet(t *testing.T) {
	bits := NewBits(128)

	testSet := map[uint64]struct{}{
		0:   {},
		7:   {},
		8:   {},
		35:  {},
		64:  {},
		71:  {},
		127: {},
	}

	for k, _ := range testSet {
		bits.Set(k)
	}

	for i := 0; i < 128; i++ {
		if _, ok := testSet[uint64(i)]; ok {
			assert.True(t, bits.HasSet(uint64(i)), i)
		} else {
			assert.False(t, bits.HasSet(uint64(i)), i)
		}
	}
}
