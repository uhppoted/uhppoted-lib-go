package entities

type Task uint8

const (
	ControlDoor          Task = 0
	UnlockDoor           Task = 1
	LockDoor             Task = 2
	DisableTimeProfiles  Task = 3
	EnableTimeProfiles   Task = 4
	EnableCardNoPassword Task = 5
	EnableCardInPassword Task = 6
	EnableCardPassword   Task = 7
	EnableMoreCards      Task = 8
	DisableMoreCards     Task = 9
	TriggerOnce          Task = 10
	DisablePushbutton    Task = 11
	EnablePushbutton     Task = 12
)

func (t Task) String() string {
	switch t {
	case ControlDoor:
		return "control door"

	case UnlockDoor:
		return "unlock door"

	case LockDoor:
		return "lock door"

	case DisableTimeProfiles:
		return "disable time profiles"

	case EnableTimeProfiles:
		return "enable time profiles"

	case EnableCardNoPassword:
		return "enable card (no PIN)"

	case EnableCardInPassword:
		return "enable card IN PIN"

	case EnableCardPassword:
		return "enable card PIN"

	case EnableMoreCards:
		return "enable more cards"

	case DisableMoreCards:
		return "disable more cards"

	case TriggerOnce:
		return "trigger once"

	case DisablePushbutton:
		return "disable pushbutton"

	case EnablePushbutton:
		return "enable pushbutton"

	default:
		return "unknown"
	}
}
