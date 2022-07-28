package it

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"io/ioutil"

	"github.com/gotracker/goaudiofile/internal/util"
	"github.com/gotracker/goaudiofile/music/tracked/it/block"
)

var (
	// ErrInvalidFileFormat is for when an invalid file format is encountered
	ErrInvalidFileFormat = errors.New("invalid file format")
)

// File is an IT internal file representation
type File struct {
	Head               ModuleHeader
	OrderList          []uint8
	InstrumentPointers []ParaPointer32
	SamplePointers     []ParaPointer32
	PatternPointers    []ParaPointer32
	Instruments        []IMPIIntf
	Samples            []FullSample
	Patterns           []PackedPattern
	EditHistory        []EditHistory
	Blocks             []block.Block
}

// FullSample is a full sample, header + data
type FullSample struct {
	Header Sample
	Data   []byte
}

// Read reads an IT file from the reader `r` and creates an internal File representation
func Read(r io.Reader) (*File, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewReader(data)

	fh, err := ReadModuleHeader(buffer)
	if err != nil {
		return nil, err
	}
	if util.GetString(fh.IMPM[:]) != "IMPM" {
		return nil, ErrInvalidFileFormat
	}

	f := File{
		Head:               *fh,
		OrderList:          make([]uint8, int(fh.OrderCount)),
		InstrumentPointers: make([]ParaPointer32, int(fh.InstrumentCount)),
		SamplePointers:     make([]ParaPointer32, int(fh.SampleCount)),
		PatternPointers:    make([]ParaPointer32, int(fh.PatternCount)),
		Instruments:        make([]IMPIIntf, 0),
		Samples:            make([]FullSample, 0),
		Patterns:           make([]PackedPattern, 0),
		Blocks:             make([]block.Block, 0),
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

	// the earliest valid position to read from
	valPos := ParaPointer32(len(data))

	for _, pp := range f.InstrumentPointers {
		if pp > 0 && pp < valPos {
			valPos = pp
		}
	}
	for _, pp := range f.SamplePointers {
		if pp > 0 && pp < valPos {
			valPos = pp
		}
	}
	for _, pp := range f.PatternPointers {
		if pp > 0 && pp < valPos {
			valPos = pp
		}
	}

	if f.Head.SpecialFlags.IsMessageAttached() {
		if valPos > f.Head.MessageOffset {
			valPos = f.Head.MessageOffset
		}
	}

	if f.Head.SpecialFlags.IsHistoryIncluded() {
		curPos, _ := buffer.Seek(0, io.SeekCurrent)
		var historyParaLen uint16
		if err := binary.Read(buffer, binary.LittleEndian, &historyParaLen); err != nil {
			return nil, err
		}

		histBlockEnd := int(curPos) + int(historyParaLen)*8 + 2
		if histBlockEnd < len(data) {
			var h EditHistory
			if err := binary.Read(buffer, binary.LittleEndian, &h); err != nil {
				return nil, err
			}
			f.EditHistory = append(f.EditHistory, h)
		} else {
			buffer.Seek(curPos, io.SeekStart)
		}
	}

	nextValPos := valPos
blockReadLoop:
	for {
		block, err := readBlock(data, nextValPos, f.Head.TrackerCompatVersion)
		if err != nil || block == nil {
			break blockReadLoop
		}

		blen := block.Length()
		if blen < 8 {
			break blockReadLoop
		}

		if block.FourCC() == 0x494d5049 { // IMPI
			break blockReadLoop
		}

		f.Blocks = append(f.Blocks, block)
		nextValPos += ParaPointer32(blen)

		if nextValPos.Offset() < len(data) {
			valPos = nextValPos
		}
	}

	for _, ptr := range f.InstrumentPointers {
		if ptr < valPos {
			return nil, ErrInvalidFileFormat
		}

		impi, err := readIMPI(data, ptr, f.Head.TrackerCompatVersion)
		if err != nil {
			return nil, ErrInvalidFileFormat
		}
		f.Instruments = append(f.Instruments, impi)
	}

	for _, ptr := range f.SamplePointers {
		if ptr < valPos {
			return nil, ErrInvalidFileFormat
		}

		imps, err := readIMPS(data, ptr, f.Head.TrackerCompatVersion)
		if err != nil {
			return nil, ErrInvalidFileFormat
		}

		fs := FullSample{
			Header: *imps,
			Data:   make([]byte, 0),
		}

		if fs.Header.Flags.DoesSampleExist() {
			slen := fs.Header.Length
			if fs.Header.Flags.Is16Bit() {
				slen *= 2
			}
			if fs.Header.Flags.IsStereo() {
				slen *= 2
			}

			fs.Data = make([]byte, slen)
			if err := readSampleData(data, fs.Header.SamplePointer, f.Head.TrackerCompatVersion, fs.Data); err != nil {
				return nil, err
			}
		}

		f.Samples = append(f.Samples, fs)
	}

	for _, ptr := range f.PatternPointers {
		if ptr == 0 {
			f.Patterns = append(f.Patterns, PackedPattern{})
		} else if ptr < valPos {
			return nil, ErrInvalidFileFormat
		}

		pat, err := readPackedPattern(data, ptr, f.Head.TrackerCompatVersion)
		if err != nil {
			return nil, ErrInvalidFileFormat
		}
		f.Patterns = append(f.Patterns, *pat)
	}

	return &f, nil
}
