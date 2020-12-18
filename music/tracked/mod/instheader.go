package mod

// InstrumentHeader is a representation of the MOD file instrument header
type InstrumentHeader struct {
	Name      [22]byte
	Len       WordLength
	FineTune  uint8
	Volume    uint8
	LoopStart WordLength
	LoopEnd   WordLength
}

// SampleData is the data associated to the instrument
type SampleData []uint8
