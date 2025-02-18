package frame1

import "github.com/behavioral-ai/domain/collective"

type lookup struct {
	Low    int `json:"low"`
	Medium int `json:"medium"`
	High   int `json:"high"`
}

type threshold struct {
	Latency    int    `json:"latency"`
	Saturation lookup `json:"saturation"`
	Gradient   lookup `json:"gradient"`
}

func (d *threshold) find(l lookup, value int) string {
	if value <= l.Low {
		return low
	}
	if value <= l.Medium {
		return medium
	}
	return high
}

func (d *threshold) comprehend(o Observation) impression {
	sat := float64(o.Latency()) / float64(d.Latency)
	return impression{Gradient: d.find(d.Gradient, o.Gradient()), Saturation: d.find(d.Saturation, int(sat*100))}
}

func newThreshold(name string, version int) (threshold, error) {
	return collective.Resolve[threshold](name, version)
}
