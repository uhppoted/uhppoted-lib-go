package codegen

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
	"text/template"
	"unicode"

	lib "github.com/uhppoted/uhppoted-codegen/model/types"
)

var Functions = template.FuncMap{
	"titleCase":   TitleCase,
	"camelCase":   camelCase,
	"hyphenate":   hyphenate,
	"clean":       clean,
	"hex":         hex,
	"testargs":    testargs,
	"testarg":     testarg,
	"fields2args": fields2args,
	"pack":        pack,
	"unpack":      unpack,
	"describe":    describe,
	"lookup":      lookup,
	"includes":    includes,
	"value":       value,
	"rpad":        rpad,
	"article":     article,
}

func TitleCase(s string) string {
	re := regexp.MustCompile(`[ \-:]+`)
	parts := re.Split(s, -1)
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}

	return strings.Join(parts, "")
}

func camelCase(s string) string {
	tokens := regexp.MustCompile(`[ \-:]+`).Split(s, -1)

	for i, token := range tokens[1:] {
		tokens[i+1] = capitalize(token)
	}

	return strings.Join(tokens, "")
}

func capitalize(s string) string {
	runes := []rune(s)

	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}

	return string(runes)
}

func hyphenate(s string) string {
	re := regexp.MustCompile(`[ -]+`)
	parts := re.Split(s, -1)
	for i := range parts {
		parts[i] = strings.ToLower(parts[i])
	}

	return strings.Join(parts, "-")
}

