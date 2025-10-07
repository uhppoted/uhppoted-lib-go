package entities

type TimeProfile struct {
	Profile       uint8         `json:"profile"`
	StartDate     Date          `json:"start-date"`
	EndDate       Date          `json:"end-date"`
	Weekdays      Weekdays      `json:"weekdays"`
	Segments      []TimeSegment `json:"segments"`
	LinkedProfile uint8         `json:"linked-profile"`
}

type TimeSegment struct {
	Start HHmm `json:"start"`
	End   HHmm `json:"end"`
}
