package it

import "time"

type EditHistory struct {
	Date    EditDate
	Time    EditTime
	Runtime Runtime
}

type EditDate uint16

func (d EditDate) Components() (int, time.Month, int) {
	year := int(uint16(d)>>9)&0x7F + 1980
	month := int(uint16(d)>>5) & 0xF
	day := int(uint16(d)) & 0x1F
	return year, time.Month(month), day
}

type EditTime uint16

func (t EditTime) Components() (int, int, int) {
	hour := int(uint16(t)>>11) & 0x1F
	minute := int(uint16(t)>>5) & 0x3F
	second := int(uint16(t)<<1) & 0x3F

	return hour, minute, second
}

type Runtime uint32

const (
	dosTimer = 100 * time.Second / 1820
)

func (r Runtime) Duration() time.Duration {
	return time.Duration(r) * dosTimer
}
