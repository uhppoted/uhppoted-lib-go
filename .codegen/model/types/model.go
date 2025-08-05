package types

import (
	lib "github.com/uhppoted/uhppoted-codegen/model/types"
)

type Request lib.Request
type RequestTest lib.RequestTest
type Field lib.Field

type Response struct {
	Name    string
	MsgType byte
	Fields  []Field
	Tests   []Test
}

type Test struct {
	Name     string `json:"name"`
	Args     []Arg
	Expected []byte
	Response []byte  `json:"packet,omitempty"`
	Values   []Value `json:"values,omitempty"`
}

type Arg struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value any    `json:"value,omitempty"`
}

type Value struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value any    `json:"value"`
}
