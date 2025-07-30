package model

import (
	lib "github.com/uhppoted/uhppoted-codegen/model"
)

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

type Test struct {
	Name     string `json:"name"`
	Args     []Arg
	Expected []byte
	Response []byte  `json:"packet,omitempty"`
	Values   []Value `json:"values,omitempty"`
}
