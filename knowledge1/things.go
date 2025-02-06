package knowledge1

const (
	Low    = "low"
	Medium = "medium"
	High   = "high"
)

// observation -> comprehend -> impression
// impression -> reason -> action

// Impression -
type Impression struct {
	Saturation string
	Gradient   string
}

// Threshold -
type Threshold struct {
	Latency int `json:"latency"`
	Low     int `json:"low"`
	Medium  int `json:"medium"`
	High    int `json:"high"` // Milliseconds
}

func (t *Threshold) Impression(o *Observation) Impression {
	s := float64(o.Latency) / float64(t.Latency)
	s *= 100
	if s <= float64(t.Low) {
		return Impression{Saturation: Low}
	}
	if s <= float64(t.Medium) {
		return Impression{Saturation: Medium}
	}
	return Impression{Saturation: High}
}
