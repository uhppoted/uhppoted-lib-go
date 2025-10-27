package types

type DoorMode uint8

const (
	NormallyOpen   DoorMode = 0x01
	NormallyClosed DoorMode = 0x02
	Controlled     DoorMode = 0x03
)

func (m DoorMode) String() string {
	switch m {
	case NormallyOpen:
		return "normally open"
	case NormallyClosed:
		return "normally closed"
	case Controlled:
		return "controlled"
	default:
		return "unknown"
	}
}
