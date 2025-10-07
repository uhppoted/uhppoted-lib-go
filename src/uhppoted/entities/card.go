package entities

type Card struct {
	Card        uint32          `json:"card"`
	StartDate   Date            `json:"start-date"`
	EndDate     Date            `json:"end-date"`
	Permissions map[uint8]uint8 `json:"permissions"`
	PIN         uint32          `json:"PIN"`
}
