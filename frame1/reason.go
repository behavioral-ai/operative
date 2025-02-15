package frame1

import (
	"fmt"
	"github.com/behavioral-ai/operative/urn"
)

type activity struct {
	Action int `json:"action"`
}

func (a activity) body() []byte {
	return []byte(fmt.Sprintf("{ \"action\" : %v }", a.Action))
}

func reason(o Observation, t threshold, i interpret) (activity, string) {
	imp := t.comprehend(o)
	action := i.action(imp)
	return activity{Action: action}, fmt.Sprintf("action: %v gradient: %v saturation: %v name:%v", action, imp.Gradient, imp.Saturation, urn.ResiliencyActivity)
}
