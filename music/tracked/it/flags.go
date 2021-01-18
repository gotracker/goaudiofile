package it

// IMPMFlags is a set of flags describing various features in the IT file
type IMPMFlags uint16

const (
	// IMPMFlagStereo :: On = Stereo, Off = Mono (panning enablement flag)
	IMPMFlagStereo = IMPMFlags(1 << 0)
	// IMPMFlagVol0Optimizations :: If on, no mixing occurs if the volume at mixing time is 0 (redundant v1.04+)
	IMPMFlagVol0Optimizations = IMPMFlags(1 << 1)
	// IMPMFlagUseInstruments :: On = Use instruments, Off = Use samples
	IMPMFlagUseInstruments = IMPMFlags(1 << 2)
	// IMPMFlagLinearSlides :: On = Linear slides, Off = Amiga slides
	IMPMFlagLinearSlides = IMPMFlags(1 << 3)
	// IMPMFlagOldEffects :: On = Old Effects, Off = IT Effects
	IMPMFlagOldEffects = IMPMFlags(1 << 4)
	// IMPMFlagEFGLinking :: On = Link Effect G's memory with Effect E/F
	IMPMFlagEFGLinking = IMPMFlags(1 << 5)
	// IMPMFlagMidiPitchController :: Use MIDI pitch controller, Pitch depth given by PitchWheelDepth
	IMPMFlagMidiPitchController = IMPMFlags(1 << 6)
	// IMPMFlagReqEmbedMidi :: Request embedded MIDI configuration
	IMPMFlagReqEmbedMidi = IMPMFlags(1 << 7)
)

// IsStereo returns true if stereo (panning) is enabled
func (f IMPMFlags) IsStereo() bool {
	return (f & IMPMFlagStereo) != 0
}

// IsVol0Optimizations returns true if vol-0 optimization is enabled
func (f IMPMFlags) IsVol0Optimizations() bool {
	return (f & IMPMFlagVol0Optimizations) != 0
}

// IsUseInstruments returns true if use-instruments (instead of samples) is enabled
func (f IMPMFlags) IsUseInstruments() bool {
	return (f & IMPMFlagUseInstruments) != 0
}

// IsLinearSlides returns true if linear slides is enabled
func (f IMPMFlags) IsLinearSlides() bool {
	return (f & IMPMFlagLinearSlides) != 0
}

// IsOldEffects returns true if old-style effects are enabled
func (f IMPMFlags) IsOldEffects() bool {
	return (f & IMPMFlagLinearSlides) != 0
}

// IsEFGLinking returns true if effect E/F/G linking is enabled
func (f IMPMFlags) IsEFGLinking() bool {
	return (f & IMPMFlagEFGLinking) != 0
}

// IsMidiPitchController returns true if midi pitch controller is enabled
func (f IMPMFlags) IsMidiPitchController() bool {
	return (f & IMPMFlagMidiPitchController) != 0
}

// IsReqEmbedMidi returns true if request embedded midi configuration is enabled
func (f IMPMFlags) IsReqEmbedMidi() bool {
	return (f & IMPMFlagReqEmbedMidi) != 0
}

// IMPMSpecialFlags is a set of flags describing various special features in the IT file
type IMPMSpecialFlags uint16

const (
	// IMPMSpecialFlagMessageAttached :: On = song message attached
	IMPMSpecialFlagMessageAttached = IMPMSpecialFlags(1 << 0)
	// IMPMSpecialFlagReservedBit1 :: Reserved
	IMPMSpecialFlagReservedBit1 = IMPMSpecialFlags(1 << 1)
	// IMPMSpecialFlagReservedBit2 :: Reserved
	IMPMSpecialFlagReservedBit2 = IMPMSpecialFlags(1 << 2)
	// IMPMSpecialFlagEmbedMidi :: MIDI configuration embedded
	IMPMSpecialFlagEmbedMidi = IMPMSpecialFlags(1 << 3)
	// IMPMSpecialFlagReservedBit4 :: Reserved
	IMPMSpecialFlagReservedBit4 = IMPMSpecialFlags(1 << 4)
	// IMPMSpecialFlagReservedBit5 :: Reserved
	IMPMSpecialFlagReservedBit5 = IMPMSpecialFlags(1 << 5)
	// IMPMSpecialFlagReservedBit6 :: Reserved
	IMPMSpecialFlagReservedBit6 = IMPMSpecialFlags(1 << 6)
	// IMPMSpecialFlagReservedBit7 :: Reserved
	IMPMSpecialFlagReservedBit7 = IMPMSpecialFlags(1 << 7)
	// IMPMSpecialFlagReservedBit8 :: Reserved
	IMPMSpecialFlagReservedBit8 = IMPMSpecialFlags(1 << 8)
	// IMPMSpecialFlagReservedBit9 :: Reserved
	IMPMSpecialFlagReservedBit9 = IMPMSpecialFlags(1 << 9)
	// IMPMSpecialFlagReservedBit10 :: Reserved
	IMPMSpecialFlagReservedBit10 = IMPMSpecialFlags(1 << 10)
	// IMPMSpecialFlagReservedBit11 :: Reserved
	IMPMSpecialFlagReservedBit11 = IMPMSpecialFlags(1 << 11)
	// IMPMSpecialFlagReservedBit12 :: Reserved
	IMPMSpecialFlagReservedBit12 = IMPMSpecialFlags(1 << 12)
	// IMPMSpecialFlagReservedBit13 :: Reserved
	IMPMSpecialFlagReservedBit13 = IMPMSpecialFlags(1 << 13)
	// IMPMSpecialFlagReservedBit14 :: Reserved
	IMPMSpecialFlagReservedBit14 = IMPMSpecialFlags(1 << 14)
	// IMPMSpecialFlagReservedBit15 :: Reserved
	IMPMSpecialFlagReservedBit15 = IMPMSpecialFlags(1 << 15)
)

// IsMessageAttached returns true if there is a special message attached to the file
func (sf IMPMSpecialFlags) IsMessageAttached() bool {
	return (sf & IMPMSpecialFlagMessageAttached) != 0
}

// IsEmbedMidi returns true if embedded midi configuration is enabled
func (sf IMPMSpecialFlags) IsEmbedMidi() bool {
	return (sf & IMPMSpecialFlagEmbedMidi) != 0
}
