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

type Activity struct {
	Action int
	Desc   string
}

func Reason(o Observation, resolver collective.Resolution) (*messaging.Status, Activity) {
	t, status := newThreshold(urn.ResiliencyThreshold, version, resolver)
	if !status.OK() {
		return status, Activity{}
	}
	i, status1 := newInterpret(urn.ResiliencyInterpret, version, resolver)
	if !status1.OK() {
		return status, Activity{}
	}
	imp := t.comprehend(o)
	action := i.action(imp)
	return messaging.StatusOK(), Activity{Action: action, Desc: fmt.Sprintf("action: %v gradient: %v saturation: %v name:%v", action, imp.Gradient, imp.Saturation, urn.ResiliencyActivity)}
}
