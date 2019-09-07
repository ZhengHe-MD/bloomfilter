package bloomfilter

type Bits []byte

const elemLen = 8

func NewBits(size uint64) Bits {
	l := size / elemLen
	if size % elemLen > 0 {
		l += 1
	}
	return make([]byte, l)
}

func (m Bits) Set(index uint64) {
	skips, offset := index/elemLen, index%elemLen
	m[skips] = m[skips] | (uint8(1) << (elemLen-offset-1))
}

func (m Bits) HasSet(index uint64) bool {
	skips, offset := index/elemLen, index%elemLen
	return m[skips] & (uint8(1) << (elemLen-offset-1)) > 0
}
