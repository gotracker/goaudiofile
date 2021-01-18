package it

// Volume defines a volume value
type Volume uint8

// Value returns the value of the volume as a floating point value between 0 and 1, inclusively
func (p Volume) Value() float32 {
	switch {
	case p >= 0 && p <= 64:
		return float32(p) / 64
	default:
		panic("unexpected value")
	}
}

// FineVolume defines a volume value with double precision
type FineVolume uint8

// Value returns the value of the fine volume as a floating point value between 0 and 1, inclusively
func (p FineVolume) Value() float32 {
	switch {
	case p >= 0 && p <= 128:
		return float32(p) / 128
	default:
		panic("unexpected value")
	}
}

const (
	// DefaultVolume is the default volume for many things in IT files
	DefaultVolume = Volume(64)

	// DefaultFineVolume is the default volume for fine volumes in IT files
	DefaultFineVolume = Volume(128)
)

// PanSeparation is the panning separation value
type PanSeparation uint8

// Value returns the value of the panning separation as a floating point value between 0 and 1, inclusively
func (p PanSeparation) Value() float32 {
	switch {
	case p >= 0 && p <= 128:
		return float32(p) / 128
	default:
		panic("unexpected value")
	}
}

// PanValue describes a panning value in the IT format
type PanValue uint8

// IsSurround returns true if the panning is in surround-sound mode
func (p PanValue) IsSurround() bool {
	return (p &^ 128) == 100
}

// IsDisabled returns true if the channel this panning value is attached to is muted
// Effects in muted channels are still processed
func (p PanValue) IsDisabled() bool {
	return (p & 128) != 0
}

// Value returns the value of the panning as a floating point value between 0 and 1, inclusively
// 0 = absolute left, 0.5 = center, 1 = absolute right
func (p PanValue) Value() float32 {
	pv := p &^ 128
	switch {
	case pv >= 0 && pv <= 64:
		return float32(pv) / 64
	case pv == 100:
		return float32(0.5)
	default:
		panic("unexpected value")
	}
}

// NewNoteAction is what to do when a new note occurs
type NewNoteAction uint8

const (
	// NewNoteActionCut means to cut the previous playback when a new note occurs
	NewNoteActionCut = NewNoteAction(0)
	// NewNoteActionContinue means to continue the previous playback when a new note occurs
	NewNoteActionContinue = NewNoteAction(1)
	// NewNoteActionOff means to note-off the previous playback when a new note occurs
	NewNoteActionOff = NewNoteAction(2)
	// NewNoteActionFade means to fade the previous playback when a new note occurs
	NewNoteActionFade = NewNoteAction(3)
)

// Percentage8 is a percentage stored as a uint8
type Percentage8 uint8

// Value returns the value of the percentage
func (p Percentage8) Value() float32 {
	return float32(p) / 100
}

// DuplicateCheckType is the duplicate check type
type DuplicateCheckType uint8

const (
	// DuplicateCheckTypeOff is for when the duplicate check type is disabled
	DuplicateCheckTypeOff = DuplicateCheckType(0)
	// DuplicateCheckTypeNote is for when the duplicate check type is set to note mode
	DuplicateCheckTypeNote = DuplicateCheckType(1)
	// DuplicateCheckTypeSample is for when the duplicate check type is set to sample mode
	DuplicateCheckTypeSample = DuplicateCheckType(2)
	// DuplicateCheckTypeInstrument is for when the duplicate check type is set to instrument mode
	DuplicateCheckTypeInstrument = DuplicateCheckType(3)
)

// DuplicateCheckAction is the action to perform when a duplicate is detected
type DuplicateCheckAction uint8

const (
	// DuplicateCheckActionCut cuts the playback when a duplicate is detected
	DuplicateCheckActionCut = DuplicateCheckAction(0)
	// DuplicateCheckActionOff performs a note-off on the playback when a duplicate is detected
	DuplicateCheckActionOff = DuplicateCheckAction(1)
	// DuplicateCheckActionFade performs a fade-out on the playback when a duplicate is detected
	DuplicateCheckActionFade = DuplicateCheckAction(2)
)

// NodePoint16 is a node point in the old instrument format
type NodePoint16 struct {
	Tick      uint8
	Magnitude uint8
}

// NodePoint24 is a node point in the new instrument format
type NodePoint24 struct {
	Y    int8
	Tick uint16
}

// Envelope is an envelope structure
type Envelope struct {
	Flags            EnvelopeFlags
	Count            uint8
	LoopBegin        uint8
	LoopEnd          uint8
	SustainLoopBegin uint8
	SustainLoopEnd   uint8
	NodePoints       [25]NodePoint24
	Reserved51       uint8
}

// EnvelopeFlags is the flagset for new instrument envelopes
type EnvelopeFlags uint8

const (
	// EnvelopeFlagEnvelopeOn :: On = Use envelope
	EnvelopeFlagEnvelopeOn = EnvelopeFlags(1 << 0)
	// EnvelopeFlagLoopOn :: On = Use loop
	EnvelopeFlagLoopOn = EnvelopeFlags(1 << 1)
	// EnvelopeFlagSustainLoopOn :: On = Use sustain loop
	EnvelopeFlagSustainLoopOn = EnvelopeFlags(1 << 2)
)
