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
	Test        FuncTest
}

type Reply struct {
	Message  []byte
	Response []types.Value
}

type FuncTest struct {
	Args    []types.Arg
	Request []byte
	Replies []Reply
}
