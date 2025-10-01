package entities

type Event struct {
	Index         uint32   `json:"index"`
	Event         uint8    `json:"event"`
	AccessGranted bool     `json:"granted"`
	Door          uint8    `json:"door"`
	Direction     uint8    `json:"direction"`
	Card          uint32   `json:"card"`
	Timestamp     DateTime `json:"timestamp"`
	Reason        uint8    `json:"reason"`
}
