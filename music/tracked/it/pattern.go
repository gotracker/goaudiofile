package it

import (
	"bytes"
	"encoding/binary"
	"errors"
)

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

func readPackedPattern(data []byte, ptr ParaPointer, cmwt uint16) (*PackedPattern, error) {
	ofs := ptr.Offset()
	if ofs == 0 {
		p := PackedPattern{
			Length: 64,
			Rows:   64,
			Data:   make([]uint8, 64), // filled with zeroes is desired
		}
		return &p, nil
	}

	var p PackedPattern
	r := bytes.NewBuffer(data[ofs:])
	if err := binary.Read(r, binary.LittleEndian, &p.Length); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.Rows); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &p.Reserved04); err != nil {
		return nil, err
	}

	p.Data = make([]uint8, int(p.Length))
	if err := binary.Read(r, binary.LittleEndian, &p.Data); err != nil {
		return nil, err
	}

	return &p, nil
}

// ReadChannelData decodes a packed pattern from the position indicated (`pos`) and returns an
// integer equal to the number of bytes used to decode the channel data, a ChannelData structure,
// and a possible error value.
// If there is no error and the row is completed, the ChannelData structure returned will be nil.
func (p *PackedPattern) ReadChannelData(pos int) (int, *ChannelData, error) {
	if pos > len(p.Data) {
		return 0, nil, errors.New("position out of bounds")
	}

	var cd ChannelData

	r := bytes.NewBuffer(p.Data[pos:])
	s := 0
	if err := binary.Read(r, binary.LittleEndian, &cd.Flags); err != nil {
		return 0, nil, err
	}
	s++

	if cd.Flags.HasNote() {
		if err := binary.Read(r, binary.LittleEndian, &cd.Note); err != nil {
			return 0, nil, err
		}
		s++
	}

	if cd.Flags.HasInstrument() {
		if err := binary.Read(r, binary.LittleEndian, &cd.Instrument); err != nil {
			return 0, nil, err
		}
		s++
	}

	if cd.Flags.HasVolPan() {
		if err := binary.Read(r, binary.LittleEndian, &cd.VolPan); err != nil {
			return 0, nil, err
		}
		s++
	}

	if cd.Flags.HasCommand() {
		if err := binary.Read(r, binary.LittleEndian, &cd.Command); err != nil {
			return 0, nil, err
		}
		s++
		if err := binary.Read(r, binary.LittleEndian, &cd.CommandData); err != nil {
			return 0, nil, err
		}
		s++
	}

	return s, &cd, nil
}
