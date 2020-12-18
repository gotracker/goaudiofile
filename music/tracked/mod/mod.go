package mod

import (
	"encoding/binary"
	"errors"
	"io"

	"github.com/heucuva/goaudiofile/internal/util"
)

// File is an MOD internal file representation
type File struct {
	Head     ModuleHeader
	Patterns []Pattern
	Samples  []SampleData
}

type formatIntf interface {
	readPattern(*modFormatDetails, io.Reader) (*Pattern, error)
	rectifyOrderList(*modFormatDetails, [128]uint8) ([128]uint8, error)
}

type modFormatDetails struct {
	sig      string
	channels int
	format   formatIntf
}

var (
	sigChannels = [...]modFormatDetails{
		// amiga noisetracker / protracker
		{"M.K.", 4, protracker}, {"M!K!", 4, protracker},
		// startracker (startrekker?)
		{"FLT4", 4, startrekker}, {"FLT8", 8, startrekker},
		// fasttracker
		{"2CHN", 2, fasttracker}, {"4CHN", 4, fasttracker},
		{"6CHN", 6, fasttracker}, {"8CHN", 8, fasttracker},
		// fasttracker 2
		{"10CH", 10, fasttracker}, {"11CH", 11, fasttracker},
		{"12CH", 12, fasttracker}, {"13CH", 13, fasttracker},
		{"14CH", 14, fasttracker}, {"15CH", 15, fasttracker},
		{"16CH", 16, fasttracker}, {"17CH", 17, fasttracker},
		{"18CH", 18, fasttracker}, {"19CH", 19, fasttracker},
		{"20CH", 20, fasttracker}, {"21CH", 21, fasttracker},
		{"22CH", 22, fasttracker}, {"23CH", 23, fasttracker},
		{"24CH", 24, fasttracker}, {"25CH", 25, fasttracker},
		{"26CH", 26, fasttracker}, {"27CH", 27, fasttracker},
		{"28CH", 28, fasttracker}, {"29CH", 29, fasttracker},
		{"30CH", 30, fasttracker}, {"31CH", 31, fasttracker},
		{"32CH", 32, fasttracker},
	}
)

// Read reads a MOD file from the reader `r` and creates an internal MOD File representation
func Read(r io.Reader) (*File, error) {
	f := File{}

	if err := binary.Read(r, binary.LittleEndian, &f.Head); err != nil {
		return nil, err
	}

	sig := util.GetString(f.Head.Sig[:])
	var ffmt *modFormatDetails
	for _, s := range sigChannels {
		if s.sig == sig {
			ffmt = &s
			break
		}
	}

	if ffmt == nil || ffmt.channels == 0 {
		return nil, errors.New("invalid file format")
	}

	processor := ffmt.format
	if processor == nil {
		return nil, errors.New("could not identify format reader")
	}

	numPatterns := 0
	orderList, err := processor.rectifyOrderList(ffmt, f.Head.Order)
	if err != nil {
		return nil, err
	}
	for i, o := range orderList {
		if i < int(f.Head.SongLen) {
			f.Head.Order[i] = o
		}
		// we count all patterns, even if we're not in the 'song' range
		// hidden/'deleted' patterns can exist...
		if numPatterns <= int(o) {
			numPatterns = int(o) + 1
		}
	}

	f.Patterns = make([]Pattern, numPatterns)
	for i := 0; i < numPatterns; i++ {
		pattern, err := processor.readPattern(ffmt, r)
		if err != nil {
			return nil, err
		}
		if pattern == nil {
			continue
		}
		f.Patterns[i] = *pattern
	}

	f.Samples = make([]SampleData, len(f.Head.Instrument))
	for instNum, inst := range f.Head.Instrument {
		samp := make([]byte, inst.Len.Value())
		if err := binary.Read(r, binary.LittleEndian, &samp); err != nil {
			return nil, err
		}
		f.Samples[instNum] = samp
	}

	return &f, nil
}
