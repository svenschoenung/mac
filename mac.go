package mac

// SizeOfEUI48 is the number of bytes that an EUI-48 address holds
const SizeOfEUI48 = 6

// SizeOfEUI64 is the number of bytes that an EUI-64 address holds
const SizeOfEUI64 = 8

// EUI48 represents an EUI-48 or MAC-48 address
type EUI48 [SizeOfEUI48]byte

// EUI64 represents an EUI-64 address
type EUI64 [SizeOfEUI64]byte

// NoSeparator indicates that a formatted address doesn't have a seperator character
const NoSeparator = rune(0)

// NoGroups indicates that a formatted address doesn't have grouped hex digits
const NoGroups = 0

// Case represents the case of a formatted address
type Case byte

// Constants that indicate the case of a formatted address
const (
	IgnoreCase = Case(iota)
	LowerCase
	UpperCase
)

// Format holds information about how an address should be formatted or is formatted
type Format struct {
	GroupsOf    int
	SeparatedBy rune
	WithCase    Case
}

func guessFormat(s string, sizeOfType int) *Format {
	if len(s) == sizeOfType*2 {
		return &Format{GroupsOf: NoGroups, SeparatedBy: NoSeparator, WithCase: IgnoreCase}
	}
	for i, r := range s {
		if r == ':' || r == '-' || r == '.' {
			return &Format{GroupsOf: i, SeparatedBy: r, WithCase: IgnoreCase}
		}
	}
	return nil
}
