// Code generated by "stringer -type=Type -linecomment -output bintype_string_gen.go"; DO NOT EDIT.

package bin

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[GardenWaste-1]
	_ = x[Recycling-2]
	_ = x[Rubbish-3]
}

const _Type_name = "UnknownGardenWasteRecyclingRubbish"

var _Type_index = [...]uint8{0, 7, 18, 27, 34}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
