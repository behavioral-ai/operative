package frame1

import (
	"fmt"
	"github.com/behavioral-ai/collective/content"
	"github.com/behavioral-ai/core/messaging"
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

func Reason(o Observation, resolver content.Resolution) (Action, *messaging.Status) {
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
	return Action{Action: action, Desc: fmt.Sprintf("act:%v sat:%v grad:%v", action, imp.Saturation, imp.Gradient)}, messaging.StatusOK()
}
