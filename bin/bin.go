package bin

import (
	"time"
)

type Bin struct {
	Type       Type   `json:"BinType"`
	PdfLink    string `json:"PdfLink,omitempty"`
	Communal   bool   `json:"Communal,string"`
	Next       BCPTime
	Subsequent BCPTime
}

type BCPTime struct {
	time.Time
}

func (t *BCPTime) UnmarshalJSON(b []byte) error {
	parsed, err := time.Parse(`"1/2/2006 15:04:05 PM"`, string(b))
	if err != nil {
		return err
	}
	t.Time = parsed

	return nil
}
