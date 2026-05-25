package types

import (
	"fmt"
	"strings"
)

type FirstCard struct {
	StartTime    HHmm     `json:"start-time"`
	EndTime      HHmm     `json:"end-time"`
	ActiveMode   DoorMode `json:"active-mode"`
	InactiveMode DoorMode `json:"inactive-mode"`
	Weekdays     Weekdays `json:"weekdays"`
}

func (f FirstCard) String() string {
	start := fmt.Sprintf("%v", f.StartTime)
	end := fmt.Sprintf("%v", f.EndTime)
	times := ""

	if start != "" && end != "" {
		times = start + "-" + end
	} else if start != "" {
		times = start + "-"
	} else {
		times = "-" + end
	}

	active := fmt.Sprintf("%4v", f.ActiveMode)
	inactive := fmt.Sprintf("%4v", f.InactiveMode)
	weekdays := fmt.Sprintf("%v", f.Weekdays)

	list := []string{}
	for _, s := range []string{"active:" + times, "control:" + active + "/" + inactive, "weekdays:" + weekdays} {
		if s != "" {
			list = append(list, s)
		}
	}

	return strings.Join(list, "  ")
}
