package model

import (
	types "github.com/uhppoted/uhppoted-codegen/model/types"
)

type Func struct {
	Name        string
	Description string
	Request     types.Request
	Response    Response
	Protocols   []string
	Test        FuncTest
}

type Reply struct {
	Message  []byte
	Response []Value
}

type FuncTest struct {
	Args    []Arg
	Request []byte
	Replies []Reply
}

type Response struct {
	Name    string
	MsgType byte
	Fields  []types.Field
	Tests   []Test
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

type Test struct {
	Name     string `json:"name"`
	Args     []Arg
	Expected []byte
	Response []byte  `json:"packet,omitempty"`
	Values   []Value `json:"values,omitempty"`
}
