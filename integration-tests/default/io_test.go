package uhppoted

import (
	"reflect"
	"testing"

	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/io"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/responses"
	"github.com/uhppoted/uhppoted-lib-go/src/uhppoted/types"
)

func TestIOSetFirstCard(t *testing.T) {
	controller := uint32(405419896)
	door := uint8(3)

	startTime := types.MustParseHHmm("08:30")
	endTime := types.MustParseHHmm("17:45")
	activeMode := types.NormallyOpen
	inactiveMode := types.NormallyClosed
	monday := true
	tuesday := true
	wednesday := false
	thursday := true
	friday := false
	saturday := true
	sunday := true

	expected := responses.SetFirstCard{
		Controller: 405419896,
		Ok:         true,
	}

	response, err := io.SetFirstCard(u,
		controller,
		door,
		startTime, endTime,
		activeMode, inactiveMode,
		monday, tuesday, wednesday, thursday, friday, saturday, sunday,
		timeout)

	if err != nil {
		t.Fatalf("%v", err)
	} else if !reflect.DeepEqual(response, expected) {
		t.Errorf("incorrect response\n   expected:%v\n   got:     %v", expected, response)
	}
}
