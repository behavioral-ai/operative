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

func Reason(o Observation, resolver collective.Resolution) (Activity, *messaging.Status) {
	t, status := newThreshold(version, resolver)
	if !status.OK() {
		return Activity{}, status
	}
	i, status1 := newInterpret(version, resolver)
	if !status1.OK() {
		return Activity{}, status
	}
	imp := t.comprehend(o)
	action := i.action(imp)
	return Activity{Action: action, Desc: fmt.Sprintf("action: %v gradient: %v saturation: %v name:%v", action, imp.Gradient, imp.Saturation, urn.ResiliencyActivity)}, messaging.StatusOK()
}
