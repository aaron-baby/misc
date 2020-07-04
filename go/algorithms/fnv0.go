package algorithms

import "hash"

type (
	sum32 uint32
)

const (
	offset32 = 0
	prime32  = 16777619
)

func (s *sum32) Reset()         { *s = offset32 }
func (s *sum32) Sum32() uint32  { return uint32(*s) }
func (s *sum32) Size() int      { return 4 }
func (s *sum32) BlockSize() int { return 1 }
func (s *sum32) Sum(in []byte) []byte {
	v := uint32(*s)
	return append(in, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
}

func NewFNV0() hash.Hash32 {
	var s sum32 = offset32
	return &s
}
func (s *sum32) Write(data []byte) (int, error) {
	hash := *s
	for _, c := range data {
		hash *= prime32
		hash ^= sum32(c)
	}
	*s = hash
	return len(data), nil
}
