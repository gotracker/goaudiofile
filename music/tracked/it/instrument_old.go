package it

// IMPIInstrumentOld is the format of the IMPI Instrument for tracker compatibility versions < 0x0200
type IMPIInstrumentOld struct {
	IMPI               [4]byte
	Filename           [12]byte
	Nul10              uint8
	Flags              IMPIOldFlags
	VolumeLoopStart    uint8
	VolumeLoopEnd      uint8
	SustainLoopStart   uint8
	SustainLoopEnd     uint8
	Fadeout            uint16
	NewNoteAction      NewNoteAction
	DuplicateNoteCheck DuplicateNoteCheck
	TrackerVersion     uint16
	SampleCount        uint8
	Reserved1F         uint8
	Name               [26]byte
	Reserved3A         [6]uint8
	NoteSampleKeyboard [240]uint8
	VolumeEnvelope     [200]uint8
	NodePoints         [25]NodePoint16
}

// IMPIOldFlags is the flagset for IMPIInstrumentOld instruments
type IMPIOldFlags uint8

const (
	// IMPIOldFlagUseVolumeEnvelope :: On = Use volume envelope
	IMPIOldFlagUseVolumeEnvelope = IMPIOldFlags(1 << 0)
	// IMPIOldFlagUseVolumeLoop :: On = Use volume loop
	IMPIOldFlagUseVolumeLoop = IMPIOldFlags(1 << 1)
	// IMPIOldFlagUseSustainVolumeLoop :: On = Use sustain volume loop
	IMPIOldFlagUseSustainVolumeLoop = IMPIOldFlags(1 << 2)
)

// DuplicateNoteCheck activates or deactivates the duplicate note checking
type DuplicateNoteCheck uint8

const (
	// DuplicateNoteCheckOff disables the duplicate note checking
	DuplicateNoteCheckOff = DuplicateNoteCheck(0)
	// DuplicateNoteCheckOn activates the duplicate note checking
	DuplicateNoteCheckOn = DuplicateNoteCheck(1)
)
