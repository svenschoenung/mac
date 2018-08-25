package mac

const SizeOfEUI48 = 6
const SizeOfEUI64 = 8

type EUI48 [SizeOfEUI48]byte
type EUI64 [SizeOfEUI64]byte

const NoSeparator = rune(0)
const NoGroups = 0

type Case byte

const (
	IgnoreCase = Case(iota)
	LowerCase
	UpperCase
)

type Format struct {
	SeparatedBy rune
	GroupsOf    int
	WithCase    Case
}

func guessFormat(s string, sizeOfType int) *Format {
	if len(s) == sizeOfType*2 {
		return &Format{SeparatedBy: NoSeparator, GroupsOf: NoGroups, WithCase: IgnoreCase}
	}
	for i, r := range s {
		if r == ':' || r == '-' || r == '.' {
			return &Format{SeparatedBy: r, GroupsOf: i, WithCase: IgnoreCase}
		}
	}
	return nil
}
