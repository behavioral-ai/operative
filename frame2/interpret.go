package frame2

import "github.com/behavioral-ai/domain/collective"

const (
	low    = "low"
	medium = "medium"
	high   = "high"
)

type impression struct {
	Saturation string `json:"saturation"`
	Gradient   string `json:"gradient"`
}

type interpret struct {
	LowLow       int `json:"low-low"`
	LowMedium    int `json:"low-medium"`
	LowHigh      int `json:"low-high"`
	MediumLow    int `json:"medium-low"`
	MediumMedium int `json:"medium-medium"`
	MediumHigh   int `json:"medium-high"`
	HighLow      int `json:"high-low"`
	HighMedium   int `json:"high-medium"`
	HighHigh     int `json:"high-high"`
}

func (a interpret) action(i impression) int {
	switch i.Saturation {
	case low:
		switch i.Gradient {
		case low:
			return a.LowLow
		case medium:
			return a.LowMedium
		case high:
			return a.LowHigh
		}
	case medium:
		switch i.Gradient {
		case low:
			return a.MediumLow
		case medium:
			return a.MediumMedium
		case high:
			return a.MediumHigh
		}
	case high:
		switch i.Gradient {
		case low:
			return a.HighLow
		case medium:
			return a.HighMedium
		case high:
			return a.HighHigh
		}
	}
	return 0
}

func newInterpret(name string, resolver collective.Resolution) (interpret, error) {
	//t, err := collectiveresolver..Get(name,version)
	return interpret{}, nil
}
