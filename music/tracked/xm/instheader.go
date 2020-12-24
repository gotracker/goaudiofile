package xm

import (
	"encoding/binary"
	"io"

	"github.com/gotracker/goaudiofile/internal/util"
)

// InstrumentHeader is a representation of the XM file instrument header
type InstrumentHeader struct {
	Size         uint32
	Name         [22]uint8
	Type         uint8
	SamplesCount uint16

	SampleHeaderSize uint32
	SampleNumber     [96]uint8
	VolEnv           [12]EnvPoint
	PanEnv           [12]EnvPoint

	VolPoints         uint8
	PanPoints         uint8
	VolSustainPoint   uint8
	VolLoopStartPoint uint8
	VolLoopEndPoint   uint8
	PanSustainPoint   uint8
	PanLoopStartPoint uint8
	PanLoopEndPoint   uint8
	VolFlags          EnvelopeFlags
	PanFlags          EnvelopeFlags
	VibratoType       uint8
	VibratoSweep      uint8
	VibratoDepth      uint8
	VibratoRate       uint8
	VolumeFadeout     uint16
	ReservedP241      [11]uint16

	Samples []SampleHeader
}

// GetName returns a string representation of the data stored in the Name field
func (ih *InstrumentHeader) GetName() string {
	return util.GetString(ih.Name[:])
}

// EnvelopeFlags is a representation of the XM file instrument envelope flags (vol/pan)
type EnvelopeFlags uint8

const (
	// EnvelopeFlagEnabled activates the envelope
	EnvelopeFlagEnabled = EnvelopeFlags(0x01)
	// EnvelopeFlagSustainEnabled enables the sustain segment of the envelope
	EnvelopeFlagSustainEnabled = EnvelopeFlags(0x02)
	// EnvelopeFlagLoopEnabled enables the loop function of the envelope
	EnvelopeFlagLoopEnabled = EnvelopeFlags(0x04)
)

// IsEnabled returns true if the envelope is enabled
func (f EnvelopeFlags) IsEnabled() bool {
	return (f & EnvelopeFlagEnabled) != 0
}

// IsSustainEnabled returns true if the envelope's sustain function is enabled
func (f EnvelopeFlags) IsSustainEnabled() bool {
	return (f & EnvelopeFlagSustainEnabled) != 0
}

// IsLoopEnabled returns true if the envelope's loop function is enabled
func (f EnvelopeFlags) IsLoopEnabled() bool {
	return (f & EnvelopeFlagLoopEnabled) != 0
}

// EnvPoint is a representation of an XM file envelope point
type EnvPoint struct {
	X uint16
	Y uint16
}

// SampleHeader is a representation of the XM file sample header
type SampleHeader struct {
	Length             uint32
	LoopStart          uint32
	LoopLength         uint32
	Volume             uint8
	Finetune           int8
	Flags              SampleFlags
	Panning            uint8
	RelativeNoteNumber int8
	ReservedP17        uint8
	Name               [22]uint8
	SampleData         []int8
}

// GetName returns a string representation of the data stored in the Name field
func (sh *SampleHeader) GetName() string {
	return util.GetString(sh.Name[:])
}

// SampleFlags is a representation of the XM file sample flags
type SampleFlags uint8

const (
	// sampleFlagLoopModeMask is the mask to pull the loop mode from the sample flags
	sampleFlagLoopModeMask = SampleFlags(0x03)
	// SampleFlag16Bit designates that the sample is 16-bit
	SampleFlag16Bit = SampleFlags(0x10)
	// SampleFlagStereo designates that the sample is stereo
	SampleFlagStereo = EnvelopeFlags(0x20)
)

// LoopMode returns the loop mode described by the sample flags
func (f SampleFlags) LoopMode() SampleLoopMode {
	return SampleLoopMode(f & sampleFlagLoopModeMask)
}

// Is16Bit returns true if the sample is 16-bit
func (f SampleFlags) Is16Bit() bool {
	return (f & SampleFlag16Bit) != 0
}

