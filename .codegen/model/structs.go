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

type Value struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value any    `json:"value"`
}

type Test struct {
	Name     string `json:"name"`
	Args     []any
	Expected []byte
	Response []byte  `json:"packet,omitempty"`
	Values   []Value `json:"values,omitempty"`
}
