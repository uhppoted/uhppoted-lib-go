package entities

type TimeProfile struct {
	Profile       uint8         `json:"profile"`
	StartDate     Date          `json:"start-date"`
	EndDate       Date          `json:"end-date"`
	Weekdays      Weekdays      `json:"weekdays"`
	Segments      []TimeSegment `json:"segments"`
	LinkedProfile uint8         `json:"linked-profile"`
}

type Weekdays struct {
	Monday    bool `json:"monday"`
	Tuesday   bool `json:"tuesday"`
	Wednesday bool `json:"wednesday"`
	Thursday  bool `json:"thursday"`
	Friday    bool `json:"friday"`
	Saturday  bool `json:"saturday"`
	Sunday    bool `json:"sunday"`
}

type TimeSegment struct {
	Start HHmm `json:"start"`
	End   HHmm `json:"end"`
}
