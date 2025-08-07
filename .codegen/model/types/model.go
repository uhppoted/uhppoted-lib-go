package types

import (
	lib "github.com/uhppoted/uhppoted-codegen/model/types"
)

type Request lib.Request
type Response lib.Response
type RequestTest lib.RequestTest
type ResponseTest lib.ResponseTest

type Field lib.Field

// type Response struct {
// 	Message lib.Message
// 	Tests   []ResponseTest
// }

// type ResponseTest struct {
// 	Name     string  `json:"name"`
// 	Response []byte  `json:"response"`
// 	Expected []Value `json:"expected"`
// }

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
