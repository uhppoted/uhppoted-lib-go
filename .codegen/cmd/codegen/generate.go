package main

import (
	_ "embed"
	"fmt"
	"os"
	"regexp"
	"strings"
	"text/template"

	"codegen/model"
)

var functions = template.FuncMap{
	"titleCase":   titleCase,
	"hyphenate":   hyphenate,
	"hex":         hex,
	"args":        args,
	"fields2args": fields2args,
	"pack":        pack,
	"describe":    describe,
	"lookup":      lookup,
	"value":       value,
}

func main() {
	if len(os.Args) > 1 {
		args := os.Args[1:]

		switch args[0] {
		case "codec":
			encode()
			encodeTest()
			decodeTest()

		case "integration-tests":
			integrationTests()
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

func args(args []any) string {
	var parts []string
	for _, a := range args {
		parts = append(parts, fmt.Sprintf("%v", a))
	}

	return strings.Join(parts, ", ")
}

func fields2args(fields []model.Field) string {
	var args []string
	for _, f := range fields {
		args = append(args, fmt.Sprintf("%v %v", f.Name, f.Type))
	}

	return strings.Join(args, ", ")
}

func pack(field model.Field) string {
	return fmt.Sprintf("packUint32(%v, packet, %v)", field.Name, field.Offset)
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

func value(v any, vtype string) string {
	switch vtype {
	case "IPv4":
		return fmt.Sprintf(`IPv4("%v")`, v)

	case "MAC":
		return fmt.Sprintf(`"%v"`, v)

	case "version":
		return fmt.Sprintf(`"%v"`, v)

	case "date":
		return fmt.Sprintf(`date("%v")`, v)

	case "string":
		return fmt.Sprintf(`"%v"`, v)

	default:
		return fmt.Sprintf("%v", v)
	}
}
