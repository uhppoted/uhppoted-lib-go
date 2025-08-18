package types

import (
	lib "github.com/uhppoted/uhppoted-codegen/model/types"
)

type Function lib.Function
type Request lib.Request
type Response lib.Response

type Arg lib.Arg
type Value lib.Value
type Field lib.Field

type RequestTest lib.RequestTest
type ResponseTest lib.ResponseTest
type FuncTest lib.FuncTest

type TestArg lib.TestArg
type TestReply lib.TestReply

// type Function struct {
// 	Name        string
// 	Description string
// 	Request     lib.Message
// 	Response    lib.Message
// 	Protocols   []string
// 	Tests       []lib.FuncTest
// }
