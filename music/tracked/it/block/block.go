package block

import (
	"encoding/binary"
)

// Block is a block interface
type Block interface {
	FourCC() uint32
	Length() int // total length, including the magic identifier
}

type blockBase struct {
	Identifier [4]byte
	BlockLen   uint32
}

// FourCC returns the big-endian representation of the block identifier
func (b *blockBase) FourCC() uint32 {
	return binary.BigEndian.Uint32(b.Identifier[:])
}

// Length returns the size of the whole block
func (b *blockBase) Length() int {
	return 8 + int(b.BlockLen)
}

type Unknown struct {
	blockBase
}

// FourCC returns the big-endian representation of the block identifier
func (b *Unknown) FourCC() uint32 {
	return b.blockBase.FourCC()
}

// Length returns the size of the whole block
func (b *Unknown) Length() int {
	return b.blockBase.Length()
}
