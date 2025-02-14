package frame1

const (
	low    = "low"
	medium = "medium"
	high   = "high"
)

type impression struct {
	Saturation string `json:"saturation"`
	Gradient   string `json:"gradient"`
}

type lookup struct {
	Low    int `json:"low"`
	Medium int `json:"medium"`
	High   int `json:"high"`
}

type data struct {
	Threshold  int    `json:"threshold"`
	Saturation lookup `json:"saturation"`
	Gradient   lookup `json:"gradient"`
}

func (d *data) find(l lookup, value int) string {
	if value <= l.Low {
		return low
	}
	if value <= l.Medium {
		return medium
	}
	return high
}

func (d *data) comprehend(o Observation) impression {
	sat := float64(o.Latency()) / float64(d.Threshold)
	return impression{Gradient: d.find(d.Gradient, o.Gradient()), Saturation: d.find(d.Saturation, int(sat*100))}
}

func comprehend(o Observation) impression {

	return impression{}
}
