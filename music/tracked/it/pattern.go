package it

// PackedPattern is a packed pattern from the IT format
type PackedPattern struct {
	Length     uint16
	Rows       uint16
	Reserved04 [4]byte
	Data       []uint8
}

// ChannelData is the partially decoded channel data from the packed pattern
type ChannelData struct {
	Flags       ChannelDataFlags
	Note        Note
	Instrument  uint8
	VolPan      uint8
	Command     uint8
	CommandData uint8
}
