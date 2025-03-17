package bin

import (
	"encoding/json"
	"fmt"
)

//go:generate stringer -type=Type -linecomment -output bintype_string_gen.go
type Type int

const (
	Unknown Type = iota
	GardenWaste
	Recycling
	Rubbish
)

func (t *Type) UnmarshalJSON(b []byte) error {
	switch string(b) {
	case `"Rubbish"`:
		*t = Rubbish
	case `"Garden Waste"`:
		*t = GardenWaste
	case `"Recycling"`:
		*t = Recycling
	default:
		*t = Unknown
		return fmt.Errorf("unable to parse bin type %s", string(b))
	}
	return nil
}

func (t Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprint(t)) // using fmt and not bt.String() so that if the go:generate stringer tool hasnt run, this will still just print the Int of the Enum
}
