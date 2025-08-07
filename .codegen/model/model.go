package model

import (
	"codegen/model/types"
)

type Func struct {
	Name        string
	Description string
	Request     types.Request
	Response    types.Response
	Protocols   []string
	Tests       []FuncTest
}

type Reply struct {
	Message  []byte
	Response []types.Value
}

type FuncTest struct {
	Name    string
	Args    []types.Arg
	Request []byte
	Replies []Reply
}
