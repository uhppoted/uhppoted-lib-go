package model

type Arg struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Field struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Offset      uint8  `json:"offset"`
	Description string `json:"description"`
}

type Test struct {
	Name     string
	Args     []any
	Expected []byte
}
