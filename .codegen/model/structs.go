package model

import (
	lib "github.com/uhppoted/uhppoted-codegen/model"
)

type Request struct {
	lib.Message
	Tests []RequestTest
}

type Arg struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value any    `json:"value,omitempty"`
}

type Field lib.Field

type Value struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value any    `json:"value"`
}

type RequestTest struct {
	Name     string `json:"name"`
	Args     []Arg  `json:"args"`
	Expected []byte `json:"expected,omitempty"`
}

type Test struct {
	Name     string `json:"name"`
	Args     []Arg
	Expected []byte
	Response []byte  `json:"packet,omitempty"`
	Values   []Value `json:"values,omitempty"`
}
