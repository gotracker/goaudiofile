package it

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"

	"github.com/gotracker/goaudiofile/internal/util"
)

// File is an IT internal file representation
type File struct {
	Head               ModuleHeader
	OrderList          []uint8
	InstrumentPointers []ParaPointer32
	SamplePointers     []ParaPointer32
	PatternPointers    []ParaPointer32
}

// Read reads an IT file from the reader `r` and creates an internal File representation
func Read(r io.Reader) (*File, error) {
	buffer := &bytes.Buffer{}
	if _, err := buffer.ReadFrom(r); err != nil {
		return nil, err
	}
	//data := buffer.Bytes()

	fh, err := ReadModuleHeader(buffer)
	if err != nil {
		return nil, err
	}
	if util.GetString(fh.IMPM[:]) != "IMPM" {
		return nil, errors.New("invalid file format")
	}

	f := File{
		Head:               *fh,
		OrderList:          make([]uint8, int(fh.OrderCount)),
		InstrumentPointers: make([]ParaPointer32, int(fh.InstrumentCount)),
		SamplePointers:     make([]ParaPointer32, int(fh.SampleCount)),
		PatternPointers:    make([]ParaPointer32, int(fh.PatternCount)),
	}
	if err := binary.Read(buffer, binary.LittleEndian, &f.OrderList); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.LittleEndian, &f.InstrumentPointers); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.LittleEndian, &f.SamplePointers); err != nil {
		return nil, err
	}
	if err := binary.Read(buffer, binary.LittleEndian, &f.PatternPointers); err != nil {
		return nil, err
	}

	// TODO: read instruments/samples

	// TODO: read patterns

	return &f, nil
}
