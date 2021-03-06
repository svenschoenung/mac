package mac

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func testGuessFormat(t *testing.T, s string, sizeOfType int, expected *Format) {
	if actual := guessFormat(s, sizeOfType); !cmp.Equal(actual, expected) {
		t.Errorf("guessFormat(\"%s\", %d)\nwas: %v\nexpected: %v\b", s, sizeOfType, actual, expected)
	}
}

func TestGuessFormat(t *testing.T) {
	testGuessFormat(t, "00-00-00-00-00-00", SizeOfEUI48, &Format{GroupsOf: 2, SeparatedBy: '-', WithCase: IgnoreCase})
	testGuessFormat(t, "00:00:00:00:00:00", SizeOfEUI48, &Format{GroupsOf: 2, SeparatedBy: ':', WithCase: IgnoreCase})
	testGuessFormat(t, "0000.0000.0000", SizeOfEUI48, &Format{GroupsOf: 4, SeparatedBy: '.', WithCase: IgnoreCase})
	testGuessFormat(t, "000000000000", SizeOfEUI48, &Format{GroupsOf: NoGroups, SeparatedBy: NoSeparator, WithCase: IgnoreCase})
	testGuessFormat(t, "00_00_00_00_00_00", SizeOfEUI48, nil)

	testGuessFormat(t, "00-00-00-00-00-00-00-00", SizeOfEUI64, &Format{GroupsOf: 2, SeparatedBy: '-', WithCase: IgnoreCase})
	testGuessFormat(t, "00:00:00:00:00:00:00:00", SizeOfEUI64, &Format{GroupsOf: 2, SeparatedBy: ':', WithCase: IgnoreCase})
	testGuessFormat(t, "0000.0000.0000.0000", SizeOfEUI64, &Format{GroupsOf: 4, SeparatedBy: '.', WithCase: IgnoreCase})
	testGuessFormat(t, "0000000000000000", SizeOfEUI64, &Format{GroupsOf: NoGroups, SeparatedBy: NoSeparator, WithCase: IgnoreCase})
	testGuessFormat(t, "00_00_00_00_00_00_00_00", SizeOfEUI64, nil)
}

func ExampleEUI48() {
	var addr = EUI48{0xC0, 0xDE, 0xC0, 0xDE, 0xC0, 0xDE}
	fmt.Println(addr)
	//Output: [192 222 192 222 192 222]
}

func ExampleEUI64() {
	var addr = EUI64{0xC0, 0xDE, 0xC0, 0xDE, 0xC0, 0xDE, 0xC0, 0xDE}
	fmt.Println(addr)
	//Output: [192 222 192 222 192 222 192 222]
}
