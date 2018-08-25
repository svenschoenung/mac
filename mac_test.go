package mac

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func testGuessFormat(t *testing.T, s string, sizeOfType int, expected *Format) {
	if actual := guessFormat(s, sizeOfType); !cmp.Equal(actual, expected) {
		t.Errorf("guessFormat(\"%s\", %d)\nwas: %v\nexpected: %v\b", s, sizeOfType, actual, expected)
	}
}

func TestGuessFormat(t *testing.T) {
	testGuessFormat(t, "00-00-00-00-00-00", SizeOfEUI48, &Format{SeparatedBy: '-', GroupsOf: 2, WithCase: IgnoreCase})
	testGuessFormat(t, "00:00:00:00:00:00", SizeOfEUI48, &Format{SeparatedBy: ':', GroupsOf: 2, WithCase: IgnoreCase})
	testGuessFormat(t, "0000.0000.0000", SizeOfEUI48, &Format{SeparatedBy: '.', GroupsOf: 4, WithCase: IgnoreCase})
	testGuessFormat(t, "000000000000", SizeOfEUI48, &Format{SeparatedBy: NoSeparator, GroupsOf: NoGroups, WithCase: IgnoreCase})
	testGuessFormat(t, "00_00_00_00_00_00", SizeOfEUI48, nil)

	testGuessFormat(t, "00-00-00-00-00-00-00-00", SizeOfEUI64, &Format{SeparatedBy: '-', GroupsOf: 2, WithCase: IgnoreCase})
	testGuessFormat(t, "00:00:00:00:00:00:00:00", SizeOfEUI64, &Format{SeparatedBy: ':', GroupsOf: 2, WithCase: IgnoreCase})
	testGuessFormat(t, "0000.0000.0000.0000", SizeOfEUI64, &Format{SeparatedBy: '.', GroupsOf: 4, WithCase: IgnoreCase})
	testGuessFormat(t, "0000000000000000", SizeOfEUI64, &Format{SeparatedBy: NoSeparator, GroupsOf: NoGroups, WithCase: IgnoreCase})
	testGuessFormat(t, "00_00_00_00_00_00_00_00", SizeOfEUI64, nil)

}
