package s3m

// ChannelID is the S3M value for a channel specification (found within the ChanenlSetting header block)
type ChannelID uint8

const (
	// ChannelIDL1 is the Left Channel 1
	ChannelIDL1 = ChannelID(0 + iota)
	// ChannelIDL2 is the Left Channel 2
	ChannelIDL2
	// ChannelIDL3 is the Left Channel 3
	ChannelIDL3
	// ChannelIDL4 is the Left Channel 4
	ChannelIDL4
	// ChannelIDL5 is the Left Channel 5
	ChannelIDL5
	// ChannelIDL6 is the Left Channel 6
	ChannelIDL6
	// ChannelIDL7 is the Left Channel 7
	ChannelIDL7
	// ChannelIDL8 is the Left Channel 8
	ChannelIDL8
	// ChannelIDR1 is the Right Channel 1
	ChannelIDR1
	// ChannelIDR2 is the Right Channel 2
	ChannelIDR2
	// ChannelIDR3 is the Right Channel 3
	ChannelIDR3
	// ChannelIDR4 is the Right Channel 4
	ChannelIDR4
	// ChannelIDR5 is the Right Channel 5
	ChannelIDR5
	// ChannelIDR6 is the Right Channel 6
	ChannelIDR6
	// ChannelIDR7 is the Right Channel 7
	ChannelIDR7
	// ChannelIDR8 is the Right Channel 8
	ChannelIDR8
	// ChannelIDOPL2Melody1 is the Adlib/OPL2 Melody Channel 1
	ChannelIDOPL2Melody1
	// ChannelIDOPL2Melody2 is the Adlib/OPL2 Melody Channel 2
	ChannelIDOPL2Melody2
	// ChannelIDOPL2Melody3 is the Adlib/OPL2 Melody Channel 3
	ChannelIDOPL2Melody3
	// ChannelIDOPL2Melody4 is the Adlib/OPL2 Melody Channel 4
	ChannelIDOPL2Melody4
	// ChannelIDOPL2Melody5 is the Adlib/OPL2 Melody Channel 5
	ChannelIDOPL2Melody5
	// ChannelIDOPL2Melody6 is the Adlib/OPL2 Melody Channel 6
	ChannelIDOPL2Melody6
	// ChannelIDOPL2Melody7 is the Adlib/OPL2 Melody Channel 7
	ChannelIDOPL2Melody7
	// ChannelIDOPL2Melody8 is the Adlib/OPL2 Melody Channel 8
	ChannelIDOPL2Melody8
	// ChannelIDOPL2Drums1 is the Adlib/OPL2 Drums Channel 1
	ChannelIDOPL2Drums1
	// ChannelIDOPL2Drums2 is the Adlib/OPL2 Drums Channel 2
	ChannelIDOPL2Drums2
	// ChannelIDOPL2Drums3 is the Adlib/OPL2 Drums Channel 3
	ChannelIDOPL2Drums3
	// ChannelIDOPL2Drums4 is the Adlib/OPL2 Drums Channel 4
	ChannelIDOPL2Drums4
	// ChannelIDOPL2Drums5 is the Adlib/OPL2 Drums Channel 5
	ChannelIDOPL2Drums5
	// ChannelIDOPL2Drums6 is the Adlib/OPL2 Drums Channel 6
	ChannelIDOPL2Drums6
	// ChannelIDOPL2Drums7 is the Adlib/OPL2 Drums Channel 7
	ChannelIDOPL2Drums7
	// ChannelIDOPL2Drums8 is the Adlib/OPL2 Drums Channel 8
	ChannelIDOPL2Drums8
)

// ChannelCategory is the type of channel
type ChannelCategory int

const (
	// ChannelCategoryUnknown is an unknown channel category
	ChannelCategoryUnknown = ChannelCategory(iota)
	// ChannelCategoryPCMLeft is a PCM audio left channel
	ChannelCategoryPCMLeft
	// ChannelCategoryPCMRight is a PCM audio right channel
	ChannelCategoryPCMRight
	// ChannelCategoryOPL2Melody is an OPL2 audio melody channel
	ChannelCategoryOPL2Melody
	// ChannelCategoryOPL2Drums is an OPL2 audio drums channel
	ChannelCategoryOPL2Drums
)

// GetChannelCategory returns the channel category for a particular channel ID
func (c ChannelID) GetChannelCategory() ChannelCategory {
	switch {
	case c >= ChannelIDL1 && c <= ChannelIDL8:
		return ChannelCategoryPCMLeft
	case c >= ChannelIDR1 && c <= ChannelIDR8:
		return ChannelCategoryPCMRight
	case c >= ChannelIDOPL2Melody1 && c <= ChannelIDOPL2Melody8:
		return ChannelCategoryOPL2Melody
	case c >= ChannelIDOPL2Drums1 && c <= ChannelIDOPL2Drums8:
		return ChannelCategoryOPL2Drums
	default:
		return ChannelCategoryUnknown
	}
}

// ChannelSetting is a full channel setting (with flags) definition from the S3M header
type ChannelSetting uint8

const (
	// ChannelSettingDisabled is the flag signifying that the channel is deactivated
	ChannelSettingDisabled = ChannelSetting(0x80)
)

// IsEnabled returns the enabled flag (bit 7 is unset)
func (cs ChannelSetting) IsEnabled() bool {
	return (cs & ChannelSettingDisabled) == 0
}

// GetChannel returns the ChannelID for the channel
func (cs ChannelSetting) GetChannel() ChannelID {
	ch := uint8(cs) & 0x7F
	return ChannelID(ch)
}

// IsPCM returns true if the channel is one of the left or right channels (non-Adlib/OPL2)
func (cs ChannelSetting) IsPCM() bool {
	ch := cs.GetChannel()
	cc := ch.GetChannelCategory()
	return (cc == ChannelCategoryPCMLeft || cc == ChannelCategoryPCMRight)
}

// IsOPL2 returns true if the channel is an Adlib/OPL2 channel (non-PCM)
func (cs ChannelSetting) IsOPL2() bool {
	ch := cs.GetChannel()
	cc := ch.GetChannelCategory()
	return (cc == ChannelCategoryOPL2Melody || cc == ChannelCategoryOPL2Drums)
}

// MakeChannelSetting returns a channel setting value based on component inputs
func MakeChannelSetting(enabled bool, cat ChannelCategory, idx int) ChannelSetting {
	cs := ChannelSetting(0)
	if !enabled {
		cs = cs | ChannelSettingDisabled
	}
	switch cat {
	case ChannelCategoryPCMLeft:
		cs = cs | ChannelSetting(ChannelIDL1)
		cs |= ChannelSetting(idx & 0x07)
	case ChannelCategoryPCMRight:
		cs = cs | ChannelSetting(ChannelIDR1)
		cs |= ChannelSetting(idx & 0x07)
	case ChannelCategoryOPL2Melody:
		cs = cs | ChannelSetting(ChannelIDOPL2Melody1)
		cs |= ChannelSetting(idx & 0x07)
	case ChannelCategoryOPL2Drums:
		cs = cs | ChannelSetting(ChannelIDOPL2Drums1)
		cs |= ChannelSetting(idx & 0x07)
	}
	return cs
}
