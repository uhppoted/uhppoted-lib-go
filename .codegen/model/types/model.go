package types

import (
	lib "github.com/uhppoted/uhppoted-codegen/model/types"
)

type Request lib.Request
type Response lib.Response
type RequestTest lib.RequestTest
type ResponseTest lib.ResponseTest

type Arg lib.Arg
type Value lib.Value
type Field lib.Field

type FuncTest lib.FuncTest
type TestArg lib.TestArg
type TestReply lib.TestReply

type Function struct {
	Name        string
	Description string
	Request     Request
	Response    Response
	Protocols   []string
	Tests       []FuncTest
}
