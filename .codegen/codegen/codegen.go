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
	"titleCase": TitleCase,
	"camelCase": CamelCase,
	"hyphenate": hyphenate,
	"trim":      trim,
	"clean":     clean,
	"hex":       hex,
	"testargs":  testargs,
	"testarg":   testarg,
	"value":     value,
	"rpad":      rpad,
	"article":   article,
}

func TitleCase(s string) string {
	re := regexp.MustCompile(`[ \-:]+`)
	parts := re.Split(s, -1)
	for i := range parts {
		parts[i] = strings.Title(parts[i])
	}

	return strings.Join(parts, "")
}

func CamelCase(s string) string {
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

func trim(s string) string {
	return strings.TrimSuffix(s, "Response")
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

func testargs(args []lib.Arg) string {
	var parts []string
	for _, a := range args {
		parts = append(parts, testarg(a))
	}

	return strings.Join(parts, ", ")
}

func testarg(arg lib.Arg) string {
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
		return fmt.Sprintf(`types.MustParseDateTime("%v")`, arg.Value)

	case "optional datetime":
		return fmt.Sprintf(`types.MustParseDateTime("%v")`, arg.Value)

	case "date":
		return fmt.Sprintf(`types.MustParseDate("%v")`, arg.Value)

	case "HHmm":
		return fmt.Sprintf(`types.MustParseHHmm("%v")`, arg.Value)

	case "pin":
		return fmt.Sprintf(`uint32(%v)`, arg.Value)

	case "mode":
		return fmt.Sprintf(`types.DoorMode(%v)`, arg.Value)

	case "task":
		return fmt.Sprintf(`types.TaskType(%v)`, arg.Value)

	case "interlock":
		return fmt.Sprintf(`types.Interlock(%v)`, arg.Value)

	case "anti-passback":
		return fmt.Sprintf(`types.AntiPassback(%v)`, arg.Value)

	default:
		return fmt.Sprintf("%v", arg.Value)
	}
}

func value(v any, vtype string) string {
	switch vtype {
	case "IPv4":
		return fmt.Sprintf(`netip.MustParseAddr("%v")`, v)

	case "addrport":
		return fmt.Sprintf(`netip.MustParseAddrPort("%v")`, v)

	case "address:port":
		return fmt.Sprintf(`netip.MustParseAddrPort("%v")`, v)

	case "MAC":
		return fmt.Sprintf(`"%v"`, v)

	case "version":
		return fmt.Sprintf(`"%v"`, v)

	case "datetime":
		return fmt.Sprintf(`types.MustParseDateTime("%v")`, v)

	case "optional datetime":
		return fmt.Sprintf(`types.MustParseDateTime("%v")`, v)

	case "date":
		return fmt.Sprintf(`types.MustParseDate("%v")`, v)

	case "shortdate":
		return fmt.Sprintf(`types.MustParseDate("%v")`, v)

	case "optional date":
		return fmt.Sprintf(`types.MustParseDate("%v")`, v)

	case "time":
		return fmt.Sprintf(`types.MustParseTime("%v")`, v)

	case "HHmm":
		return fmt.Sprintf(`types.MustParseHHmm("%v")`, v)

	case "string":
		return fmt.Sprintf(`"%v"`, v)

	case "event-type":
		return fmt.Sprintf(`types.EventType(%v)`, v)

	case "mode":
		return fmt.Sprintf(`types.DoorMode(%v)`, v)

	case "direction":
		return fmt.Sprintf(`types.Direction(%v)`, v)

	case "reason":
		return fmt.Sprintf(`types.Reason(%v)`, v)

	case "anti-passback":
		return fmt.Sprintf(`types.AntiPassback(%v)`, v)

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
