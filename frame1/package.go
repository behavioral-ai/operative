package frame1

import (
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

// IFrame - frame interface
type IFrame struct {
	Reason func(agent messaging.Agent, o Observation, append collective.Appender, resolver collective.Resolution)
}

var Frame = func() *IFrame {
	return &IFrame{
		Reason: func(agent messaging.Agent, o Observation, append collective.Appender, resolver collective.Resolution) {
			t, status := newThreshold(urn.ResiliencyThreshold, version, resolver)
			if !status.OK() {
				agent.Notify(status)
				return
			}
			i, status1 := newInterpret(urn.ResiliencyInterpret, version, resolver)
			if !status1.OK() {
				agent.Notify(status1)
				return
			}
			reason(o, t, i)
			append.Activity(agent, urn.ResiliencyActivity, "", nil)
			// Do we want to trace on error??
			//handler.Trace(handler, messaging.MasterChannel, messaging.ObservationEvent, result)
		},
	}
}()
