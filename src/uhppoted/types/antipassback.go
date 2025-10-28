package types

type AntiPassback uint8

const (
	// disabled
	NoAntiPassback AntiPassback = 0x00

	// readers 1:2; 3:4 (independently)
	Readers12_34 AntiPassback = 0x01

	// readers (1,3):(2,4)
	Readers13_24 AntiPassback = 0x02

	// readers 1:(2,3)
	Readers1_23 AntiPassback = 0x03

	// readers 1:(2,3,4)
	Readers1_234 AntiPassback = 0x04
)

func (a AntiPassback) String() string {
	switch a {
	case NoAntiPassback:
		return "disabled"
	case Readers12_34:
		return "readers 1:2; 3:4 (independently)"
	case Readers13_24:
		return "readers (1,3):(2,4)"
	case Readers1_23:
		return "readers 1:(2,3)"
	case Readers1_234:
		return "readers 1:(2,3,4)3"
	default:
		return "unknown"
	}
}