// IsStereo returns true if the sample is stereo
func (f EnvelopeFlags) IsStereo() bool {
	return (f & SampleFlagStereo) != 0
}

// SampleLoopMode is a representation of the XM file sample loop mode
type SampleLoopMode uint8

const (
	// SampleLoopModeDisabled is no loop mode
	SampleLoopModeDisabled = SampleLoopMode(0x00)
	// SampleLoopModeEnabled describes a loop mode where the sample plays from start to loopEnd, then repeats from loopBegin to loopEnd
	SampleLoopModeEnabled = SampleLoopMode(0x01)
	// SampleLoopModePingPong describes a loop mode where the sample plays from start to loopend, then inverts playback from loopEnd to
	// loopBegin, then inverts back (and continues this way)
	SampleLoopModePingPong = SampleLoopMode(0x02)
	// SampleLoopModeUnknown is an invalid/unknown loop mode
	SampleLoopModeUnknown = SampleLoopMode(0x03)
)

func readInstrumentHeaderPartial(r io.Reader) (*InstrumentHeader, error) {
	ih := InstrumentHeader{}

	if err := binary.Read(r, binary.LittleEndian, &ih.Size); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.Name); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.Type); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.SamplesCount); err != nil {
		return nil, err
	}

	if ih.SamplesCount == 0 {
		// it's empty - we're done!
		return &ih, nil
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.SampleHeaderSize); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.SampleNumber); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.VolEnv); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.PanEnv); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.VolPoints); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.PanPoints); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.VolSustainPoint); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.VolLoopStartPoint); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.VolLoopEndPoint); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.PanSustainPoint); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.PanLoopStartPoint); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.PanLoopEndPoint); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.VolFlags); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.PanFlags); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.VibratoType); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.VibratoSweep); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.VibratoDepth); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.VibratoRate); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.VolumeFadeout); err != nil {
		return nil, err
	}

	if err := binary.Read(r, binary.LittleEndian, &ih.ReservedP241); err != nil {
		return nil, err
	}

	return &ih, nil
}

func readInstrumentHeader(r io.Reader) (*InstrumentHeader, error) {
	ih, err := readInstrumentHeaderPartial(r)
	if err != nil {
		return nil, err
	}

	for i := uint16(0); i < ih.SamplesCount; i++ {
		s := SampleHeader{}

		if err := binary.Read(r, binary.LittleEndian, &s.Length); err != nil {
			return nil, err
		}

		if err := binary.Read(r, binary.LittleEndian, &s.LoopStart); err != nil {
			return nil, err
		}

		if err := binary.Read(r, binary.LittleEndian, &s.LoopLength); err != nil {
			return nil, err
		}

		if err := binary.Read(r, binary.LittleEndian, &s.Volume); err != nil {
			return nil, err
		}

		if err := binary.Read(r, binary.LittleEndian, &s.Finetune); err != nil {
			return nil, err
		}

		if err := binary.Read(r, binary.LittleEndian, &s.Flags); err != nil {
			return nil, err
		}

		if err := binary.Read(r, binary.LittleEndian, &s.Panning); err != nil {
			return nil, err
		}

		if err := binary.Read(r, binary.LittleEndian, &s.RelativeNoteNumber); err != nil {
			return nil, err
		}

		if err := binary.Read(r, binary.LittleEndian, &s.ReservedP17); err != nil {
			return nil, err
		}

		if err := binary.Read(r, binary.LittleEndian, &s.Name); err != nil {
			return nil, err
		}

		s.SampleData = make([]int8, int(s.Length))

		ih.Samples = append(ih.Samples, s)
	}

	for _, s := range ih.Samples {
		data := s.SampleData
		if err := binary.Read(r, binary.LittleEndian, &data); err != nil {
			return nil, err
		}

		// convert the sample in the background
		go func() {
			old := int8(0)
			for i, s := range data {
				new := s + old
				data[i] = new
				old = new
			}
		}()
	}
	return ih, nil
}
