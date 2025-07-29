package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"

	"codegen/api"
	"codegen/model"
)

var functions = template.FuncMap{
	"titleCase":   titleCase,
	"hyphenate":   hyphenate,
	"hex":         hex,
	"args":        args,
	"arg":         arg,
	"fields2args": fields2args,
	"pack":        pack,
	"unpack":      unpack,
	"describe":    describe,
	"lookup":      lookup,
	"includes":    includes,
	"value":       value,
}

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]

		switch args[0] {
		case "codec":
			codec()

		case "integration-tests":
			integrationTests()

		case "API":
			api.API()
		}
	}
}

func titleCase(s string) string {
	re := regexp.MustCompile(`[ -]+`)
	parts := re.Split(s, -1)
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}

	return strings.Join(parts, "")
}

func hyphenate(s string) string {
	re := regexp.MustCompile(`[ -]+`)
	parts := re.Split(s, -1)
	for i := range parts {
		parts[i] = strings.ToLower(parts[i])
	}

	return strings.Join(parts, "-")
}

func hex(bytes []byte) string {
	lines := []string{}
	hex := "0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x,"

	for i := 0; i < 4; i++ {
		offset := i * 16
		u := bytes[offset : offset+8]
		v := bytes[offset+8 : offset+16]

		p := fmt.Sprintf(hex, u[0], u[1], u[2], u[3], u[4], u[5], u[6], u[7])
		q := fmt.Sprintf(hex, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7])

		lines = append(lines, fmt.Sprintf("            %v %v", p, q))
	}

	return strings.Join(lines, "\n")
}

func args(args []model.Arg) string {
	var parts []string
	for _, a := range args {
		parts = append(parts, arg(a))
	}

	return strings.Join(parts, ", ")
}

func arg(arg model.Arg) string {
	switch arg.Type {
	case "uint8":
		return fmt.Sprintf(`uint8(%v)`, arg.Value)

	case "uint32":
		return fmt.Sprintf(`uint32(%v)`, arg.Value)

	case "IPv4":
		return fmt.Sprintf(`netip.MustParseAddr("%v")`, arg.Value)

	case "addrport":
		return fmt.Sprintf(`netip.MustParseAddrPort("%v")`, arg.Value)

	case "datetime":
		return fmt.Sprintf(`string2datetime("%v")`, arg.Value)

	default:
		return fmt.Sprintf("%v", arg.Value)
	}
}

func fields2args(fields []model.Field) string {
	var args []string
	for _, f := range fields {
		switch f.Type {
		case "IPv4":
			args = append(args, fmt.Sprintf("%v netip.Addr", f.Name))

		case "addrport":
			args = append(args, fmt.Sprintf("%v netip.AddrPort", f.Name))

		case "datetime":
			args = append(args, fmt.Sprintf("%v time.Time", f.Name))

		case "magic":
			// skip

		default:
			args = append(args, fmt.Sprintf("%v %v", f.Name, f.Type))
		}
	}

	return strings.Join(args, ", ")
}

func pack(field model.Field) string {
	switch field.Type {
	case "uint8":
		return fmt.Sprintf("packUint8(%v, packet, %v)", field.Name, field.Offset)

	case "uint16":
		return fmt.Sprintf("packUint16(%v, packet, %v)", field.Name, field.Offset)

	case "uint32":
		return fmt.Sprintf("packUint32(%v, packet, %v)", field.Name, field.Offset)

	case "IPv4":
		return fmt.Sprintf("packIPv4(%v, packet, %v)", field.Name, field.Offset)

	case "addrport":
		return fmt.Sprintf("packAddrPort(%v, packet, %v)", field.Name, field.Offset)

	case "datetime":
		return fmt.Sprintf("packDateTime(%v, packet, %v)", field.Name, field.Offset)

	case "magic":
		return fmt.Sprintf("packUint32(0x55aaaa55, packet, %v)", field.Offset)

	default:
		panic(fmt.Sprintf("*** ERROR unsupported field type (%v)", field.Type))
	}
}

func unpack(field model.Field) string {
	switch field.Type {
	case "bool":
		return fmt.Sprintf("unpackBool(packet, %v)", field.Offset)

	case "uint8":
		return fmt.Sprintf("unpackUint8(packet, %v)", field.Offset)

	case "uint32":
		return fmt.Sprintf("unpackUint32(packet, %v)", field.Offset)

	case "datetime":
		return fmt.Sprintf("unpackYYYYMMDDHHMMSS(packet, %v)", field.Offset)

	case "date":
		return fmt.Sprintf("unpackYYYYMMDD(packet, %v)", field.Offset)

	case "shortdate":
		return fmt.Sprintf("unpackYYMMDD(packet, %v)", field.Offset)

	case "time":
		return fmt.Sprintf("unpackHHMMSS(packet, %v)", field.Offset)

	case "IPv4":
		return fmt.Sprintf("unpackIPv4(packet, %v)", field.Offset)

	case "addrport":
		return fmt.Sprintf("unpackAddrPort(packet, %v)", field.Offset)

	case "MAC":
		return fmt.Sprintf("unpackMAC(packet, %v)", field.Offset)

	case "version":
		return fmt.Sprintf("unpackVersion(packet, %v)", field.Offset)

	default:
		panic(fmt.Sprintf("*** ERROR unsupported field type (%v)", field.Type))
	}
}

func describe(field model.Field) string {
	return fmt.Sprintf("%v  (%v)  %v", field.Name, field.Type, field.Description)
}

var types = map[string]string{
	"uint8":      "uint8",
	"uint16":     "uint16",
	"uint32":     "uint32",
	"bool":       "bool",
	"IPv4":       "netip.Addr",
	"MAC":        "string",
	"version":    "string",
	"date":       "time.Time",
	"shortdate":  "Date",
	"time":       "Time",
	"datetime":   "DateTime",
	"HHmm":       "HHmm",
	"pin":        "PIN",
	"controller": "Controller",

	"optional date":     "Date",
	"optional datetime": "DateTime",
}

func lookup(path, key, defval string) any {
	table := types

	if v, ok := table[key]; ok {
		return v
	}

	return defval
}

func includes(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}

	return false
}

func value(v any, vtype string) string {
	switch vtype {
	case "IPv4":
		return fmt.Sprintf(`IPv4("%v")`, v)

	case "addrport":
		return fmt.Sprintf(`addrport("%v")`, v)

	case "MAC":
		return fmt.Sprintf(`"%v"`, v)

	case "version":
		return fmt.Sprintf(`"%v"`, v)

	case "datetime":
		return fmt.Sprintf(`string2datetime("%v")`, v)

	case "date":
		return fmt.Sprintf(`string2date("%v")`, v)

	case "time":
		return fmt.Sprintf(`string2time("%v")`, v)

	case "string":
		return fmt.Sprintf(`"%v"`, v)

	default:
		return fmt.Sprintf("%v", v)
	}
}
