package frame1

import (
	"fmt"
	"github.com/behavioral-ai/operative/urn"
)

type activity struct {
	action int
}

func (a activity) body() []byte {
	return []byte(fmt.Sprintf("{ \"action\" : %v }", a.action))
}

func reason(o Observation, t threshold, i interpret) (activity, string) {
	imp := t.comprehend(o)
	action := i.action(imp)
	return activity{action: action}, fmt.Sprintf("action: %v gradient: %v saturation: %v name:%v", action, imp.Gradient, imp.Saturation, urn.ResiliencyActivity)
}