func clean(s string) string {
	re := regexp.MustCompile(`[ :\-]+`)

	return strings.ToLower(re.ReplaceAllString(s, ""))
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

func testargs(args []lib.TestArg) string {
	var parts []string
	for _, a := range args {
		parts = append(parts, testarg(a))
	}

	return strings.Join(parts, ", ")
}

func testarg(arg lib.TestArg) string {
	switch arg.Type {
	case "uint8":
		return fmt.Sprintf(`uint8(%v)`, arg.Value)

	case "uint16":
		return fmt.Sprintf(`uint16(%v)`, arg.Value)

	case "uint32":
		return fmt.Sprintf(`uint32(%v)`, arg.Value)

	case "IPv4":
		return fmt.Sprintf(`netip.MustParseAddr("%v")`, arg.Value)

	case "addrport":
		return fmt.Sprintf(`netip.MustParseAddrPort("%v")`, arg.Value)

	case "address:port":
		return fmt.Sprintf(`netip.MustParseAddrPort("%v")`, arg.Value)

	case "datetime":
		return fmt.Sprintf(`string2datetime("%v")`, arg.Value)

	case "date":
		return fmt.Sprintf(`string2date("%v")`, arg.Value)

	case "HHmm":
		return fmt.Sprintf(`string2HHmm("%v")`, arg.Value)

	case "pin":
		return fmt.Sprintf(`uint32(%v)`, arg.Value)

	default:
		return fmt.Sprintf("%v", arg.Value)
	}
}

func fields2args(fields []lib.Field) string {
	var args []string
	for _, f := range fields {
		name := regexp.MustCompile(`\s+`).ReplaceAllString(f.Name, "")

		switch f.Type {
		case "IPv4":
			args = append(args, fmt.Sprintf("%v netip.Addr", name))

		case "address:port":
			args = append(args, fmt.Sprintf("%v netip.AddrPort", name))

		case "datetime":
			args = append(args, fmt.Sprintf("%v time.Time", name))

		case "date":
			args = append(args, fmt.Sprintf("%v time.Time", name))

		case "HHmm":
			args = append(args, fmt.Sprintf("%v time.Time", name))

		case "pin":
			args = append(args, fmt.Sprintf("%v uint32", name))

		case "magic":
			// skip

		default:
			args = append(args, fmt.Sprintf("%v %v", name, f.Type))
		}
	}

	return strings.Join(args, ", ")
}

func pack(field lib.Field) string {
	name := regexp.MustCompile(`\s+`).ReplaceAllString(field.Name, "")

	switch field.Type {
	case "bool":
		return fmt.Sprintf("packBool(%v, packet, %v)", name, field.Offset)

	case "uint8":
		return fmt.Sprintf("packUint8(%v, packet, %v)", name, field.Offset)

	case "uint16":
		return fmt.Sprintf("packUint16(%v, packet, %v)", name, field.Offset)

	case "uint32":
		return fmt.Sprintf("packUint32(%v, packet, %v)", name, field.Offset)

	case "IPv4":
		return fmt.Sprintf("packIPv4(%v, packet, %v)", name, field.Offset)

	case "address:port":
		return fmt.Sprintf("packAddrPort(%v, packet, %v)", name, field.Offset)

	case "datetime":
		return fmt.Sprintf("packDateTime(%v, packet, %v)", name, field.Offset)

	case "date":
		return fmt.Sprintf("packDate(%v, packet, %v)", name, field.Offset)

	case "HHmm":
		return fmt.Sprintf("packHHmm(%v, packet, %v)", name, field.Offset)

	case "pin":
		return fmt.Sprintf("packPIN(%v, packet, %v)", name, field.Offset)

	case "passcode":
		return fmt.Sprintf("packPasscode(%v, packet, %v)", name, field.Offset)

	case "magic":
		return fmt.Sprintf("packUint32(0x55aaaa55, packet, %v)", field.Offset)

	default:
		panic(fmt.Sprintf("*** ERROR unsupported field type (%v)", field.Type))
	}
}

func unpack(field lib.Field) string {
	switch field.Type {
	case "bool":
		return fmt.Sprintf("unpackBool(packet, %v)", field.Offset)

	case "uint8":
		return fmt.Sprintf("unpackUint8(packet, %v)", field.Offset)

	case "uint16":
		return fmt.Sprintf("unpackUint16(packet, %v)", field.Offset)

	case "uint32":
		return fmt.Sprintf("unpackUint32(packet, %v)", field.Offset)

	case "datetime":
		return fmt.Sprintf("unpackDateTime(packet, %v)", field.Offset)

	case "optional datetime":
		return fmt.Sprintf("unpackOptionalDateTime(packet, %v)", field.Offset)

	case "date":
		return fmt.Sprintf("unpackDate(packet, %v)", field.Offset)

	case "shortdate":
		return fmt.Sprintf("unpackShortDate(packet, %v)", field.Offset)

	case "optional date":
		return fmt.Sprintf("unpackOptionalDate(packet, %v)", field.Offset)

	case "time":
		return fmt.Sprintf("unpackHHMMSS(packet, %v)", field.Offset)

	case "HHmm":
		return fmt.Sprintf("unpackHHMM(packet, %v)", field.Offset)

	case "IPv4":
		return fmt.Sprintf("unpackIPv4(packet, %v)", field.Offset)

	case "address:port":
		return fmt.Sprintf("unpackAddrPort(packet, %v)", field.Offset)

	case "addrport":
		return fmt.Sprintf("unpackAddrPort(packet, %v)", field.Offset)

	case "MAC":
		return fmt.Sprintf("unpackMAC(packet, %v)", field.Offset)

	case "version":
		return fmt.Sprintf("unpackVersion(packet, %v)", field.Offset)

	case "pin":
		return fmt.Sprintf("unpackPIN(packet, %v)", field.Offset)

	default:
		panic(fmt.Sprintf("*** ERROR unsupported field type (%v)", field.Type))
	}
}

func describe(field lib.Field) string {
	return fmt.Sprintf("%v  (%v)  %v", field.Name, field.Type, field.Description)
}

func lookup(path, key, defval string) any {
	table := map[string]string{
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
		"HHmm":       "time.Time",
		"pin":        "PIN",
		"controller": "Controller",

		"optional date":     "Date",
		"optional datetime": "DateTime",
	}

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

	case "address:port":
		return fmt.Sprintf(`addrport("%v")`, v)

	case "MAC":
		return fmt.Sprintf(`"%v"`, v)

	case "version":
		return fmt.Sprintf(`"%v"`, v)

	case "datetime":
		return fmt.Sprintf(`string2datetime("%v")`, v)

	case "date":
		return fmt.Sprintf(`entities.MustParseDate("%v")`, v)

	case "shortdate":
		return fmt.Sprintf(`string2date("%v")`, v)

	case "optional date":
		return fmt.Sprintf(`string2date("%v")`, v)

	case "time":
		return fmt.Sprintf(`string2time("%v")`, v)

	case "HHmm":
		return fmt.Sprintf(`string2HHmm("%v")`, v)

	case "string":
		return fmt.Sprintf(`"%v"`, v)

	default:
		return fmt.Sprintf("%v", v)
	}
}

func rpad(v any, width int) string {
	format := fmt.Sprintf("%%-%vv", width)

	return fmt.Sprintf(format, v)
}

func article(v any) string {
	s := fmt.Sprintf("%v", v)

	if regexp.MustCompile("^[aeiouAEIOY].*").MatchString(s) {
		return "an"
	} else {
		return "a"
	}
}
