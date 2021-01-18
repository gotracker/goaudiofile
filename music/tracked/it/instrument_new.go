package it

// IMPIInstrument is the format of the IMPI Instrument for tracker compatibility versions >= 0x0200
type IMPIInstrument struct {
	IMPI                   [4]byte
	Filename               [12]byte
	Nul10                  uint8
	NewNoteAction          NewNoteAction
	DuplicateCheckType     DuplicateCheckType
	DuplicateCheckAction   DuplicateCheckAction
	Fadeout                uint16
	PitchPanSeparation     int8
	PitchPanCenter         uint8
	GlobalVolume           FineVolume
	DefaultPan             PanValue
	RandomVolumeVariation  Percentage8
	RandomPanVariation     Percentage8
	TrackerVersion         uint16
	SampleCount            uint8
	Reserved1F             uint8
	Name                   [26]byte
	InitialFilterCutoff    uint8
	InitialFilterResonance uint8
	MidiChannel            uint8
	MidiProgram            uint8
	MidiBank               uint16
	NoteSampleKeyboard     [240]uint8
	VolumeEnvelope         Envelope
	PanningEnvelope        Envelope
	PitchEnvelope          Envelope
}
