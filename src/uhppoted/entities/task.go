package entities

type TaskType uint8

const (
	ControlDoor          TaskType = 0
	UnlockDoor           TaskType = 1
	LockDoor             TaskType = 2
	DisableTimeProfiles  TaskType = 3
	EnableTimeProfiles   TaskType = 4
	EnableCardNoPassword TaskType = 5
	EnableCardInPassword TaskType = 6
	EnableCardPassword   TaskType = 7
	EnableMoreCards      TaskType = 8
	DisableMoreCards     TaskType = 9
	TriggerOnce          TaskType = 10
	DisablePushbutton    TaskType = 11
	EnablePushbutton     TaskType = 12
)

func (t TaskType) String() string {
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

type Task struct {
	Task      TaskType `json:"task"`
	Door      uint8    `json:"door"`
	StartDate Date     `json:"start-date"`
	EndDate   Date     `json:"end-date"`
	StartTime HHmm     `json:"start-time"`
	Weekdays  Weekdays `json:"weekdays"`
	MoreCards uint8    `json:"more-cards"`
}
