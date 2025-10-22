package api

import (
	"fmt"
	"os"
)

func writeln(f *os.File, s string) {
	if _, err := f.WriteString(s + "\n"); err != nil {
		panic(fmt.Errorf("error writing to %v (%v)", f.Name(), err))
	}
}
