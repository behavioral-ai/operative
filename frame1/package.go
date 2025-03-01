package frame1

import (
	"fmt"
	"github.com/behavioral-ai/core/messaging"
	"github.com/behavioral-ai/domain/collective"
	"github.com/behavioral-ai/operative/urn"
)

const (
	PkgPath = "github/behavioral-ai/operative/frame1"
	version = 1
)

type Observation interface {
	Gradient() int
	Latency() int
}

type Action struct {
	Action int
	Desc   string
}

func Reason(o Observation, resolver collective.Resolution) (Action, *messaging.Status) {
	t, status := newThreshold(version, resolver)
	if !status.OK() {
		return Action{}, status
	}
	i, status1 := newInterpret(version, resolver)
	if !status1.OK() {
		return Action{}, status
	}
	imp := t.comprehend(o)
	action := i.action(imp)
	return Action{Action: action, Desc: fmt.Sprintf("action: %v gradient: %v saturation: %v %v", action, imp.Gradient, imp.Saturation, urn.ResiliencyAction)}, messaging.StatusOK()
}
