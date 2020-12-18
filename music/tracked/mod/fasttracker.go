package mod

import (
	"encoding/binary"
	"errors"
	"io"
)

type fmtFT struct {
	formatIntf
}

var (
	fasttracker = &fmtFT{}
)

func (f *fmtFT) readPattern(ffmt *modFormatDetails, r io.Reader) (*Pattern, error) {
	if r == nil {
		return nil, errors.New("r is nil")
	}

	p := NewPattern(ffmt.channels)
	for _, row := range p {
		for c := 0; c < ffmt.channels; c++ {
			if err := binary.Read(r, binary.LittleEndian, &row[c]); err != nil {
				return nil, err
			}
		}
	}

	return &p, nil
}

func (f *fmtFT) rectifyOrderList(ffmt *modFormatDetails, in [128]uint8) ([128]uint8, error) {
	return in, nil
}
