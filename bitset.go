package bloomfilter

type BitSet []byte

const elemLen = 8

func NewBitSet(size uint64) BitSet {
	l := size / elemLen
	if size%elemLen > 0 {
		l += 1
	}
	return make([]byte, l)
}

func (m BitSet) Set(index uint64) {
	skips, offset := index/elemLen, index%elemLen
	m[skips] = m[skips] | (uint8(1) << (elemLen - offset - 1))
}

func (m BitSet) HasSet(index uint64) bool {
	skips, offset := index/elemLen, index%elemLen
	return m[skips]&(uint8(1)<<(elemLen-offset-1)) > 0
}
