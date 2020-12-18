package mod

// ModuleHeader is a representation of the MOD file header
type ModuleHeader struct {
	Name       [20]byte
	Samples    [31]InstrumentHeader
	SongLen    uint8
	RestartPos uint8
	Order      [128]uint8
	Sig        [4]uint8
}
