package integration_tests

import (
	"fmt"

	"go/token"

	"github.com/dave/dst"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"

	"codegen/codegen"
)

func buildAPITestExpected(fn lib.Function, test lib.FuncTest) dst.Stmt {
	name := codegen.TitleCase(test.Name)

	return &dst.AssignStmt{
		Lhs: []dst.Expr{dst.NewIdent("expected")},
		Tok: token.DEFINE,
		Rhs: []dst.Expr{
			&dst.SelectorExpr{
				X: &dst.SelectorExpr{
					X:   &dst.Ident{Name: "test"},
					Sel: &dst.Ident{Name: "Expected"},
				},

				Sel: &dst.Ident{Name: name},
			},
		},

		Decs: dst.AssignStmtDecorations{
			NodeDecs: dst.NodeDecs{
				After: dst.EmptyLine,
			},
		},
	}
}

func buildAPITestArg(arg lib.Arg) dst.Expr {
	switch arg.Type {
	case "bool":
		return &dst.BasicLit{
			Kind:  token.IDENT,
			Value: fmt.Sprintf(`%v`, arg.Value),
		}

	case "uint8":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "uint8"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`%v`, arg.Value),
				},
			},
		}

	case "uint16":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "uint16"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`%v`, arg.Value),
				},
			},
		}

	case "uint32":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "uint32"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`%v`, arg.Value),
				},
			},
		}

	case "IPv4":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "netip.MustParseAddr"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, arg.Value),
				},
			},
		}

	case "address:port":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "netip.MustParseAddrPort"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, arg.Value),
				},
			},
		}

	case "datetime":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "types.MustParseDateTime"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, arg.Value),
				},
			},
		}

	case "date":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "types.MustParseDate"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, arg.Value),
				},
			},
		}

	case "HHmm":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "types.MustParseHHmm"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`"%v"`, arg.Value),
				},
			},
		}

	case "pin":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "uint32"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: fmt.Sprintf(`%v`, arg.Value),
				},
			},
		}

	case "mode":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "types.DoorMode"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.INT,
					Value: fmt.Sprintf(`%v`, arg.Value),
				},
			},
		}

	case "task":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "types.TaskType"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.INT,
					Value: fmt.Sprintf(`%v`, arg.Value),
				},
			},
		}

	case "interlock":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "types.Interlock"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.INT,
					Value: fmt.Sprintf(`%v`, arg.Value),
				},
			},
		}

	case "anti-passback":
		return &dst.CallExpr{
			Fun: &dst.Ident{Name: "types.AntiPassback"},
			Args: []dst.Expr{
				&dst.BasicLit{
					Kind:  token.INT,
					Value: fmt.Sprintf(`%v`, arg.Value),
				},
			},
		}

	default:
		panic(fmt.Sprintf("unknown arg type '%v'", arg.Type))
	}
}

func buildAPITestExec(fn lib.Function, test lib.FuncTest) dst.Stmt {
	name := codegen.TitleCase(fn.Name)

	args := []dst.Expr{
		&dst.Ident{
			Name: "u",
		},
	}

	for _, arg := range test.Args {
		args = append(args, &dst.Ident{
			Name: codegen.CamelCase(arg.Name),
		})
	}

	args = append(args, &dst.Ident{
		Name: "timeout",
	})

	return &dst.AssignStmt{
		Lhs: []dst.Expr{
			dst.NewIdent("response"),
			dst.NewIdent("err"),
		},
		Tok: token.DEFINE,
		Rhs: []dst.Expr{
			&dst.CallExpr{
				Fun: &dst.SelectorExpr{
					X:   dst.NewIdent("uhppoted"),
					Sel: dst.NewIdent(name),
				},
				Args: args,
			},
		},

		Decs: dst.AssignStmtDecorations{
			NodeDecs: dst.NodeDecs{
				Before: dst.EmptyLine,
				After:  dst.EmptyLine,
			},
		},
	}
}

func buildAPITestValidate(fn lib.Function, test lib.FuncTest) dst.Stmt {
	return &dst.IfStmt{
		Cond: &dst.BinaryExpr{
			X:  &dst.Ident{Name: "err"},
			Op: token.NEQ,
			Y:  &dst.Ident{Name: "nil"},
		},
		Body: &dst.BlockStmt{
			List: []dst.Stmt{
				&dst.ExprStmt{
					X: &dst.CallExpr{
						Fun: &dst.SelectorExpr{
							X:   &dst.Ident{Name: "t"},
							Sel: &dst.Ident{Name: "Fatalf"},
						},
						Args: []dst.Expr{
							&dst.BasicLit{Kind: token.STRING, Value: `"%v"`},
							&dst.Ident{Name: "err"},
						},
					},
				},
			},
		},
		Else: &dst.IfStmt{
			Cond: &dst.UnaryExpr{
				Op: token.NOT,
				X: &dst.CallExpr{
					Fun: &dst.SelectorExpr{
						X:   &dst.Ident{Name: "reflect"},
						Sel: &dst.Ident{Name: "DeepEqual"},
					},
					Args: []dst.Expr{
						&dst.Ident{Name: "response"},
						&dst.Ident{Name: "expected"},
					},
				},
			},
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ExprStmt{
						X: &dst.CallExpr{
							Fun: &dst.SelectorExpr{
								X:   &dst.Ident{Name: "t"},
								Sel: &dst.Ident{Name: "Errorf"},
							},
							Args: []dst.Expr{
								&dst.BasicLit{
									Kind:  token.STRING,
									Value: `"incorrect response\n   expected:%#v\n   got:     %#v"`,
								},
								&dst.Ident{Name: "expected"},
								&dst.Ident{Name: "response"},
							},
						},
					},
				},
			},
		},
	}
}
