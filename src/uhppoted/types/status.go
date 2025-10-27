package types

type Status struct {
	System struct {
		Time  DateTime `json:"datetime"`
		Error uint8    `json:"error"`
		Info  uint8    `json:"info"`
	} `json:"system"`

	Doors map[uint8]struct {
		Open     bool `json:"open"`
		Button   bool `json:"button"`
		Unlocked bool `json:"unlocked"`
	} `json:"doors"`

	Alarms struct {
		Flags      uint8 `json:"flags"`
		Fire       bool  `json:"fire"`
		LockForced bool  `json:"lock-forced"`
	} `json:"alarms"`

	Event Event `json:"event"`
}
