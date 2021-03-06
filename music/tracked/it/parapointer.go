package it

// ParaPointer is a pointer offset within the IT file format
type ParaPointer interface {
	Offset() int
}

// ParaPointer16 is a 16-bit pointer offset within the IT file format
type ParaPointer16 uint16

// Offset returns the actual offset
func (p ParaPointer16) Offset() int {
	return int(p) << 4
}

// ParaPointer24 is a 24-bit pointer offset within the IT file format
type ParaPointer24 struct {
	Hi uint8
	Lo ParaPointer16
}

// Offset returns the actual offset
func (p ParaPointer24) Offset() int {
	return (int(p.Hi)<<16 | int(p.Lo)) << 4
}

// ParaPointer32 is a 32-bit pointer offset within the IT file format
type ParaPointer32 uint32

// Offset returns the actual offset
func (p ParaPointer32) Offset() int {
	return int(p)
}
